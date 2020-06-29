package generator

import (
	"os"
	"path/filepath"
)

type Workspace struct {
	HomeDir         string
	TemplateDir     string
	Module          string // go module name for this project
	ResourceDirs    []string
	CurrentResource string
}

func DetectWorkspace(baseDir string) (*Workspace, error) {
	var p Workspace
	var err error
	p.HomeDir, err = DetectGoProjectHome(baseDir)
	if err != nil {
		return nil, err
	}

	p.Module, err = DetectGoModule(p.HomeDir)
	if err != nil {
		return nil, err
	}

	//p.ResourceDirs, err = DetectServiceDirs(p.HomeDir)
	//if err != nil {
	//	return nil, err
	//}
	//
	//exists, err := ProtoFileExists(baseDir)
	//if err != nil {
	//	return nil, err
	//}
	//if exists {
	//	dir, _ := os.Getwd()
	//	p.CurrentResource = filepath.Base(dir)
	//}

	taoHomePath := filepath.Join(os.Getenv("HOME"), ".tao")
	p.TemplateDir = filepath.Join(taoHomePath, "src/tools/tao/templates")

	return &p, nil
}

func (w Workspace) TemplatesExists() (bool, error) {
	_, err := os.Stat(w.TemplateDir)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
