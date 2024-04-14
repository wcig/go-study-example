package main

type Config struct {
	DataSource string
}

func NewConfig() *Config {
	return &Config{
		DataSource: "root:123456@tcp(127.0.0.1:3306)/test",
	}
}
