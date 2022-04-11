package msgrouter

import "github.com/nndergunov/tgBot/bot/pkg/domain/messenger"

func (r *MsgRouter) editCollectionDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch r.currentPosition[msg.ChatID] {
	case receivingDetailsPos:
		return r.editCollDetailsDialogue(msg)
	case addingNewPos:
		return r.addingNewDialogue(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) editCollDetailsDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch msg.Text {
	case "Add new":
		r.currentPosition[msg.ChatID] = addingNewPos

		return r.communicator.AddToCollectionResponser(msg)
	case "Choose which to edit":
		r.currentPosition[msg.ChatID] = editingPos

		return r.communicator.EditInCollectionResponser(msg)
	case "Delete":
		r.currentPosition[msg.ChatID] = deletingPos

		return r.communicator.DeletingFromCollectionResponser(msg)
	case "Back":
		r.currentDialogue[msg.ChatID] = homeDialogue
		r.currentPosition[msg.ChatID] = homePos

		return r.communicator.BackResponser(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) addingNewDialogue(msg messenger.ReceiveMessage) messenger.SendMessage {
	success, resp := r.communicator.Adder(msg)

	if success == true {
		r.currentDialogue[msg.ChatID] = homeDialogue
		r.currentPosition[msg.ChatID] = homePos
	}

	return resp
}
