package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	logger "github.com/nndergunov/tgBot/server/pkg/log"
)

type ChatBot struct {
	bot          *tgbotapi.BotAPI
	log          *logger.Logger
	startButtons tgbotapi.ReplyKeyboardMarkup
	editPrompt   tgbotapi.InlineKeyboardMarkup
	viewPrompt   tgbotapi.InlineKeyboardMarkup
	editOptions  tgbotapi.InlineKeyboardMarkup
}

func (b *ChatBot) Init(token string, l *logger.Logger) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("connect %w", err)
	}

	b.bot = bot
	b.log = l

	b.generateButtons()

	return nil
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