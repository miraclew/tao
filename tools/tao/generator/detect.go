package generator

import (
	"bufio"
	"github.com/miraclew/tao/pkg/tos"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var ErrNotInWorkingDir = errors.New("tao works in project dir (has go.mod) or it's children resource dir (has .proto)")

// Detect the home dir of the go project (contains go.mod)
// Check current dir and parent dir
func DetectGoProjectHome(currentDir string) (string, error) {
	exists, err := tos.FileExists(filepath.Join(currentDir, "go.mod"))
	if err != nil {
		return "", nil
	}
	if !exists {
		exists, err = tos.FileExists(filepath.Join(currentDir, "../go.mod"))
		if err != nil {
			return "", nil
		}

		if !exists {
			return "", ErrNotInWorkingDir
		}

		dir, _ := filepath.Abs(filepath.Join(currentDir, ".."))
		return dir, nil
	} else {
		dir, _ := filepath.Abs(filepath.Join(currentDir))
		return dir, nil
	}
}

func ProtoFileExists(dir string) (bool, error) {
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

// Parse go.mod and extract go module
func DetectGoModule(homeDir string) (string, error) {
	file, err := os.Open(filepath.Join(homeDir, "go.mod"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			parts := strings.Split(line, " ")
			return parts[1], err
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module not found in go.mod")
}

// Detect service dirs
func DetectServiceDirs(home string) ([]string, error) {
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
