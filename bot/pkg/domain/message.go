package domain

import "time"

type TextMessage struct {
	ChatID    int64
	Text      string
	FirstName string
	LastName  string
	UserName  string
	Time      time.Time
}

type Photo struct {
	ChatID       int64
	Height       int64
	Width        int64
	Caption      string
	FileUniqueID string
	FirstName    string
	LastName     string
	UserName     string
	Time         time.Time
}

type VoiceMessage struct {
	ChatID    int64
	Duration  int
	FirstName string
	LastName  string
	UserName  string
	Time      time.Time
}

type VideoNote struct {
	ChatID    int64
	Duration  int
	FirstName string
	LastName  string
	UserName  string
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
	Caption   string
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
