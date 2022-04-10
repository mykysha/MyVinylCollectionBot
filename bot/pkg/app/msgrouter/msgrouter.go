package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/app/conversationer"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

const (
	homeDialogue           = "home"
	viewCollectionDialogue = "view collection"
	editCollectionDialogue = "edit collection"
	viewWishlistDialogue   = "view wishlist"
	editWishlistDialogue   = "edit wishlist"
	moveDialogue           = "move from wishlist to collection"
)

type MsgRouter struct {
	currentDialogue map[int64]string
	currentPosition map[int64]string
	communicator    *conversationer.Conversationer
}

func NewMsgRouter(communicator *conversationer.Conversationer) *MsgRouter {
	currDialogue := make(map[int64]string)
	currPos := make(map[int64]string)

	return &MsgRouter{
		currentDialogue: currDialogue,
		currentPosition: currPos,
		communicator:    communicator,
	}
}

func (r *MsgRouter) Route(msg messenger.ReceiveMessage) messenger.SendMessage {
	if _, ok := r.currentDialogue[msg.ChatID]; !ok {
		r.currentDialogue[msg.ChatID] = homeDialogue
	}

	return r.routeByType(msg)
}

// todo

func (r MsgRouter) routeByType(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch {
	case msg.Voice != nil:
		return r.communicator.VoiceResponser(msg)
	case msg.VideoNote != nil:
		return r.communicator.VideoNoteResponser(msg)
	case msg.Photo != nil:
		return r.photoRouter(msg)
	case msg.Video != nil:
		return r.communicator.VideoResponser(msg)
	case msg.Poll != nil:
		return r.communicator.PollResponser(msg)
	case msg.Text != "":
		return r.textRouter(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}
