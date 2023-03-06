package Telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//msg.Text = "Невалиднаы сслыка!"
//msg.Text = "Ты не авторизирован"
//msg.Text = "Не удалось созранить ссылку"

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user us  not authorized")
	errUnableToSave = errors.New("Unamle to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Default)
	switch err {
	case errInvalidURL:
		msg.Text = b.messages.InvalidURL
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "Ты не авторизирован!"
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = b.messages.UnableToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
