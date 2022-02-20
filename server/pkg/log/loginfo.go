package logger

import (
	"time"
)

type logInfo struct {
	chatID    int64
	firstName string
	lastName  string
	userName  string
	time      time.Time
}

func (l Logger) NewLog(chatID int64, firstName, lastName, userName string, time time.Time) logInfo {
	return logInfo{
		chatID:    chatID,
		firstName: firstName,
		lastName:  lastName,
		userName:  userName,
		time:      time,
	}
}
