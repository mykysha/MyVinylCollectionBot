package msgrouter

import "github.com/nndergunov/tgBot/bot/pkg/domain/messenger"

func (r *MsgRouter) viewCollectionDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	return r.communicator.ViewCollectionResponser(msg)
}
