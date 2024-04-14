package config

import (
	"encoding/json"
	"go-app/third/dependency_injection/wire/example-project/pkg/jsonx"
	"os"
)

type Server struct {
	Address string `json:"address"`
	Mode    string `json:"mode"`
}

type Database struct {
	DriveName string `json:"drive_name"`
	Dsn       string `json:"dsn"`
	MaxIdle   int    `json:"max_idle"`
	MaxOpen   int    `json:"max_open"`
	ShowSql   bool   `json:"show_sql"`
}

type Data struct {
	Database *Database `json:"database"`
	Redis    *Redis    `json:"redis"`
}

type Redis struct {
	Address string `json:"address"`
}

type Logger struct {
	Path  string `json:"path"`
	Level string `json:"level"`
}

type Config struct {
	Server *Server `json:"server"`
	Data   Data    `json:"data"`
	Logger *Logger `json:"logger"`
}

func Load(filePath string) *Config {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		panic(err)
	}
	jsonx.PrintStr(&cfg)
	return &cfg
}
