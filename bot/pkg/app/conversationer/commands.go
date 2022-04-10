package conversationer

import (
	"fmt"
	"log"

	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (c Conversationer) StartResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := fmt.Sprintf("%v, the bot had started successfully", msg.FirstName)

	keyboard := c.keyboards[StartKeyboardKey]

	return messenger.MakeKeyedTextMessage(msg.ChatID, text, nil, keyboard)
}

func (c Conversationer) HelpResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.Help

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) BotInfoResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	startTime, err := c.database.GetInfo()
	if err != nil {
		log.Println(err)
	}

	text := fmt.Sprintf("%s\nBot Start time: %v", c.answers.BotInfo, startTime.Starttime.Format(timeFormat))

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) ViewGenresResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	// TODO logic.
	return messenger.SendMessage{}
}

func (c Conversationer) ViewArtistsResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	// TODO logic.
	return messenger.SendMessage{}
}

func (c Conversationer) UnsupportedResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.Unsupported

	return messenger.MakeTextMessage(msg.ChatID, text)
}
