package conversationer

import (
	"strconv"
	"strings"

	"github.com/nndergunov/tgBot/bot/pkg/domain/entities"
	"github.com/nndergunov/tgBot/bot/pkg/domain/messenger"
)

func (c Conversationer) EditCollectionResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := "Which action do you want to perform?"

	return messenger.MakeKeyedTextMessage(
		msg.ChatID,
		text,
		c.keyboards[EditKeyboardKey],
		nil,
	)
}

func (c Conversationer) AddToCollectionResponser(msg messenger.ReceiveMessage) messenger.SendMessage {
	text := "Then send me cover of desired album with caption consisting of:\nName of the artist\nName of the album\n" +
		"Genre\nOriginal release year\nYear of the album issue\nLabel\nWhether it is coloured or not (Yes/No)\n" +
		"Location, in which the album is stored\n\nall in one line divided by commas."

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) Adder(msg messenger.ReceiveMessage) (bool, messenger.SendMessage) {
	if msg.Photo == nil {
		text := "And where is the photo?"

		return false, messenger.MakeTextMessage(msg.ChatID, text)
	}

	caption := msg.Text
	caption = strings.ReplaceAll(caption, ", ", ",")
	args := strings.Split(caption, ",")
	numOfArgs := 8

	if len(args) != numOfArgs {
		text := "Looks like not all arguments are in place."

		return false, messenger.MakeTextMessage(msg.ChatID, text)
	}

	relYear, err := strconv.Atoi(args[3])
	if err != nil {
		text := "Release year should be integer, like '1984'."

		return false, messenger.MakeTextMessage(msg.ChatID, text)
	}

	reisYear, err := strconv.Atoi(args[4])
	if err != nil {
		text := "Reissue year should be integer, like '1984'."

		return false, messenger.MakeTextMessage(msg.ChatID, text)
	}

	isColoured := strings.ToLower(args[6]) == "yes"

	album := entities.Album{
		Artist: entities.Artist{
			Name: args[0],
		},
		Name:        args[1],
		Genre:       args[2],
		ReleaseYear: relYear,
		ReissueYear: reisYear,
		Label:       args[5],
		Coloured:    isColoured,
		CoverID:     msg.Photo.FileUniqueID,
	}

	location := entities.Location{
		Owner: entities.User{
			ChatID:   msg.ChatID,
			UserName: msg.UserName,
		},
		Name: args[7],
	}

	err = c.database.AddAlbumToCollection(album, location)
	if err != nil {
		text := "Some error working with database, try again later"

		return true, messenger.MakeTextMessage(msg.ChatID, text)
	}

	text := "Added successfully!"

	return true, messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) DeletingFromCollectionResponser(msg messenger.ReceiveMessage, curAlbum int) messenger.SendMessage {
	err := c.database.DeleteAlbum(curAlbum, msg.ChatID)
	if err != nil {
		text := "Some error working with database, try again later"

		return messenger.MakeTextMessage(msg.ChatID, text)
	}

	text := "Deleted successfully!"

	return messenger.MakeTextMessage(msg.ChatID, text)
}

func (c Conversationer) EditInCollectionResponser(msg messenger.ReceiveMessage, curAlbum int) (bool, messenger.SendMessage) {
	err := c.database.DeleteAlbum(curAlbum, msg.ChatID)
	if err != nil {
		text := "Some error working with database, try again later"

		return false, messenger.MakeTextMessage(msg.ChatID, text)
	}

	text := "Deleted successfully!"

	return true, messenger.MakeTextMessage(msg.ChatID, text)
}
