package main

import (
	"os"

	"github.com/nndergunov/tgBot/bot/api"
	"github.com/nndergunov/tgBot/bot/pkg/configreader"
	"github.com/nndergunov/tgBot/bot/pkg/logger"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	logFile, err := os.Create("log.log")
	if err != nil {
		mainLogger.Panicln("error", err)
	}

	botLogger := logger.NewLogger(logFile, "bot")

	err = configreader.SetConfigFile("config.yaml")
	if err != nil {
		mainLogger.Panicln(err)
	}

	token := configreader.GetString("bot.token")

	bot, err := api.NewChatBot(token, botLogger)
	if err != nil {
		mainLogger.Println(err)
	}

	bot.Handle()
}
