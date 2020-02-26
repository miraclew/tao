package engine

import (
	"bytes"

	"github.com/miraclew/tao/tools/tao/generator"
	"github.com/miraclew/tao/tools/tao/mapper"

	"github.com/miraclew/tao/tools/tao/mapper/dart"

	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/alecthomas/participle"
	"github.com/miraclew/tao/tools/tao/mapper/golang"
	"github.com/miraclew/tao/tools/tao/mapper/openapiv3"
	"github.com/miraclew/tao/tools/tao/parser"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type Engine struct {
	Workspace *generator.Workspace
	Config    *Config
}

func NewEngine() (*Engine, error) {
	workspace, err := generator.NewWorkspace()
	if err != nil {
		return nil, err
	}

	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return &Engine{Workspace: workspace, Config: config}, nil
}

func (e Engine) GenerateLocator() error {
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
			Name:     r,
			HasEvent: proto.EventService != nil,
		})
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
		_ = r
		//proto, err := parser.ParseProto3(filepath.Join(e.Workspace.HomeDir, r, r+".proto"))
		//if err != nil {
		//	return err
		//}
		//schemaModel, err := sqlschema.MapMessage2CreateTable(proto.ResourceMessage)
		//if err != nil {
		//	return err
		//}
		//
		//enums := gocode.ProtoEnums(proto.Proto)
		//// fix enum
		//for _, field := range schemaModel.Fields {
		//	for _, enum := range enums {
		//		if enum.Name == field.Type {
		//			field.Type = "int(11)"
		//		}
		//	}
		//}
		//
		//sqlDir := filepath.Join(e.Workspace.HomeDir, "doc/sql")
		//_ = os.Mkdir(sqlDir, 0755)
		//fileName := filepath.Join(sqlDir, r+".sql")
		//
		//outputFile, err := os.Create(fileName)
		//if err != nil {
		//	return err
		//}
		//tplFile := filepath.Join(e.Workspace.TemplateDir, "doc/sql_schema/mysql.sql.tpl")
		//
		//tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
		//if err != nil {
		//	return err
		//}
		//err = tpl.Execute(outputFile, schemaModel)
		//if err != nil {
		//	return err
		//}
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
		tplFile := filepath.Join(e.Workspace.TemplateDir, "sdk/dart/client.dart.tpl")

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
	//dir, _ := os.Getwd()
	//pkg := filepath.Base(dir)

	//model := api.Proto{
	//	Pkg:      pkg,
	//	Resource: strings.Title(pkg),
	//}
	//
	//tplFile := filepath.Join(e.Workspace.TemplateDir, "api/resource.proto.tpl")
	//fileName := fmt.Sprintf("%s.proto", strcase.ToSnake(pkg))
	//f, err := os.Create(fileName)
	//if err != nil {
	//	return err
	//}
	//
	//tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
	//if err != nil {
	//	return err
	//}
	//
	//return tpl.Execute(f, model)
	return nil
}

func (e Engine) GenerateAPI() error {
	if e.Workspace.CurrentResource == "" {
		return errors.New("this command should be execute in resource dir")
	}

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

	protoGolang, err := golang.Map(proto)
	if err != nil {
		return err
	}
	protoGolang.Module = e.Workspace.Module

	files := []string{"api", "client"}
	for _, file := range files {
		outputFile, err := os.Create(fmt.Sprintf("%s.go", file))
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

func (e Engine) GenerateService() error {
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

	protoGolang, err := golang.Map(proto)
	if err != nil {
		return err
	}
	protoGolang.Module = e.Workspace.Module

	files := []string{"service", "handler.api", "handler.event"}
	for _, file := range files {
		outputFile, err := os.Create(fmt.Sprintf("svc/%s.go", file))
		if err != nil {
			return err
		}

		tplFile := filepath.Join(e.Workspace.TemplateDir, fmt.Sprintf("svc/%s.go.tpl", file))
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

func (e Engine) GenerateRepo() error {
	if e.Workspace.CurrentResource == "" {
		return errors.New("this command should be execute in resource dir")
	}
	_ = os.MkdirAll("svc", 0755)

	//files := []string{"repo", "repo.mysql", "repo.redis"}
	//for _, file := range files {
	//	outputFile, err := os.Create(fmt.Sprintf("svc/%s.go", file))
	//	if err != nil {
	//		return err
	//	}
	//
	//	proto, err := parser.ParseProto3(e.Workspace.CurrentResource + ".proto")
	//	if err != nil {
	//		return err
	//	}
	//	model := svc.Repo{
	//		Pkg:      e.Workspace.CurrentResource,
	//		Module:   e.Workspace.Module,
	//		Resource: e.Workspace.CurrentResource,
	//		Table:    strings.Title(e.Workspace.CurrentResource),
	//		Fields:   parser.ResourceFields(proto.ResourceMessage, false),
	//	}
	//
	//	tplFile := filepath.Join(e.Workspace.TemplateDir, fmt.Sprintf("api/svc/%s.go.tpl", file))
	//	tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(tplFile)
	//	if err != nil {
	//		return err
	//	}
	//	err = tpl.Execute(outputFile, model)
	//	if err != nil {
	//		return err
	//	}
	//}

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
