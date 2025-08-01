package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"packcalculator/config/constants/messages"
	"sync"
)

type Config struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port string `envconfig:"SERVICE_PORT" default:"80"`
}

var (
	once sync.Once
	Data *Config
)

func LoadConfig() *Config {
	once.Do(func() {
		//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC822})
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		logLevel, _ := zerolog.ParseLevel("debug")
		zerolog.SetGlobalLevel(logLevel)

		var err error
		err = godotenv.Load()
		if err != nil {
			fmt.Println(messages.ErrorloadingEnvFileMsg)
		}

		Data = &Config{}
		err = envconfig.Process("", Data)
		if err != nil {
			log.Fatal().Err(err).Msg(messages.FailedProcessConfigMsg)
		}
	})

	if Data == nil {
		log.Fatal().Msg(messages.ConfigNilErrorMsg)
	}

	return Data
}
