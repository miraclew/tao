package generator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/alecthomas/participle"
	"github.com/iancoleman/strcase"
	"github.com/miraclew/tao/tools/tao/mapper"
	"github.com/miraclew/tao/tools/tao/mapper/dart"
	"github.com/miraclew/tao/tools/tao/mapper/golang"
	"github.com/miraclew/tao/tools/tao/mapper/openapiv3"
	"github.com/miraclew/tao/tools/tao/mapper/sqlschema"
	"github.com/miraclew/tao/tools/tao/mapper/swift"
	"github.com/miraclew/tao/tools/tao/parser"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type Engine struct {
	Workspace *Workspace
	Config    *Config
}

func NewEngine() (*Engine, error) {
	return NewEngineWithBaseDir(".")
}

func NewEngineWithBaseDir(dir string) (*Engine, error) {
	workspace, err := DetectWorkspace(dir)
	if err != nil {
		return nil, err
	}

	config, err := NewConfig(workspace.HomeDir)
	if err != nil {
		return nil, err
	}
	return &Engine{Workspace: workspace, Config: config}, nil
}

func (e Engine) GenerateLocator() error {
	// create locator.go
	dir := filepath.Join(e.Workspace.HomeDir, "locator")
	_ = os.Mkdir(dir, 0755)
	fileName := filepath.Join(dir, "locator.go")
	outputFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	var rs []mapper.Resource
	for _, r := range e.Workspace.ResourceDirs {
		proto, err := parser.ParseProto3(filepath.Join(e.Workspace.HomeDir, r, r+".proto"))
		if err != nil {
			return err
		}
		rs = append(rs, mapper.Resource{
			Module:   e.Workspace.Module,
			Pkg:      r,
			Name:     proto.ResourceName,
			HasEvent: proto.EventService != nil,
		})
	}

	// fetch dependencies
	if e.Config != nil {
		for _, dir := range e.Config.Dependencies {
			w, err := DetectWorkspace(dir)
			if err != nil {
				return err
			}
			for _, r := range w.ResourceDirs {
				proto, err := parser.ParseProto3(filepath.Join(w.HomeDir, r, r+".proto"))
				if err != nil {
					return err
				}
				rs = append(rs, mapper.Resource{
					Module:   w.Module,
					Pkg:      r,
					Name:     proto.ResourceName,
					HasEvent: proto.EventService != nil,
				})
			}
		}
	}

	model := mapper.Locator{
		Module:    e.Workspace.Module,
		Resources: rs,
	}

	tplFile := filepath.Join(e.Workspace.TemplateDir, "locator/locator.go.tpl")
	tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
	if err != nil {
		return err
	}
	return tpl.Execute(outputFile, model)
}

