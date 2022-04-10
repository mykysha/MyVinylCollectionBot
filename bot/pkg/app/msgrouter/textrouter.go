package msgrouter

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (r MsgRouter) textRouter(msg messenger.ReceiveMessage) messenger.SendMessage {
	switch r.currentDialogue[msg.ChatID] {
	case homeDialogue:
		return r.homeDialogue(msg)
	default:
		return r.communicator.UnknownTypeResponser(msg)
	}
}

func (r *MsgRouter) homeDialogue(msg messenger.ReceiveMessage) (answer messenger.SendMessage) {
	switch msg.Text {
	case "View collection":
		return r.communicator.ViewCollectionResponser(msg)
		// TODO change home to viewCollection.
	case "Edit collection":
		return r.communicator.EditCollectionResponser(msg)
		// TODO change home to editCollection.
	case "View genres":
		return r.communicator.ViewGenresResponser(msg)
	case "View artists":
		return r.communicator.ViewArtistsResponser(msg)
	case "View wishlist":
		return r.communicator.ViewWishlistResponser(msg)
		// TODO change home to viewWishlist.
	case "Edit wishlist":
		return r.communicator.EditWishlistResponser(msg)
		// TODO change home to viewWishlist.
	case "Move from wishlist to collection":
		return r.communicator.MoveResponser(msg)
	case "Bot info":
		return r.communicator.BotInfoResponser(msg)
	default:
		return r.communicator.UnsupportedResponser(msg)
	}
}
