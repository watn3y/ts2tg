package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	LogLevel      int `env:"TS2TG_LOGLEVEL, default=1"`
	SleepInterval int `env:"TS2TG_SLEEPINTERVAL, default=60"`

	Telegram struct {
		APIToken string `env:"TS2TG_TELEGRAM_APITOKEN, required"`
		ChatID   int64  `env:"TS2TG_TELEGRAM_CHATID, required"`
	}

	Teamspeak struct {
		BaseURL string `env:"TS2TG_TEAMSPEAK_BASEURL, required"`
		APIKey  string `env:"TS2TG_TEAMSPEAK_APIKEY, required"`
		JoinURL string `env:"TS2TG_TEAMSPEAK_JOINURL"`
	}
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Info().Msg("No .env file found, relying on system environment variables")
	} else {
		log.Info().Msg(".env file loaded successfully")
	}

	var cfg Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config from env variables")
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.LogLevel))

	log.Info().Msg("Config loaded successfully")
	log.Debug().Interface("config", cfg).Send()

	return &cfg
}
