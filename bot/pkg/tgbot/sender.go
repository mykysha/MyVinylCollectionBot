package tgbot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func stringToInt64(s string) (int64, error) {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("stringToInt64: %w", err)
	}

	return res, nil
}

func (tg TgBot) Send(message messenger.SendMessage) error {
	var err error

	if message.Photo != nil {
		err = tg.sendPhoto(message)
	} else {
		if message.File != nil {
			err = tg.sendFile(message)
		} else {
			err = tg.sendText(message)
		}
	}

	if err != nil {
		return fmt.Errorf("TgBot.Send: %w", err)
	}

	return nil
}

func (tg TgBot) sendPhoto(message messenger.SendMessage) error {
	id, err := stringToInt64(message.ChatID)
	if err != nil {
		return fmt.Errorf("sendPhoto: %w", err)
	}

	msg := tgbotapi.NewPhoto(id, tgbotapi.FileID(message.Photo.FileUniqueID))

	if message.Text != "" {
		msg.Caption = message.Text
	}

	if message.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(message.ReplyKeyboard)
	}

	if message.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(message.InlineKeyboard)
	}

	_, err = tg.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("sendPhoto: %w", err)
	}

	return nil
}

func (tg TgBot) sendText(message messenger.SendMessage) error {
	id, err := stringToInt64(message.ChatID)
	if err != nil {
		return fmt.Errorf("sendText: %w", err)
	}

	msg := tgbotapi.NewMessage(id, message.Text)

	if message.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(message.ReplyKeyboard)
	}

	if message.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(message.InlineKeyboard)
	}

	_, err = tg.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("sendText: %w", err)
	}

	return nil
}

func (tg TgBot) sendFile(message messenger.SendMessage) error {
	id, err := stringToInt64(message.ChatID)
	if err != nil {
		return fmt.Errorf("sendText: %w", err)
	}

	f := tgbotapi.FileReader{
		Name:   "collection.xlsx",
		Reader: message.File,
	}

	msg := tgbotapi.NewDocument(id, f)

	_, err = tg.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("sendFile: %w", err)
	}

	return nil
}
