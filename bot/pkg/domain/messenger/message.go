package messenger

import "time"

const (
	TextType      = "text"
	PhotoType     = "photo"
	VoiceType     = "voice"
	VideoNoteType = "videoNote"
	VideoType     = "video"
	PollType      = "poll"
)

type SendMessage struct {
	ChatID         string
	Text           string
	InlineKeyboard [][]string
	ReplyKeyboard  [][]string
	Photo          *Photo
	Voice          *Voice
	VideoNote      *VideoNote
	Video          *Video
	Poll           *Poll
}

func MakeTextMessage(chatID, text string) SendMessage {
	return SendMessage{
		ChatID:         chatID,
		Text:           text,
		InlineKeyboard: nil,
		ReplyKeyboard:  nil,
		Photo:          nil,
		Voice:          nil,
		VideoNote:      nil,
		Video:          nil,
		Poll:           nil,
	}
}

func MakeKeyedTextMessage(chatID, text string, inline [][]string, reply [][]string) SendMessage {
	return SendMessage{
		ChatID:         chatID,
		Text:           text,
		InlineKeyboard: inline,
		ReplyKeyboard:  reply,
		Photo:          nil,
		Voice:          nil,
		VideoNote:      nil,
		Video:          nil,
		Poll:           nil,
	}
}

func MakePhotoMessage(chatID, text string, photo *Photo) SendMessage {
	return SendMessage{
		ChatID:         chatID,
		Text:           text,
		InlineKeyboard: nil,
		ReplyKeyboard:  nil,
		Photo:          photo,
		Voice:          nil,
		VideoNote:      nil,
		Video:          nil,
		Poll:           nil,
	}
}

func MakeKeyedPhotoMessage(chatID, text string, photo *Photo, inline [][]string, reply [][]string) SendMessage {
	return SendMessage{
		ChatID:         chatID,
		Text:           text,
		InlineKeyboard: inline,
		ReplyKeyboard:  reply,
		Photo:          photo,
		Voice:          nil,
		VideoNote:      nil,
		Video:          nil,
		Poll:           nil,
	}
}

type ReceiveMessage struct {
	ChatID    string
	Text      string
	FirstName string
	LastName  string
	UserName  string
	Time      time.Time
	Photo     *Photo
	Voice     *Voice
	VideoNote *VideoNote
	Video     *Video
	Poll      *Poll
}

func (msg ReceiveMessage) GetType() string {
	switch {
	case msg.Photo != nil:
		return PhotoType
	case msg.Voice != nil:
		return VoiceType
	case msg.VideoNote != nil:
		return VideoNoteType
	case msg.Video != nil:
		return VideoType
	case msg.Poll != nil:
		return PollType
	default:
		return TextType
	}
}

func (msg ReceiveMessage) IsPhoto() bool {
	return msg.Photo != nil
}

func (msg ReceiveMessage) IsVoice() bool {
	return msg.Voice != nil
}

func (msg ReceiveMessage) IsVideoNote() bool {
	return msg.VideoNote != nil
}

func (msg ReceiveMessage) IsVideo() bool {
	return msg.Video != nil
}

func (msg ReceiveMessage) IsPoll() bool {
	return msg.Poll != nil
}

type Photo struct {
	FileUniqueID string
}

func NewPhoto(photoID string) *Photo {
	return &Photo{photoID}
}

type Voice struct {
	FileUniqueID string
}

func NewVoice(fileUniqueID string) *Voice {
	return &Voice{FileUniqueID: fileUniqueID}
}

type VideoNote struct {
	FileUniqueID string
}

func NewVideoNote(fileUniqueID string) *VideoNote {
	return &VideoNote{FileUniqueID: fileUniqueID}
}

type Video struct {
	FileUniqueID string
}

func NewVideo(fileUniqueID string) *Video {
	return &Video{FileUniqueID: fileUniqueID}
}

type Poll struct {
	Question string
	Options  []string
}

func NewPoll(question string, options []string) *Poll {
	return &Poll{Question: question, Options: options}
}
