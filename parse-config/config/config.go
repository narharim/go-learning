package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig(path string) (*Config, error) {

	bytes, err := ReadConfig(path)
	if err != nil {
		return nil, err
	}

	config, err := ParseConfigFromJson(bytes)
	if err != nil {
		return nil, err
	}

	return config, err
}

func ParseConfigFromJson(bytes []byte) (*Config, error) {
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %v", err)
	}
	return &config, nil
}

func ReadConfig(path string) ([]byte, error) {
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}
	return config, err
}
