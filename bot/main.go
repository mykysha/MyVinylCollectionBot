package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nndergunov/tgBot/bot/api"
	"github.com/nndergunov/tgBot/bot/pkg/app/conversationer"
	"github.com/nndergunov/tgBot/bot/pkg/app/msgrouter"
	"github.com/nndergunov/tgBot/bot/pkg/configreader"
	"github.com/nndergunov/tgBot/bot/pkg/db"
	"github.com/nndergunov/tgBot/bot/pkg/domain/answerer"
	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
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

	communicator, err := createCommunicator()
	if err != nil {
		mainLogger.Panicln(err)
	}

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

func createCommunicator() (*conversationer.Conversationer, error) {
	dbURL := fmt.Sprintf(
		"host=" + configreader.GetString("database.host") +
			" port=" + configreader.GetString("database.port") +
			" user=" + configreader.GetString("database.user") +
			" password=" + configreader.GetString("database.password") +
			" dbname=" + configreader.GetString("database.dbname") +
			" sslmode=" + configreader.GetString("database.ssl"),
	)

	database, err := db.NewDatabase(dbURL)
	if err != nil {
		return nil, fmt.Errorf("createCommunicator: %w", err)
	}

	err = database.PutInfo(entities.Info{Starttime: time.Now()})
	if err != nil {
		return nil, fmt.Errorf("createCommunicator: %w", err)
	}

	answers := getAnswers()
	keyboards := getKeyboards()

	comm := conversationer.NewConver(database, answers, keyboards)

	return comm, nil
}

func getAnswers() answerer.Answers {
	return answerer.Answers{
		Help:        configreader.GetString("answers.help"),
		Voice:       configreader.GetString("answers.voice"),
		VideoNote:   configreader.GetString("answers.videoNote"),
		Photo:       configreader.GetString("answers.photo"),
		Video:       configreader.GetString("answers.video"),
		Poll:        configreader.GetString("answers.poll"),
		Unsupported: configreader.GetString("answers.unsupported"),
		Unknown:     configreader.GetString("answers.unknown"),
		BotInfo:     configreader.GetString("answers.botInfo"),
	}
}

func getKeyboards() map[string][][]string {
	keyboardMap := make(map[string][][]string)

	keyboardMap[conversationer.StartKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.startKeyboard")
	keyboardMap[conversationer.EditKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.editKeyboard")
	keyboardMap[conversationer.ViewKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.viewKeyboard")
	keyboardMap[conversationer.ExpandedViewKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.expandedView")
	keyboardMap[conversationer.EditViewKeyboardKey] = configreader.GetStringSliceOfSlices("keyboards.editView")

	return keyboardMap
}
