package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b ChatBot) voiceHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've recieved a voice message, " +
		"but I dont't have ears to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func (b ChatBot) videoNoteHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've recieved a video message, " +
		"but I dont't have ears or eyes to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func (b ChatBot) photoHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've recieved a picture, " +
		"but I don't have eyes to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}

	photoID := update.Message.Photo

	photo := tgbotapi.FileID(photoID[len(photoID)-1].FileUniqueID)

	photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, photo)

	photoMsg.Caption = update.Message.Caption

	if _, err := b.bot.Send(photoMsg); err != nil {
		log.Panic(err)
	}
}

func (b ChatBot) videoHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've recieved a video, " +
		"but I dont't have ears or eyes to fully enjoy it.")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func (b ChatBot) pollHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("I've recieved a poll, " +
		"but due to intergalactic convention of bot rights, " +
		"I have no right to vote")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}