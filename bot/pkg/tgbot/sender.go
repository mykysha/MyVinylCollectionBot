package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (tg TgBot) Send(messasge domain.SendMessage) error {
	msg := tgbotapi.NewMessage(messasge.ChatID, messasge.Text)

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("SendTextMessage: %w", err)
	}

	return nil
}

func (tg TgBot) sendTextMessage(toID int, messasge domain.SendMessage) error {
	msg := tgbotapi.NewMessage(int64(toID), messasge.Text)

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("SendTextMessage: %w", err)
	}

	return nil
}

func (tg TgBot) sendTextWithInlineKeyboard(toID int, message domain.SendMessage, keyboard [][]string) error {
	kbrd := newInlineKeyboard(keyboard)

	if err := tg.sendWithKeyboard(toID, message, kbrd); err != nil {
		return fmt.Errorf("SendTextWithInlineKeyboard: %w", err)
	}

	return nil
}

func (tg TgBot) sendTextWithReplyKeyboard(toID int, message domain.SendMessage, keyboard [][]string) error {
	kbrd := newReplyKeyboard(keyboard)

	if err := tg.sendWithKeyboard(toID, message, kbrd); err != nil {
		return fmt.Errorf("SendTextWithInlineKeyboard: %w", err)
	}

	return nil
}

func (tg TgBot) sendWithKeyboard(toID int, message domain.SendMessage, keyboard interface{}) error {
	msg := tgbotapi.NewMessage(int64(toID), message.Text)

	msg.ReplyMarkup = keyboard

	if _, err := tg.bot.Send(msg); err != nil {
		return fmt.Errorf("sendWithKeyboard: %w", err)
	}

	return nil
}
