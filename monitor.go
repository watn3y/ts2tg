package main

import (
	"fmt"
	"strings"
	"time"
	"watn3y/ts2tg/teamspeak"
	"watn3y/ts2tg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func monitor(tsClient *teamspeak.Client, tgClient *telegram.Client, sleepInterval int, chatID int64, joinURL string) {

	previous, err := tsClient.GetClientListWithInfo()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get initial client list")
	}

	for {
		time.Sleep(time.Duration(sleepInterval) * time.Second)

		current, err := tsClient.GetClientListWithInfo()
		if err != nil {
			log.Error().Err(err).Msg("Failed to get client list")
			continue
		}

		prevBySession := make(map[string]teamspeak.ClientList, len(previous))
		for _, c := range previous {
			// Type 0 is a regular connection, Type 1 a serverquery connection
			if c.Type != "1" {
				prevBySession[c.SessionID] = c
			}
		}

		currBySession := make(map[string]teamspeak.ClientList, len(current))
		for _, c := range current {
			// Type 0 is a regular connection, Type 1 a serverquery connection
			if c.Type != "1" {
				currBySession[c.SessionID] = c
			}
		}

		for id, c := range prevBySession {
			if _, online := currBySession[id]; !online {
				log.Info().Str("nickname", c.Nickname).Msg("Client left")
				log.Debug().Str("nickname", c.Nickname).Str("databaseID", c.DatabaseID).Str("sessionID", c.SessionID).Str("platform", c.Platform).Send()
				log.Trace().Interface("client", c).Send()

				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("<b>%s</b> left", c.Nickname))
				msg.ParseMode = "html"
				result, err := tgClient.SendMessage(msg)
				if err != nil {
					log.Error().Err(err).Msg("Failed to send leave notification")
				} else {
					log.Debug().Str("nickname", c.Nickname).Int64("chatID", result.Chat.ID).Int("messageID", result.MessageID).Msg("Sent leave notification")
					log.Trace().Interface("message", result).Send()
				}
			}
		}

		for id, c := range currBySession {
			if _, online := prevBySession[id]; !online {
				log.Info().Str("nickname", c.Nickname).Msg("Client joined")
				log.Debug().Str("nickname", c.Nickname).Str("databaseID", c.DatabaseID).Str("sessionID", c.SessionID).Str("platform", c.Platform).Send()
				log.Trace().Interface("client", c).Send()

				version, _, _ := strings.Cut(c.Version, " [")
				text := fmt.Sprintf("<b>%s</b> joined using TeamSpeak <b>%s</b> for <b>%s</b>", c.Nickname, version, c.Platform)
				if joinURL != "" {
					text += fmt.Sprintf("\n<a href=\"%s\">Click to join them</a>", joinURL)
				}
				msg := tgbotapi.NewMessage(chatID, text)
				msg.ParseMode = "html"
				result, err := tgClient.SendMessage(msg)
				if err != nil {
					log.Error().Err(err).Msg("Failed to send join notification")
				} else {
					log.Debug().Str("nickname", c.Nickname).Int64("chatID", result.Chat.ID).Int("messageID", result.MessageID).Msg("Sent join notification")
					log.Trace().Interface("message", result).Send()
				}
			}
		}

		previous = current
	}
}
