package config

import (
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	//home, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, configFileName)

	return path, nil
}
