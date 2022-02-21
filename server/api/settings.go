package api

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/server/pkg/db/service"
	logger "github.com/nndergunov/tgBot/server/pkg/log"
)

type ChatBot struct {
	bot          *tgbotapi.BotAPI
	log          *logger.Logger
	db           *service.DB
	startButtons tgbotapi.ReplyKeyboardMarkup
	editPrompt   tgbotapi.InlineKeyboardMarkup
	viewPrompt   tgbotapi.InlineKeyboardMarkup
}

func NewChatBot(token string, l *logger.Logger, db *service.DB) (*ChatBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("connect %w", err)
	}

	b := ChatBot{
		bot:          bot,
		log:          l,
		db:           db,
		startButtons: tgbotapi.ReplyKeyboardMarkup{},
		editPrompt:   tgbotapi.InlineKeyboardMarkup{},
		viewPrompt:   tgbotapi.InlineKeyboardMarkup{},
	}

	b.generateButtons()
	b.generateEditPrompt()
	b.generateViewPrompt()

	return &b, nil
}

func (b *ChatBot) generateButtons() {
	startBttns := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("View collection"),
			tgbotapi.NewKeyboardButton("Edit collection"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("View genres"),
			tgbotapi.NewKeyboardButton("View artists"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("View wishlist"),
			tgbotapi.NewKeyboardButton("Edit wishlist"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Move from wishlist to collection"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Bot info"),
		),
	)

	b.startButtons = startBttns
}

func (b *ChatBot) generateEditPrompt() {
	editPrompt := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Add new", "Add new"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Delete", "Delete"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Back", "Back"),
		),
	)

	b.editPrompt = editPrompt
}

func (b *ChatBot) generateViewPrompt() {
	viewPrompt := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("View all full", "View all full"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("View full by choice", "View full by choice"),
		),
	)

	b.viewPrompt = viewPrompt
}
