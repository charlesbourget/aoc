package utils

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func CreateDir(path string) error {
	if FileExists(path) {
		return nil
	}

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func CreateFile(path string) (*os.File, error) {
	if FileExists(path) {
		return nil, fmt.Errorf("file %s already exists", path)
	}

	file, err := os.Create(path)
	if err != nil {
		return &os.File{}, err
	}

	return file, nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ReadTomlFile(filePath []byte, config interface{}) error {
	return toml.Unmarshal(filePath, config)
}
