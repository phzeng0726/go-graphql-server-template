package config

import (
	"os"
)

var (
	Env *AppConfig
)

const (
	keyEnv  = "ENV"
	keyHost = "HOST"
	keyPort = "PORT"
)

type AppConfig struct {
	Env  string
	Host string
	Port string
}

type DataPathConfig struct {
	LogFolder string
}

func InitConfig() {
	Env = &AppConfig{
		Env:  os.Getenv(keyEnv),
		Host: os.Getenv(keyHost),
		Port: os.Getenv(keyPort),
	}
}
