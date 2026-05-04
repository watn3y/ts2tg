package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewClient(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	updates := tgbotapi.NewUpdate(0)
	updates.Timeout = 60

	return &Client{
		bot:     bot,
		Updates: bot.GetUpdatesChan(updates),
	}, nil
}

func (c *Client) GetSelf() tgbotapi.User {
	return c.bot.Self
}

func (c *Client) Request(config tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	return c.bot.Request(config)
}
