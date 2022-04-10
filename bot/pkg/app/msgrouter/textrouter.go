package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (r MsgRouter) textRouter(msg domain.ReceiveMessage) domain.SendMessage {
	switch r.currentDialogue[msg.ChatID] {
	case homeDialogue:
		return r.homeDialogue(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) homeDialogue(msg domain.ReceiveMessage) (answer domain.SendMessage) {
	return domain.MakeTextMessage(msg.ChatID, msg.Text)
}
