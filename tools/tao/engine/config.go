package engine

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DartOutputDir string
}

func NewConfig() (*Config, error) {
	configFilePath := filepath.Join(os.Getenv("HOME"), ".tao/tao.json")
	_, err := os.Stat(configFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return nil, nil
	}

	f, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
