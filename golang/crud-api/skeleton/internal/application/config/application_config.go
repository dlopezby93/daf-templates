package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ApplicationConfig struct {
	MongoURI string `env:"MONGO_URI"`
	Database string `env:"MONGO_DATABASE"`
}

func NewApplicationConfig() *ApplicationConfig {
	var config ApplicationConfig

	err := godotenv.Load()

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to load .env file")
	}

	err = env.Parse(&config)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse ennvironment variables")
	}
	return &config
}
