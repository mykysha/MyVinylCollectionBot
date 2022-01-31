package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ChatBot struct {
	bot *tgbotapi.BotAPI
	log *log.Logger
}

func (b *ChatBot) Init(token string, l *log.Logger) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("connect %w", err)
	}

	b.bot = bot
	b.log = l

	return nil
}

func (b *ChatBot) Handle() {
	b.bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.log.Println(update.Message)

		switch update.Message.Command() {
		case "start":
			go b.startHandler(update)
		case "help":
			go b.helpHandler(update)
		default:
			go b.otherHandler(update)
		}
	}
}

func (b ChatBot) startHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("%v, I am alive!", update.Message.Chat.FirstName)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}

func (b ChatBot) helpHandler(update tgbotapi.Update) {
	text := "insert help"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}

func (b ChatBot) otherHandler(update tgbotapi.Update) {
	text := fmt.Sprintf("%v, you've written '%v'", update.Message.Chat.FirstName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := b.bot.Send(msg); err != nil {
		b.log.Println(err)
	}
}