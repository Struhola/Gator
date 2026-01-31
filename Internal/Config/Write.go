package Config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Write(cfg *Config) error {
	//home, err := os.UserHomeDir()
	home, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(home, configFileName)

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
