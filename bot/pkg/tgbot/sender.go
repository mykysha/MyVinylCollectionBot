package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (tg TgBot) Send(messasge domain.SendMessage) error {
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

func (tg TgBot) sendPhoto(messasge domain.SendMessage) error {
	msg := tgbotapi.NewPhoto(messasge.ChatID, tgbotapi.FileID(messasge.Photo.FileUniqueID))

	if messasge.Text != "" {
		msg.Caption = messasge.Text
	}

	if messasge.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(messasge.ReplyKeyboard)
	}

	if messasge.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(messasge.ReplyKeyboard)
	}

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("SendTextMessage: %w", err)
	}

	return nil
}

func (tg TgBot) sendText(messasge domain.SendMessage) error {
	msg := tgbotapi.NewMessage(messasge.ChatID, messasge.Text)

	if messasge.ReplyKeyboard != nil {
		msg.ReplyMarkup = newReplyKeyboard(messasge.ReplyKeyboard)
	}

	if messasge.InlineKeyboard != nil {
		msg.ReplyMarkup = newInlineKeyboard(messasge.ReplyKeyboard)
	}

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("SendTextMessage: %w", err)
	}

	return nil
}
