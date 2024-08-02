package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	App struct {
		Name string `env:"NAME"`
		Env  string `env:"ENV"`
		Port int    `env:"PORT"`
	}

	Database struct {
		Host     string `env:"DATABASE_HOST"`
		Port     string `env:"DATABASE_PORT"`
		Username string `env:"DATABASE_USERNAME"`
		Password string `env:"DATABASE_PASSWORD"`
		Name     string `env:"DATABASE_NAME"`
		Dsn      string `env:"DATABASE_DSN,expand"`
	}

	Config struct {
		App      App
		Database Database
	}
)

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}
	opts := env.Options{RequiredIfNoDef: true}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}
