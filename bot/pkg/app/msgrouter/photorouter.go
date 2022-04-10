package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (r MsgRouter) photoRouter(msg messenger.ReceiveMessage) messenger.SendMessage {
	return messenger.MakePhotoMessage(msg.ChatID, msg.Text, msg.Photo)
}
