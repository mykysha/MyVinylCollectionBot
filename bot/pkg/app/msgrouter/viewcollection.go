package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (r *MsgRouter) viewCollectionDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch r.currentPosition[msg.ChatID] {
	case homePos:
		return r.showCollectionDialogue(msg)
	case shownPos:
		return r.gettingNextStepDialogue(msg)
	case receivingDetailsPos:
		return r.expandingDialogue(msg)
	case shownFullPos:
		return r.editOneDialogue(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) showCollectionDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	r.currentPosition[msg.ChatID] = shownPos

	return r.communicator.ViewCollectionResponser(msg)
}

func (r *MsgRouter) gettingNextStepDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch msg.Text {
	case "Choose one to expand":
		r.currentPosition[msg.ChatID] = receivingDetailsPos

		return r.communicator.ExpandingResponser(msg)
	case "Back":
		r.currentDialogue[msg.ChatID] = homeDialogue
		r.currentPosition[msg.ChatID] = homePos

		return r.communicator.BackResponser(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) expandingDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	success, id, resp := r.communicator.ShowingFullResponser(msg)

	if success == true {
		r.currentPosition[msg.ChatID] = shownFullPos

		r.currentAlbum[msg.ChatID] = id
	} else {
		r.currentDialogue[msg.ChatID] = homeDialogue
		r.currentPosition[msg.ChatID] = homePos
	}

	return resp
}
