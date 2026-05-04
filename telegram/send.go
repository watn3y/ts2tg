package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Client) SendMessage(message tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	return c.bot.Send(message)
}
