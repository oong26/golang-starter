package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	// server
	GoMode     string `mapstructure:"GO_MODE"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	// keys
	GeminiApiKey string `mapstructure:"GEMINI_API_KEY"`
}

func NewEnv(filename string, override bool) *Env {
	env := Env{}
	viper.SetConfigFile(filename)

	if override {
		viper.AutomaticEnv()
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading environment file", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error loading environment file", err)
	}

	return &env
}
