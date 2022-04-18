package conversationer

import "github.com/nndergunov/tgBot/bot/pkg/domain/messenger"

func (c Conversationer) AddHome(msg messenger.SendMessage) messenger.SendMessage {
	msg.ReplyKeyboard = append(msg.ReplyKeyboard, []string{"Take me home!"})

	return msg
}
