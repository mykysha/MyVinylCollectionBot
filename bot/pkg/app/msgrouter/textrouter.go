package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (r MsgRouter) textRouter(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch r.currentDialogue[msg.ChatID] {
	case homeDialogue:
		return r.homeDialogue(msg)
	case viewCollectionDialogue:
		return r.viewCollectionDialogue(msg)
	case editCollectionDialogue:
		return r.editCollectionDialogue(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) homeDialogue(msg messenger.ReceiveMessage) (answer messenger.SendMessage) {
	switch msg.Text {
	case "View collection":
		r.currentDialogue[msg.ChatID] = viewCollectionDialogue

		return r.viewCollectionDialogue(msg)
	case "Edit collection":
		r.currentDialogue[msg.ChatID] = editCollectionDialogue
		r.currentPosition[msg.ChatID] = receivingDetailsPos

		return r.communicator.EditCollectionResponser(msg)
	case "View genres":
		return r.communicator.ViewGenresResponser(msg)
	case "View artists":
		return r.communicator.ViewArtistsResponser(msg)
	case "Bot info":
		return r.communicator.BotInfoResponser(msg)
	case "Convert to XSLX":
		return r.communicator.ConvertToXLSX(msg)
	default:
		return r.communicator.UnsupportedResponser(msg)
	}
}
