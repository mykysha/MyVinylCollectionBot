package conversationer

import (
	"fmt"
	"strings"

	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (c Conversationer) ViewCollectionResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	albums, err := c.database.GetCollection(int(msg.ChatID))
	if err != nil {
		text := "Some error working with database, try again later"

		return messenger.MakeTextMessage(msg.ChatID, text)
	}

	var text string

	for i := 0; i < len(albums); i++ {
		currAlbum := albums[i]

		text += fmt.Sprintf("%d. Name: %s; Artist: %s\n", i+1, currAlbum.Name, currAlbum.Artist.Name)
	}

	text = strings.TrimRight(text, "\n")

	return messenger.MakeKeyedTextMessage(msg.ChatID, text, c.keyboards[ViewKeyboardKey], nil)
}
