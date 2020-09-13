package generator

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	GoOutputDir    string
	SwiftOutputDir string
	DartOutputDir  string
	UseSnackCase   bool
	Dependencies   []string
}

func NewConfig(homeDir string) (*Config, error) {
	var conf Config

	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configFilePath := filepath.Join(userHome, ".tao/tao.json")

	err = load(&conf, configFilePath)
	if err != nil {
		return nil, err
	}
	// merge current working directory config
	configFilePath = filepath.Join(homeDir, "tao.json")
	err = load(&conf, configFilePath)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func load(conf *Config, filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	return json.NewDecoder(f).Decode(&conf)
}
