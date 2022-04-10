package main

import (
	"fmt"
	"os"

	"github.com/nndergunov/tgBot/bot/api"
	"github.com/nndergunov/tgBot/bot/pkg/app/conversationer"
	"github.com/nndergunov/tgBot/bot/pkg/app/msgrouter"
	"github.com/nndergunov/tgBot/bot/pkg/configreader"
	"github.com/nndergunov/tgBot/bot/pkg/domain"
	"github.com/nndergunov/tgBot/bot/pkg/logger"
	"github.com/nndergunov/tgBot/bot/pkg/pooler"
	"github.com/nndergunov/tgBot/bot/pkg/tgbot"
)

func main() {
	mainLogger := logger.NewLogger(os.Stdout, "main")

	err := configreader.SetConfigFile("config.yaml")
	if err != nil {
		mainLogger.Panicln(err)
	}

	bot, err := createBot()
	if err != nil {
		mainLogger.Panicln(err)
	}

	communicator := createCommunicator()

	router := msgrouter.NewMsgRouter(communicator)

	workerNumber := 16
	workerpool := pooler.NewPool(workerNumber)

	logFile, err := os.Create("log.log")
	if err != nil {
		mainLogger.Panicln("error", err)
	}

	apiLogger := logger.NewLogger(logFile, "botAPI")

	botConfig := api.Configuration{
		Pool:   workerpool,
		Router: router,
		Bot:    bot,
		Logger: apiLogger,
	}

	botAPI, err := api.NewChatBot(botConfig)
	if err != nil {
		mainLogger.Println(err)
	}

	botAPI.Handle()
}

func createBot() (*tgbot.TgBot, error) {
	token := configreader.GetString("bot.token")

	bot, err := tgbot.NewTelegramBot(token)
	if err != nil {
		return nil, fmt.Errorf("createBot: %w", err)
	}

	return bot, nil
}

func createCommunicator() *conversationer.Conversationer {
	answers := getAnswers()
	keyboards := getKeyboards()

	comm := conversationer.NewConver(answers, keyboards)

	return comm
}

func getAnswers() domain.Answers {
	return domain.Answers{
		Help:      configreader.GetString("answers.help"),
		Voice:     configreader.GetString("answers.voice"),
		VideoNote: configreader.GetString("answers.videoNote"),
		Photo:     configreader.GetString("answers.photo"),
		Video:     configreader.GetString("answers.video"),
		Poll:      configreader.GetString("answers.poll"),
		Unknown:   configreader.GetString("answers.unknown"),
		BotInfo:   configreader.GetString("answers.botInfo"),
	}
}

func getKeyboards() map[string][][]string {
	keyboardMap := make(map[string][][]string)

	keyboardMap[conversationer.StartKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.startKeyboard")
	keyboardMap[conversationer.EditKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.editKeyboard")
	keyboardMap[conversationer.ViewKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.viewKeyboard")

	return keyboardMap
}
