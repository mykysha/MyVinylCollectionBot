package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

type TgBot struct {
	MessageChan chan domain.ReceiveMessage
	bot         *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("NewTelegramBot: %w", err)
	}

	messageChan := make(chan domain.ReceiveMessage)

	return &TgBot{
		MessageChan: messageChan,
		bot:         bot,
	}, nil
}

func (tg *TgBot) Listen() {
	tg.bot.Debug = true

	go tg.getUpdates()
}

func (tg *TgBot) getUpdates() {
	updateConfig := tgbotapi.NewUpdate(0)

	updates := tg.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := tg.parseToMessage(update)

		tg.MessageChan <- message
	}
}

func (tg TgBot) parseToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      update.Message.Text,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     nil,
		Voice:     nil,
		VideoNote: nil,
		Video:     nil,
		Poll:      nil,
	}
}
