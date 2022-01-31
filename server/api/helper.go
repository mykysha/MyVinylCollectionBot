package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b ChatBot) logMessage(msg *tgbotapi.Message) {
	logInfo := b.log.NewLog(msg.Chat.ID, msg.Chat.FirstName, msg.Chat.LastName, msg.Chat.UserName, msg.Time())

	switch {
	case msg.Voice != nil:
		b.log.VoiceLog(logInfo, msg.Voice.Duration)
	case msg.VideoNote != nil:
		b.log.VideoNoteLog(logInfo, msg.VideoNote.Duration)
	case msg.Photo != nil:
		b.log.PhotoLog(logInfo, msg.Photo[len(msg.Photo)-1].Height, msg.Photo[len(msg.Photo)-1].Width)
	case msg.Video != nil:
		b.log.VideoLog(logInfo, msg.Video.Duration, msg.Video.Height, msg.Video.Width)
	case msg.Poll != nil:
		b.log.PollLog(logInfo, msg.Poll.Question)
	case msg.Text != "":
		b.log.TextLog(logInfo, msg.Text)
	default:
		b.log.DefaultLog(logInfo)
	}
}