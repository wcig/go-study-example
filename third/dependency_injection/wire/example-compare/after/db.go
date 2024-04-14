package main

type DB struct {
	cfg *Config
}

func NewDB(cfg *Config) *DB {
	return &DB{
		cfg: cfg,
	}
}

func (db *DB) Ping() string {
	return "Pong"
}
