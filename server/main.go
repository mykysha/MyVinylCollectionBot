package main

import (
	"github.com/nndergunov/tgBot/server/api"
	"github.com/nndergunov/tgBot/server/cmd/log"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	f, err := os.Create("log.log")
	if err != nil {
		panic(err)
	}

	l := log.New(f, "bot ", log.LstdFlags)

	botLogger := &logger.Logger{}

	botLogger.Init(l)

	if err = godotenv.Load(".env"); err != nil {
		l.Printf("env file read: %v", err)
	}

	token := os.Getenv("apitoken")

	bot := api.ChatBot{}

	err = bot.Init(token, botLogger)
	if err != nil {
		l.Println(err)
	}

	bot.Handle()
}