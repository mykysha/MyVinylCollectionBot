package api

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	logger "github.com/nndergunov/tgBot/bot/pkg/logger"
)

type ChatBot struct {
	bot          *tgbotapi.BotAPI
	log          *logger.Logger
	startButtons tgbotapi.ReplyKeyboardMarkup
	editPrompt   tgbotapi.InlineKeyboardMarkup
	viewPrompt   tgbotapi.InlineKeyboardMarkup
}

func NewChatBot(token string, logger *logger.Logger) (*ChatBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("connect %w", err)
	}

	chatBot := ChatBot{
		bot: bot,
		log: logger,
		startButtons: tgbotapi.ReplyKeyboardMarkup{
			Keyboard:              nil,
			ResizeKeyboard:        true,
			OneTimeKeyboard:       false,
			InputFieldPlaceholder: "Select a button",
			Selective:             false,
		},
		editPrompt: tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: nil,
		},
		viewPrompt: tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: nil,
		},
	}

	chatBot.generateButtons()
	chatBot.generateEditPrompt()
	chatBot.generateViewPrompt()

	return &chatBot, nil
}

func (b *ChatBot) generateButtons() {
	startButtons := tgbotapi.NewReplyKeyboard(
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

	b.startButtons = startButtons
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
