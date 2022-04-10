package conversationer

import (
	"fmt"

	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (c Conversationer) StartResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := fmt.Sprintf("%v, the bot had started successfully", msg.FirstName)

	keyboard := c.keyboards[StartKeyboardKey]

	return domain.MakeTextKeyboardMessage(msg.ChatID, text, nil, keyboard)
}

func (c Conversationer) HelpResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.Help

	return domain.MakeTextMessage(msg.ChatID, text)
}
