package main

import (
	"github.com/nndergunov/tgBot/server/pkg/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	f, err := os.Create("log.log")
	if err != nil {
		panic(err)
	}

	logger := log.New(f, "bot", log.LstdFlags)

	if err = godotenv.Load(".env"); err != nil {
		logger.Printf("env file read: %v", err)
	}

	token := os.Getenv("apitoken")

	bot := api.ChatBot{}

	err = bot.Init(token, logger)
	if err != nil {
		logger.Println(err)
	}

	bot.Handle()
}