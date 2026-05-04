package main

import (
	"os"
	"strconv"
	"strings"
	"time"
	"watn3y/ts2tg/teamspeak"
	"watn3y/ts2tg/telegram"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configureLogger()
	log.Info().Msg("Starting ts2tg...")
	cfg := loadConfig()

	tgClient, err := telegram.NewClient(cfg.Telegram.APIToken)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to authenticate to Telegram")
	}
	log.Info().Int64("ID", tgClient.GetSelf().ID).Str("username", tgClient.GetSelf().UserName).Msg("Authenticated to Telegram successfully")
	go bot(tgClient)

	tsClient := teamspeak.NewClient(cfg.Teamspeak.BaseURL, cfg.Teamspeak.APIKey)
	serverInfo, err := tsClient.GetServerInfo()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to TeamSpeak WebQuery")
	}
	log.Info().Str("platform", serverInfo.Platform).Str("name", serverInfo.Name).Str("version", serverInfo.Version).Msg("Connected to TeamSpeak WebQuery successfully")

	monitor(tsClient, tgClient, cfg.SleepInterval, cfg.Telegram.ChatID, cfg.Teamspeak.JoinURL)
}

func configureLogger() {

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		const prefix = "ts2tg/"

		index := strings.Index(file, prefix)
		if index != -1 {
			return file[index+len(prefix):] + ":" + strconv.Itoa(line)
		}
		return file + ":" + strconv.Itoa(line)
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}

	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	log.Info().Msg("Logger started successfully")

}
