package generator

import (
	"github.com/Masterminds/sprig"
	"os"
	"path/filepath"
	"text/template"
)

type APIGenerator int

func (g APIGenerator) Exec(w *Workspace) error {
	//templateFile string, data interface{}, outputFile string
	templateFile := filepath.Join(w.TemplateDir, "api/api.go.tpl")
	outputFile := filepath.Join(w.HomeDir, w.CurrentResource, "api.go")
	var data interface{}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	tpl, err := template.New(filepath.Base(templateFile)).Funcs(sprig.TxtFuncMap()).ParseFiles(templateFile)
	if err != nil {
		return err
	}
	if err = tpl.Execute(f, data); err != nil {
		return err
	}

	return nil
}
