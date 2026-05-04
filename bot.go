package main

import (
	"strings"
	"time"

	"watn3y/ts2tg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func bot(tgClient *telegram.Client) {
	go setCommands(tgClient)

	for update := range tgClient.Updates {
		log.Trace().Interface("update", update).Msg("Update received")

		if update.Message == nil || update.Message.Text == "" {
			log.Debug().Int("UpdateID", update.UpdateID).Msg("Skipping non-message update")
			continue
		}
		if update.Message.Time().UTC().Unix()+60 < time.Now().UTC().Unix() {
			log.Debug().Int("UpdateID", update.UpdateID).Msg("Skipping old update")
			continue
		}

		if update.Message.IsCommand() {
			log.Info().Int64("ChatID", update.Message.Chat.ID).Int64("UserID", update.Message.From.ID).Str("Text", update.Message.Text).Msg("Command received")
			handleCommand(update, tgClient)
		}
	}
}

func setCommands(tgClient *telegram.Client) {
	github := tgbotapi.BotCommand{Command: "github", Description: "Source GitHub repo"}

	_, err := tgClient.Request(tgbotapi.NewSetMyCommands(github))
	if err != nil {
		log.Error().Err(err).Msg("Failed to publish commands to Telegram")
		return
	}

	log.Info().Msg("Published commands to Telegram successfully")
}

func handleCommand(update tgbotapi.Update, tgClient *telegram.Client) {
	cmd := strings.ToLower(update.Message.Command())

	switch cmd {
	case "start", "github":
		cmdGithub(update, tgClient)
	}
}

func cmdGithub(update tgbotapi.Update, tgClient *telegram.Client) {
	message := tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{ChatID: update.Message.Chat.ID, ReplyToMessageID: update.Message.MessageID},
		ParseMode:             "html",
		DisableWebPagePreview: false,
		Text:                  "Check out: https://github.com/watn3y/ts2tg",
	}

	result, err := tgClient.SendMessage(message)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
		return
	}

	log.Debug().Int64("chat", result.Chat.ID).Str("msg", result.Text).Msg("Sent message successfully")
	log.Trace().Interface("msg", result).Send()
}
