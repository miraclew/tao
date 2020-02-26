package tos

import (
	"os"
)

func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func MustFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			panic(err)
		}
		return false
	}
	return true
}
