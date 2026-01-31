package Config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	var cfg Config

	path, err := GetConfigFilePath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
