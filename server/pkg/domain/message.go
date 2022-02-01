package domain

import "time"

type TextMessage struct {
	ChatID    int64
	FirstName string
	LastName  string
	UserName  string
	Text      string
	Time      time.Time
}

func NewTextMessage(chatID int64, FirstName, LastName, UserName, Text string, time time.Time) TextMessage {
	return TextMessage{
		ChatID:    chatID,
		FirstName: FirstName,
		LastName:  LastName,
		UserName:  UserName,
		Text:      Text,
		Time:      time,
	}
}

type Photo struct {
	ChatID       int64
	Height       int64
	Width        int64
	FirstName    string
	LastName     string
	UserName     string
	Caption      string
	FileUniqueID string
	Time         time.Time
}

func NewPhoto(chatID, height, width int64, firstName, lastName, userName, caption, uniqueID string, time time.Time) Photo {
	return Photo{
		ChatID:       chatID,
		Height:       height,
		Width:        width,
		FirstName:    firstName,
		LastName:     lastName,
		UserName:     userName,
		Caption:      caption,
		FileUniqueID: uniqueID,
		Time:         time,
	}
}

type VoiceMessage struct {
	ChatID    int64
	FirstName string
	LastName  string
	UserName  string
	Text      string
	Time      time.Time
}

type VideoNote struct {
	ChatID    int64
	FirstName string
	LastName  string
	UserName  string
	Text      string
	Time      time.Time
}

type Video struct {
	ChatID    int64
	Duration  int64
	Height    int64
	Width     int64
	FirstName string
	LastName  string
	UserName  string
	Text      string
	Time      time.Time
}

type Poll struct {
	ChatID    int64
	FirstName string
	LastName  string
	UserName  string
	Question  string
	Time      time.Time
}