package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"packcalculator/config/constants/messages"
	"sync"
)

type Config struct {
	Debug bool   `envconfig:"SERVICE_DEBUG" default:"false"`
	Host  string `envconfig:"SERVICE_HOST" default:"0.0.0.0"`
	Port  string `envconfig:"SERVICE_PORT" default:"8180"`
}

var (
	once sync.Once
	Data *Config
)

func LoadConfig() *Config {
	once.Do(func() {
		Data = &Config{}

		// We structure log output
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		logLevel, _ := zerolog.ParseLevel("debug")
		zerolog.SetGlobalLevel(logLevel)

		// This will load .env file and set env values
		_ = godotenv.Load()
		if err := envconfig.Process("", Data); err != nil {
			log.Fatal().Err(err).Msg(messages.FailedProcessConfigMsg)
		}
	})

	if Data == nil {
		log.Fatal().Msg(messages.ConfigNilErrorMsg)
	}

	return Data
}
