package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path             string `yaml:"path"`
	Country          string `yaml:"country"`
	XL               bool   `yaml:"xl"`
	DeleteOldVersion bool   `yaml:"delete_old_version"`
}

const configFileName = "config.yml"

func loadConfig() Config {
	config := Config{
		Path:             "",
		Country:          "us",
		XL:               true,
		DeleteOldVersion: true,
	}

	data, err := os.ReadFile(configFileName)
	if err != nil {
		fmt.Println("Config file not found or unreadable, using default values.")
		return config
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing config file, using default values.")
		return config
	}

	validCountries := map[string]bool{"us": true, "es": true, "d": true, "br": true}
	if _, ok := validCountries[config.Country]; !ok {
		fmt.Println("Invalid country in config, defaulting to 'us'.")
		config.Country = "us"
	}

	return config
}

func getHabboExe(config Config) string {
	if config.XL {
		return fmt.Sprintf("HabboHotel-o%s-xl.exe", config.Country)
	}
	return fmt.Sprintf("HabboHotel-o%s.exe", config.Country)
}
