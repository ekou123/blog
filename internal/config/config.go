package config

import (
	"encoding/json"
	"github.com/ekou123/blog/internal/database"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL string `json:"db_url"`
	User  string `json:"current_user"`
}

type State struct {
	Db  *database.Queries
	Cfg *Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handler map[string]func(*State, Command) error
}

const ConfigFileName = ".gatorconfig.json"

func Read() (Config, error) {

	fullPath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil

}

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, ConfigFileName)

	return fullPath, nil
}

func (cfg *Config) SetUser(username string) error {
	// Update the current user name
	cfg.User = username

	// Write the updated config to file
	return Write(*cfg)

}

func Write(cfg Config) error {

	fullPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
