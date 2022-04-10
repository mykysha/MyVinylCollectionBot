package api

import (
	"github.com/nndergunov/tgBot/bot/pkg/app/msgrouter"
	"github.com/nndergunov/tgBot/bot/pkg/logger"
	pool "github.com/nndergunov/tgBot/bot/pkg/pooler"
	"github.com/nndergunov/tgBot/bot/pkg/tgbot"
)

type ChatBot struct {
	pool   *pool.Pool
	router *msgrouter.MsgRouter
	bot    *tgbot.TgBot
	log    *logger.Logger
}

type Configuration struct {
	Pool   *pool.Pool
	Router *msgrouter.MsgRouter
	Bot    *tgbot.TgBot
	Logger *logger.Logger
}

func NewChatBot(configuration Configuration) (*ChatBot, error) {
	chatBot := ChatBot{
		pool:   configuration.Pool,
		router: configuration.Router,
		bot:    configuration.Bot,
		log:    configuration.Logger,
	}

	return &chatBot, nil
}
