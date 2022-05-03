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

func (tg TgBot) Send(messasge messenger.SendMessage) error {
	var err error

	if messasge.Photo != nil {
		err = tg.sendPhoto(messasge)
	} else {
		err = tg.sendText(messasge)
	}

	if err != nil {
		return fmt.Errorf("TgBot.Send: %w", err)
	}

	return nil
}

func (tg TgBot) sendPhoto(messasge messenger.SendMessage) error {
	id, err := stringToInt64(messasge.ChatID)
	if err != nil {
		return fmt.Errorf("sendPhoto: %w", err)
	}

	msg := tgbotapi.NewPhoto(id, tgbotapi.FileID(messasge.Photo.FileUniqueID))

	if messasge.Text != "" {
		msg.Caption = messasge.Text
	}

	if messasge.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(messasge.ReplyKeyboard)
	}

	if messasge.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(messasge.InlineKeyboard)
	}

	_, err = tg.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("sendPhoto: %w", err)
	}

	return nil
}

func (tg TgBot) sendText(messasge messenger.SendMessage) error {
	id, err := stringToInt64(messasge.ChatID)
	if err != nil {
		return fmt.Errorf("sendText: %w", err)
	}

	msg := tgbotapi.NewMessage(id, messasge.Text)

	if messasge.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(messasge.ReplyKeyboard)
	}

	if messasge.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(messasge.InlineKeyboard)
	}

	_, err = tg.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("sendText: %w", err)
	}

	return nil
}
