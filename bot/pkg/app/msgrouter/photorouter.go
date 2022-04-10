package msgrouter

import "github.com/nndergunov/tgBot/bot/pkg/domain"

func (r MsgRouter) photoRouter(msg domain.ReceiveMessage) domain.SendMessage {
	return domain.MakePhotoMessage(msg.ChatID, msg.Text, msg.Photo)
}
