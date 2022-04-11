package conversationer

import (
	"fmt"
	"strings"

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
		text := "Some error working with database, try again later"

		return messenger.MakeTextMessage(msg.ChatID, text)
	}

	text := fmt.Sprintf("%s\nBot Start time: %v", c.answers.BotInfo, startTime.Starttime.Format(timeFormat))

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) ViewGenresResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	genres, err := c.database.GetGenres(int(msg.ChatID))
	if err != nil {
		text := "Some error working with database, try again later"

		return messenger.MakeTextMessage(msg.ChatID, text)
	}

	var text string

	for i := 0; i < len(genres); i++ {
		text += fmt.Sprintf("%d. %s\n", i+1, genres[i])
	}

	text = strings.TrimRight(text, "\n")

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) ViewArtistsResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	artists, err := c.database.GetArtists(int(msg.ChatID))
	if err != nil {
		text := "Some error working with database, try again later"

		return messenger.MakeTextMessage(msg.ChatID, text)
	}

	var text string

	for i := 0; i < len(artists); i++ {
		text += fmt.Sprintf("%d. %s\n", i+1, artists[i].Name)
	}

	text = strings.TrimRight(text, "\n")

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) BackResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := "Ok, returning you back"

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) TakeHomeResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := "Ok, returning you to start menu"

	return messenger.MakeTextMessage(msg.ChatID, text)
}
