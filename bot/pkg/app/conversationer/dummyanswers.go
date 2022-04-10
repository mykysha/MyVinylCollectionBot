package conversationer

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (c Conversationer) VoiceResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.Voice

	return domain.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) VideoNoteResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.VideoNote

	return domain.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) PhotoResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.Photo

	return domain.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) VideoResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.VideoNote

	return domain.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) PollResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.Poll

	return domain.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) UnknownTypeResponser(msg domain.ReceiveMessage) domain.SendMessage {
	text := c.answers.Unknown

	return domain.MakeTextMessage(msg.ChatID, text)
}
