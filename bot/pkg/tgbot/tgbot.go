package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

type TgBot struct {
	MessageChan chan messenger.ReceiveMessage
	bot         *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("NewTelegramBot: %w", err)
	}

	messageChan := make(chan messenger.ReceiveMessage)

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
