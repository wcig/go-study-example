package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/wire"
)

var Provider = wire.NewSet(NewConfig)

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	DSN string `json:"dsn"`
}

func NewConfig() (*Config, error) {
	file, err := os.Open("../config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg, nil
}
