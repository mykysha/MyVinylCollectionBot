package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nndergunov/tgBot/server/api"
	db "github.com/nndergunov/tgBot/server/pkg/db/service"
	"github.com/nndergunov/tgBot/server/pkg/log"
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

	dbSource := fmt.Sprintf(
		"host=" + os.Getenv("HOST") +
			" port=" + os.Getenv("PORT") +
			" user=" + os.Getenv("USER") +
			" password=" + os.Getenv("PASS") +
			" dbname=" + os.Getenv("NAME") +
			" sslmode=" + os.Getenv("SSL"),
	)

	database, err := db.NewDB(dbSource)
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("APITOKEN")

	bot, err := api.NewChatBot(token, botLogger, database)
	if err != nil {
		l.Println(err)
	}

	bot.Handle()
}
