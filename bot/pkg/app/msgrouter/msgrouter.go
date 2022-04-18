package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/app/conversationer"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

const (
	homeDialogue           = "home"
	viewCollectionDialogue = "view collection"
	editCollectionDialogue = "edit collection"
)

const (
	homePos             = "home"
	receivingDetailsPos = "receiving details"
	addingNewPos        = "adding new"
	shownPos            = "shown albums"
	shownFullPos        = "shown full album"
)

type MsgRouter struct {
	currentDialogue map[int64]string
	currentPosition map[int64]string
	currentAlbum    map[int64]int
	communicator    *conversationer.Conversationer
}

func NewMsgRouter(communicator *conversationer.Conversationer) *MsgRouter {
	currDialogue := make(map[int64]string)
	currPos := make(map[int64]string)
	currAlbum := make(map[int64]int)

	return &MsgRouter{
		currentDialogue: currDialogue,
		currentPosition: currPos,
		currentAlbum:    currAlbum,
		communicator:    communicator,
	}
}

func (r *MsgRouter) Route(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch msg.Text {
	case "/start":
		r.currentDialogue[msg.ChatID] = homeDialogue
		r.currentPosition[msg.ChatID] = homePos

		return r.communicator.StartResponser(msg)
	case "/help":
		return r.communicator.HelpResponser(msg)
	}

	if _, ok := r.currentDialogue[msg.ChatID]; !ok {
		r.currentDialogue[msg.ChatID] = homeDialogue
	}

	if _, ok := r.currentPosition[msg.ChatID]; !ok {
		r.currentPosition[msg.ChatID] = homePos
	}

	if msg.Text == "Take me home!" {
		r.currentDialogue[msg.ChatID] = homeDialogue
		r.currentPosition[msg.ChatID] = homePos

		return r.communicator.TakeHomeResponser(msg)
	}

	resp := r.routeByType(msg)
	resp = r.communicator.AddHome(resp)

	return resp
}

func (r MsgRouter) routeByType(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch {
	case msg.Voice != nil:
		return r.communicator.VoiceResponser(msg)
	case msg.VideoNote != nil:
		return r.communicator.VideoNoteResponser(msg)
	case msg.Video != nil:
		return r.communicator.VideoResponser(msg)
	case msg.Poll != nil:
		return r.communicator.PollResponser(msg)
	case msg.Text != "" || msg.Photo != nil:
		return r.textRouter(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}