func (e Engine) GenerateSql() error {
	for _, r := range e.Workspace.ResourceDirs {
		proto, err := parser.ParseProto3(filepath.Join(e.Workspace.HomeDir, r, r+".proto"))
		if err != nil {
			return err
		}

		models := getModelMessages(proto.Proto)
		schemaModel, err := sqlschema.MapCreateTables(models)
		if err != nil {
			return err
		}

		protoGolang, err := golang.Map(proto.Proto, e.Config.UseSnackCase)
		if err != nil {
			return err
		}

		// fix enum
		for _, model := range schemaModel.Items {
			for _, field := range model.Columns {
				for _, enum := range protoGolang.Enums {
					if enum.Name == field.Type {
						field.Type = "int(11)"
					}
				}
			}
		}

		sqlDir := filepath.Join(e.Workspace.HomeDir, "doc/sql")
		_ = os.Mkdir(sqlDir, 0755)
		fileName := filepath.Join(sqlDir, r+".sql")

		outputFile, err := os.Create(fileName)
		if err != nil {
			return err
		}
		tplFile := filepath.Join(e.Workspace.TemplateDir, "doc/sql_schema/mysql.sql.tpl")

		tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
		if err != nil {
			return err
		}
		err = tpl.Execute(outputFile, schemaModel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Engine) GenerateSwift(protoFile string) error {
	baseFileName := filepath.Base(protoFile)
	baseFileName = strings.TrimSuffix(baseFileName, filepath.Ext(baseFileName))

	var p = participle.MustBuild(&proto3.Proto{}, participle.UseLookahead(2))
	proto := &proto3.Proto{}
	file, err := os.Open(protoFile)
	if err != nil {
		return err
	}
	defer file.Close()
	err = p.Parse(file, proto)
	if err != nil {
		return err
	}

	pm := swift.NewProtoMapper()
	model, err := pm.Map(proto)
	if err != nil {
		return err
	}

	var fileName string
	if e.Config != nil && e.Config.SwiftOutputDir != "" {
		fileName = filepath.Join(e.Config.SwiftOutputDir, strings.Title(baseFileName)+".swift")
	} else {
		swiftDir := filepath.Join(e.Workspace.HomeDir, "doc/swift")
		_ = os.Mkdir(swiftDir, 0755)
		fileName = filepath.Join(swiftDir, strings.Title(baseFileName)+".swift")
	}

	fmt.Printf("output: %s", fileName)
	outputFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	tplFile := filepath.Join(e.Workspace.TemplateDir, "sdk/swift/client.swift.tpl")

	tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
	if err != nil {
		return err
	}
	err = tpl.Execute(outputFile, model)
	if err != nil {
		return err
	}
	return nil
}

func (e Engine) GenerateDart() error {
	for _, r := range e.Workspace.ResourceDirs {
		var p = participle.MustBuild(&proto3.Proto{}, participle.UseLookahead(2))
		proto := &proto3.Proto{}
		file, err := os.Open(filepath.Join(e.Workspace.HomeDir, r, r+".proto"))
		if err != nil {
			return err
		}
		defer file.Close()
		err = p.Parse(file, proto)
		if err != nil {
			return err
		}

		pm := dart.NewProtoMapper()
		model, err := pm.Map(proto)
		if err != nil {
			return err
		}

		var fileName string
		if e.Config != nil && e.Config.DartOutputDir != "" {
			fileName = filepath.Join(e.Config.DartOutputDir, r+".dart")
		} else {
			dartDir := filepath.Join(e.Workspace.HomeDir, "doc/dart")
			_ = os.Mkdir(dartDir, 0755)
			fileName = filepath.Join(dartDir, r+".dart")
		}

		outputFile, err := os.Create(fileName)
		if err != nil {
			return err
		}
		tplFile := filepath.Join(e.Workspace.TemplateDir, "sdk/dart/client.swift.tpl")

		tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
		if err != nil {
			return err
		}
		err = tpl.Execute(outputFile, model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Engine) GenerateOpenAPIV3() error {
	for _, r := range e.Workspace.ResourceDirs {
		proto, err := parser.ParseProto3(filepath.Join(e.Workspace.HomeDir, r, r+".proto"))
		if err != nil {
			return err
		}
		schemaModel, err := openapiv3.MapProto2Doc(proto)
		if err != nil {
			return err
		}

		sqlDir := filepath.Join(e.Workspace.HomeDir, "doc/openapiv3")
		_ = os.Mkdir(sqlDir, 0755)
		fileName := filepath.Join(sqlDir, r+".yaml")

		outputFile, err := os.Create(fileName)
		if err != nil {
			return err
		}
		tplFile := filepath.Join(e.Workspace.TemplateDir, "doc/open_api_v3/open_api.yaml.tpl")

		var tpl *template.Template
		fm := template.FuncMap{
			"include": func(name string, data interface{}) (string, error) {
				buf := bytes.NewBuffer(nil)
				if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
					return "", err
				}
				return buf.String(), nil
			},
		}

		tpl, err = template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).Funcs(fm).ParseFiles(tplFile)
		if err != nil {
			return err
		}
		err = tpl.Execute(outputFile, schemaModel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Engine) GenerateProto() error {
	// e.Workspace.CurrentResource not available
	dir, _ := os.Getwd()
	pkg := filepath.Base(dir)

	tplFile := filepath.Join(e.Workspace.TemplateDir, "api/resource.proto.tpl")
	fileName := fmt.Sprintf("%s.proto", strcase.ToSnake(pkg))
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
	if err != nil {
		return err
	}

	return tpl.Execute(f, map[string]interface{}{
		"Resource": strcase.ToCamel(pkg),
	})
}

func (e Engine) GenerateAPI(pbFile string) error {
	res, err := parser.ParseProto3(pbFile)
	if err != nil {
		return err
	}

	protoGolang, err := golang.Map(res.Proto, e.Config.UseSnackCase)
	if err != nil {
		return err
	}
	//protoGolang.Module = e.Workspace.Module
	protoGolang.Module = "hello.world.module"

	files := []string{"api", "client"}
	for _, file := range files {
		outputFile, err := os.Create(filepath.Join(e.Config.GoOutputDir, fmt.Sprintf("%s.go", file)))
		if err != nil {
			return err
		}

		tplFile := filepath.Join(e.Workspace.TemplateDir, fmt.Sprintf("api/%s.go.tpl", file))
		tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
		if err != nil {
			return err
		}
		err = tpl.Execute(outputFile, protoGolang)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e Engine) GenerateService(useDefault bool) error {
	if e.Workspace.CurrentResource == "" {
		return errors.New("this command should be execute in resource dir")
	}
	_ = os.MkdirAll("svc", 0755)

	var p = participle.MustBuild(&proto3.Proto{}, participle.UseLookahead(2))
	proto := &proto3.Proto{}
	r, err := os.Open(e.Workspace.CurrentResource + ".proto")
	if err != nil {
		return err
	}
	defer r.Close()
	err = p.Parse(r, proto)
	if err != nil {
		return err
	}

	protoGolang, err := golang.Map(proto, e.Config.UseSnackCase)
	if err != nil {
		return err
	}
	protoGolang.Module = e.Workspace.Module

	files := []string{"handler.api", "handler.event", "service"}
	tplFiles := []string{"handler.api", "handler.event", "service"}
	if useDefault {
		tplFiles[2] = "service_default"
	}
	for i, file := range files {
		outputFile, err := os.Create(fmt.Sprintf("svc/%s.go", file))
		if err != nil {
			return err
		}

		tplFile := filepath.Join(e.Workspace.TemplateDir, fmt.Sprintf("svc/%s.go.tpl", tplFiles[i]))
		tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
		if err != nil {
			return err
		}
		err = tpl.Execute(outputFile, protoGolang)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: proto support more than one model messages,  how to generate repo code then?
func (e Engine) GenerateRepo() error {
	if e.Workspace.CurrentResource == "" {
		return errors.New("this command should be execute in resource dir")
	}
	_ = os.MkdirAll("svc", 0755)

	files := []string{"repo", "repo.mysql", "repo.redis"}
	for _, file := range files {
		outputFile, err := os.Create(fmt.Sprintf("svc/%s.go", file))
		if err != nil {
			return err
		}

		proto, err := parser.ParseProto3(e.Workspace.CurrentResource + ".proto")
		if err != nil {
			return err
		}

		model, err := golang.Map(proto.Proto, e.Config.UseSnackCase)
		if err != nil {
			return err
		}
		model.Module = e.Workspace.Module

		tplFile := filepath.Join(e.Workspace.TemplateDir, fmt.Sprintf("svc/%s.go.tpl", file))
		tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
		if err != nil {
			return err
		}
		err = tpl.Execute(outputFile, model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e Engine) UpdateTemplates() error {
	taoHomePath := filepath.Join(os.Getenv("HOME"), ".tao")
	_, err := os.Stat(taoHomePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = os.Mkdir(taoHomePath, 0755)
		if err != nil {
			return err
		}
	}

	repositoryPath := filepath.Join(taoHomePath, "src")
	_, err = os.Stat(repositoryPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		// clone repository
		cmd := exec.Command("git", "clone", "https://github.com/miraclew/tao.git", repositoryPath)
		cmd.Dir = taoHomePath
		if err = cmd.Run(); err != nil {
			return err
		}
	} else {
		// pull
		cmd := exec.Command("git", "pull")
		cmd.Dir = repositoryPath
		if err = cmd.Run(); err != nil {
			return err
		}
	}
	s := filepath.Join(repositoryPath, "tools/tao/templates")

	fmt.Printf("update template done. (%s)\n", s)
	return nil
}

func getModelMessages(proto *proto3.Proto) []*proto3.Message {
	var msgs []*proto3.Message
	for _, e := range proto.Entries {
		if e.Message == nil {
			continue
		}

		var isModel bool
		for _, e2 := range e.Message.Entries {
			if e2.Option != nil {
				if e2.Option.Name == "model" {
					isModel = true
					break
				}
			}
		}
		if !isModel {
			continue
		}

		msgs = append(msgs, e.Message)
	}

	return msgs
}
