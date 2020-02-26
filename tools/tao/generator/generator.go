package generator

import (
	"errors"
	"os"
	"path/filepath"
	"text/template"
)

var (
	ErrModelNotFound   = errors.New("model message not found")
	ErrServiceNotFound = errors.New("service not found")
)

type Generator interface {
	Generate(templateFile string, data interface{}, outputFile string) error
}

type SimpleGenerator int

func (s SimpleGenerator) Generate(templateFile string, data interface{}, outputFile string) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	tpl, err := template.New(filepath.Base(templateFile)).ParseFiles(templateFile)
	if err != nil {
		return err
	}

	return tpl.Execute(f, data)
}
