package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/server/pkg/domain"
)

func (b ChatBot) decodeMessage(apiMsg *tgbotapi.Message) {
	logInfo := b.log.NewLog(apiMsg.Chat.ID, apiMsg.Chat.FirstName, apiMsg.Chat.LastName, apiMsg.Chat.UserName, apiMsg.Time())

	switch {
	case apiMsg.Voice != nil:
		b.log.VoiceLog(logInfo, apiMsg.Voice.Duration)
	case apiMsg.VideoNote != nil:
		b.log.VideoNoteLog(logInfo, apiMsg.VideoNote.Duration)
	case apiMsg.Photo != nil:
		b.log.PhotoLog(logInfo, apiMsg.Photo[len(apiMsg.Photo)-1].Height, apiMsg.Photo[len(apiMsg.Photo)-1].Width)
	case apiMsg.Video != nil:
		b.log.VideoLog(logInfo, apiMsg.Video.Duration, apiMsg.Video.Height, apiMsg.Video.Width)
	case apiMsg.Poll != nil:
		b.log.PollLog(logInfo, apiMsg.Poll.Question)
	case apiMsg.Text != "":
		msg := domain.NewTextMessage(apiMsg.Chat.ID, apiMsg.Chat.FirstName, apiMsg.Chat.LastName,
			apiMsg.Chat.UserName, apiMsg.Text, apiMsg.Time())
		b.log.TextLog(msg)
	default:
		b.log.DefaultLog(logInfo)
	}
}

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
	default:
		b.log.DefaultLog(logInfo)
	}
}
