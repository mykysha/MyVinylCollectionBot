package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (tg TgBot) parseToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	switch {
	case update.Message.Photo != nil:
		return tg.photoToMessage(update)
	case update.Message.Voice != nil:
		return tg.voiceToMessage(update)
	case update.Message.VideoNote != nil:
		return tg.videoNoteToMessage(update)
	case update.Message.Video != nil:
		return tg.videoToMessage(update)
	case update.Message.Poll != nil:
		return tg.pollToMessage(update)
	default:
		return tg.textToMessage(update)
	}
}

func (tg TgBot) photoToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	bestQualityPhoto := len(update.Message.Photo) - 1

	photo := domain.NewPhoto(update.Message.Photo[bestQualityPhoto].FileUniqueID)

	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      update.Message.Caption,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     photo,
		Voice:     nil,
		VideoNote: nil,
		Video:     nil,
		Poll:      nil,
	}
}

func (tg TgBot) voiceToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	voice := domain.NewVoice(update.Message.Voice.FileUniqueID)

	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      update.Message.Caption,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     nil,
		Voice:     voice,
		VideoNote: nil,
		Video:     nil,
		Poll:      nil,
	}
}

func (tg TgBot) videoNoteToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	videoNote := domain.NewVideoNote(update.Message.VideoNote.FileUniqueID)

	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      "",
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     nil,
		Voice:     nil,
		VideoNote: videoNote,
		Video:     nil,
		Poll:      nil,
	}
}

func (tg TgBot) videoToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	video := domain.NewVideo(update.Message.Video.FileUniqueID)

	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      update.Message.Caption,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     nil,
		Voice:     nil,
		VideoNote: nil,
		Video:     video,
		Poll:      nil,
	}
}

func (tg TgBot) pollToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	tgOptions := update.Message.Poll.Options
	options := make([]string, 0, len(tgOptions))

	for _, option := range tgOptions {
		options = append(options, option.Text)
	}

	poll := domain.NewPoll(update.Message.Poll.Question, options)

	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      "",
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     nil,
		Voice:     nil,
		VideoNote: nil,
		Video:     nil,
		Poll:      poll,
	}
}

func (tg TgBot) textToMessage(update tgbotapi.Update) domain.ReceiveMessage {
	return domain.ReceiveMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      update.Message.Text,
		FirstName: update.Message.Chat.FirstName,
		LastName:  update.Message.Chat.LastName,
		UserName:  update.Message.Chat.UserName,
		Time:      update.Message.Time(),
		Photo:     nil,
		Voice:     nil,
		VideoNote: nil,
		Video:     nil,
		Poll:      nil,
	}
}
