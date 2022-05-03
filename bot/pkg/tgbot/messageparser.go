package tgbot

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func (tg TgBot) parseToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	switch {
	case update.CallbackQuery != nil:
		return tg.callbackMessage(update)
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

func (tg TgBot) photoToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	bestQualityPhoto := len(update.Message.Photo) - 1

	photo := messenger.NewPhoto(update.Message.Photo[bestQualityPhoto].FileID)

	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.Message.Chat.ID),
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

func (tg TgBot) voiceToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	voice := messenger.NewVoice(update.Message.Voice.FileUniqueID)

	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.Message.Chat.ID),
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

func (tg TgBot) videoNoteToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	videoNote := messenger.NewVideoNote(update.Message.VideoNote.FileUniqueID)

	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.Message.Chat.ID),
		Text:      update.Message.Caption,
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

func (tg TgBot) videoToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	video := messenger.NewVideo(update.Message.Video.FileUniqueID)

	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.Message.Chat.ID),
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

func (tg TgBot) pollToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	tgOptions := update.Message.Poll.Options
	options := make([]string, 0, len(tgOptions))

	for _, option := range tgOptions {
		options = append(options, option.Text)
	}

	poll := messenger.NewPoll(update.Message.Poll.Question, options)

	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.Message.Chat.ID),
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

func (tg TgBot) textToMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.Message.Chat.ID),
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

func (tg TgBot) callbackMessage(update tgbotapi.Update) messenger.ReceiveMessage {
	return messenger.ReceiveMessage{
		ChatID:    int64ToString(update.CallbackQuery.From.ID),
		Text:      update.CallbackQuery.Data,
		FirstName: update.CallbackQuery.From.FirstName,
		LastName:  update.CallbackQuery.From.LastName,
		UserName:  update.CallbackQuery.From.UserName,
		Time:      update.CallbackQuery.Message.Time(),
		Photo:     nil,
		Voice:     nil,
		VideoNote: nil,
		Video:     nil,
		Poll:      nil,
	}
}
