package generator

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/miraclew/tao/pkg/tos"
)

type Workspace struct {
	HomeDir         string
	TemplateDir     string
	Module          string // go module name for this project
	ResourceDirs    []string
	CurrentResource string
}

func NewWorkspace() (*Workspace, error) {
	var p Workspace
	var err error
	if err = detectWorkingPath(&p); err != nil {
		return nil, err
	}
	if err = detectGoModule(&p); err != nil {
		return nil, err
	}

	p.ResourceDirs, err = detectResourceDirs(p.HomeDir)
	if err != nil {
		return nil, err
	}

	exists, err := protoFileExists(".")
	if err != nil {
		return nil, err
	}
	if exists {
		dir, _ := os.Getwd()
		p.CurrentResource = filepath.Base(dir)
	}

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

func detectResourceDirs(home string) ([]string, error) {
	fs, err := ioutil.ReadDir(home)
	if err != nil {
		return nil, err
	}
	// collect resource dir
	var dirs []string
	for _, info := range fs {
		if !info.IsDir() {
			continue
		}

		name := info.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}

		_, err := os.Stat(filepath.Join(home, name, name+".proto"))
		if err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
			continue
		}

		dirs = append(dirs, name)
	}
	return dirs, nil
}

func detectGoModule(c *Workspace) error {
	file, err := os.Open(filepath.Join(c.HomeDir, "go.mod"))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			parts := strings.Split(line, " ")
			c.Module = parts[1]
			return nil
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return fmt.Errorf("module not found in go.mod")
}

var ErrNotInWorkingDir = errors.New("tao works in project dir (has go.mod) or it's children resource dir (has .proto)")

// workingPath get home/resource path
func detectWorkingPath(c *Workspace) (err error) {
	var exists bool
	exists, err = tos.FileExists("./go.mod")
	if err != nil {
		return
	}
	if !exists {
		exists, err = tos.FileExists("../go.mod")
		if err != nil {
			return
		}

		if !exists {
			err = ErrNotInWorkingDir
			return
		}

		// we should in resource dir, ensure we have proto file
		//exists, err = protoFileExists(".")
		//if err != nil {
		//	return
		//}
		//if !exists {
		//	err = ErrNotInWorkingDir
		//	return
		//}
		c.HomeDir, _ = filepath.Abs("..")
	} else {
		c.HomeDir, _ = filepath.Abs(".")
	}
	return nil
}

func protoFileExists(dir string) (bool, error) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}

	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".proto") {
			return true, nil
		}
	}

	return false, nil
}
