package config

import (
	"os"
)

type Config struct {
	App struct {
		Port string
	}
}

func New() *Config {
	c := new(Config)
	c.loadApp()

	return c
}

func (c *Config) loadApp() *Config {
	port := os.Getenv("PORT")

	c.App.Port = port

	return c
}
