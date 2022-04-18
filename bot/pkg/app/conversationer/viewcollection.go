package conversationer

import (
	"fmt"
	"strconv"
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

func (c Conversationer) ExpandingResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := "Choose number of the album you want to expand"

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) ShowingFullResponser(msg messenger.ReceiveMessage) (bool, int, messenger.SendMessage) {
	id, err := strconv.Atoi(msg.Text)
	if err != nil {
		text := "ID should be integer, like '1', '2', '3'"

		return false, 0, messenger.MakeTextMessage(msg.ChatID, text)
	}

	albums, err := c.database.GetCollection(int(msg.ChatID))
	if err != nil {
		text := "Could not find your collection. Are you sure you have it?"

		return false, 0, messenger.MakeTextMessage(msg.ChatID, text)
	}

	if id > len(albums) || id <= 0 {
		text := "ID should be integer from the list."

		return false, 0, messenger.MakeTextMessage(msg.ChatID, text)
	}

	album := albums[id-1]

	text := fmt.Sprintf("Name: %s, Artist: %s, Genre: %s, Label: %s, Release Year: %d, Reissue Year: %d",
		album.Name, album.Artist, album.Genre, album.Label, album.ReleaseYear, album.ReissueYear)

	if album.Coloured {
		text += ", Coloured"
	}

	photo := messenger.NewPhoto(album.CoverID)

	return true, id, messenger.MakeKeyedPhotoMessage(msg.ChatID, text, photo, c.keyboards[EditViewKeyboardKey], nil)
}
