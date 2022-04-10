package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func newInlineKeyboard(keyboard [][]string) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, len(keyboard))

	for _, row := range keyboard {
		keyboardRow := make([]tgbotapi.InlineKeyboardButton, 0, len(row))

		for _, button := range row {
			keyboardButton := tgbotapi.NewInlineKeyboardButtonData(button, button)

			keyboardRow = append(keyboardRow, keyboardButton)
		}

		rows = append(rows, keyboardRow)
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func newReplyKeyboard(keyboard [][]string) tgbotapi.ReplyKeyboardMarkup {
	rows := make([][]tgbotapi.KeyboardButton, 0, len(keyboard))

	for _, row := range keyboard {
		keyboardRow := make([]tgbotapi.KeyboardButton, 0, len(row))

		for _, button := range row {
			keyboardButton := tgbotapi.NewKeyboardButton(button)

			keyboardRow = append(keyboardRow, keyboardButton)
		}

		rows = append(rows, keyboardRow)
	}

	return tgbotapi.NewOneTimeReplyKeyboard(rows...)
}
