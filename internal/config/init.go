package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type config struct {
	Host         string
	Port         int
	DatabasePath string
	Password     string
	Username     string
}

var CONFIG config

func Init() error {
	CONFIG = config{
		Host:         "localhost",
		Port:         8080,
		DatabasePath: "./poultracker.db",
		Username:     "admin",
		Password:     "password",
	}

	opts := env.Options{
		UseFieldNameByDefault: true,
		Prefix:                "POULTRACKER_",
	}

	if err := env.ParseWithOptions(&CONFIG, opts); err != nil {
		return err
	}

	log.Println("Loaded config")

	return nil
}

func GetConfig() config {
	return CONFIG
}

func Username() string {
	return CONFIG.Username
}

func Password() string {
	return CONFIG.Password
}

func DatabasePath() string {
	return CONFIG.DatabasePath
}

func Host() string {
	return CONFIG.Host
}

func Port() int {
	return CONFIG.Port
}
