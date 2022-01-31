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

type Photo struct {
	ChatID    int64
	FirstName string
	LastName  string
	UserName  string
	Caption   string
	Time      time.Time
}