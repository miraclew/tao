package generator

import (
	"testing"
)

func TestEngine_GenerateAPI(t *testing.T) {
	protoFile := "../testdata/demo.proto"
	e, err := NewEngineWithBaseDir("../../..")
	if err != nil {
		t.Error(err)
		return
	}
	e.Config = &Config{
		GoOutputDir:    "../testdata",
		SwiftOutputDir: "",
		DartOutputDir:  "",
		UseSnackCase:   false,
		Dependencies:   nil,
	}
	e.Workspace = &Workspace{
		HomeDir:         "",
		TemplateDir:     "../templates",
		Module:          "Hello",
		ResourceDirs:    nil,
		CurrentResource: "",
	}

	err = e.GenerateAPI(protoFile)
	if err != nil {
		t.Error(err)
	}
}

func TestEngine_GenerateSwift(t *testing.T) {
	protoFile := "../testdata/demo.proto"
	e, err := NewEngineWithBaseDir("../../..")
	if err != nil {
		t.Error(err)
		return
	}
	e.Config = &Config{
		GoOutputDir:    "../testdata",
		SwiftOutputDir: "../testdata",
		DartOutputDir:  "",
		UseSnackCase:   false,
		Dependencies:   nil,
	}
	e.Workspace = &Workspace{
		HomeDir:         "",
		TemplateDir:     "../templates",
		Module:          "hello",
		ResourceDirs:    nil,
		CurrentResource: "",
	}
	err = e.GenerateSwift(protoFile)
	if err != nil {
		t.Error(err)
	}
}
