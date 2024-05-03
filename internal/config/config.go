package config

import (
	"os"
)

type Config struct {
	Token string
}

func MustLoad() *Config {
	var config Config
	config.Token = os.Getenv("BOT_TOKEN")
	return &config
}
