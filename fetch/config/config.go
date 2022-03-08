package config

import (
	"os"
)

type (
	// Configuration is application configuration
	Configuration struct {
		AppName   string `json:"app_name"`
		JWTSecret string `json:"jwt_secret"`
		Port      string `json:"port"`
	}
)

// LoadConfiguration is a function to load all application configuration
func LoadConfiguration(appName string) *Configuration {
	return &Configuration{
		AppName:   appName,
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port:      ":" + os.Getenv("PORT"),
	}
}
