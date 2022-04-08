package api

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b ChatBot) voiceHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've received a voice message, " +
		"but I don't have ears to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}

func (b ChatBot) videoNoteHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've received a video message, " +
		"but I don't have ears or eyes to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}

func (b ChatBot) photoHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've received a picture, " +
		"but I don't have eyes to fully enjoy it. " +
		"However, I can add a photo as an album cover, but first, you should ask me to.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}

func (b ChatBot) videoHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've received a video, " +
		"but I dont't have ears nor eyes to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}

func (b ChatBot) pollHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've received a poll, " +
		"but due to intergalactic convention of bot rights, " +
		"I have no right to vote.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}
