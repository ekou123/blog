package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL string `json:"db_url"`
	User  string `json:"current_user"`
}

const ConfigFileName = ".gatorconfig.json"

func Read() (Config, error) {

	config := Config{}

	filePath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("cannot get config file path: %v", err)
	}

	configFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("cannot open config file: %v", err)
	}

	defer configFile.Close()

	var data []byte
	data, err = io.ReadAll(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("cannot read config file: %v", err)
	}

	err = json.Unmarshal(data, &config)

	return config, nil

}

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get home dir: %v", err)
	}

	return filepath.Join(homeDir, ConfigFileName), nil
}

func SetUser(cfg Config) error {
	currentConfig, err := Read()
	if err != nil {
		return fmt.Errorf("cannot read config file: %v", err)
	}

	if currentConfig.DbURL == "" {
		currentConfig.DbURL = "postgres://example"
	}

	currentConfig.User = cfg.User

	err = Write(currentConfig)
	if err != nil {
		return fmt.Errorf("cannot write config file: %v", err)
	}

	return nil
}

func Write(cfg Config) error {

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("cannot marshal current config: %v", err)
	}

	filePath, err := GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("cannot get config file path: %v", err)
	}

	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		return fmt.Errorf("cannot write config file: %v", err)
	}

	return nil
}
