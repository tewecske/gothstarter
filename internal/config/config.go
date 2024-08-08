package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := Config{
		Port: os.Getenv("LISTEN_ADDR"),
	}

	return &cfg
}
