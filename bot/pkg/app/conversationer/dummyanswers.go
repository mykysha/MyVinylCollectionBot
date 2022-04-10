package conversationer

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (c Conversationer) VoiceResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.Voice

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) VideoNoteResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.VideoNote

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) PhotoResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.Photo

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) VideoResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.VideoNote

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) PollResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.Poll

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) UnknownTypeResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := c.answers.Unknown

	return messenger.MakeTextMessage(msg.ChatID, text)
}
