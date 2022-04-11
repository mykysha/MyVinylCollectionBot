package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

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
	msg := tgbotapi.NewPhoto(messasge.ChatID, tgbotapi.FileID(messasge.Photo.FileUniqueID))

	if messasge.Text != "" {
		msg.Caption = messasge.Text
	}

	if messasge.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(messasge.ReplyKeyboard)
	}

	if messasge.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(messasge.InlineKeyboard)
	}

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("SendTextMessage: %w", err)
	}

	return nil
}

func (tg TgBot) sendText(messasge messenger.SendMessage) error {
	msg := tgbotapi.NewMessage(messasge.ChatID, messasge.Text)

	if messasge.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(messasge.ReplyKeyboard)
	}

	if messasge.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(messasge.InlineKeyboard)
	}

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("SendTextMessage: %w", err)
	}

	return nil
}
