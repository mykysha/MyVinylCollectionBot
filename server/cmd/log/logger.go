package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	log *log.Logger
}

func (l *Logger) Init(log *log.Logger) {
	l.log = log
}

func (l Logger) ErrorLog(err error) {
	l.log.Println(err)
}

func (l Logger) VoiceLog(li logInfo, voiceDuration int) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, "+
		"VoiceMessage duration: %d seconds, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, voiceDuration, li.time)

	l.log.Println(logMsg)
}

func (l Logger) VideoNoteLog(li logInfo, videoNoteDuration int) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, "+
		"VideoNote duration: %d seconds, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, videoNoteDuration, li.time)

	l.log.Println(logMsg)
}

func (l Logger) PhotoLog(li logInfo, height, width int) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, "+
		"Photo height: %d, Photo width: %d, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, height, width, li.time)

	l.log.Println(logMsg)
}

func (l Logger) VideoLog(li logInfo, videoDuration, height, width int) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, "+
		"Video duration: %d, Video height: %d, Video width: %d, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, videoDuration, height, width, li.time)

	l.log.Println(logMsg)
}

func (l Logger) PollLog(li logInfo, question string) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, PollQuestion: %s, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, question, li.time)

	l.log.Println(logMsg)
}

func (l Logger) TextLog(li logInfo, text string) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, Text: %s, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, text, li.time)

	l.log.Println(logMsg)
}

func (l Logger) DefaultLog(li logInfo) {
	logMsg := fmt.Sprintf("ID: %v, FirstName: %s, LastName: %s, UserName: %s, Time: %v",
		li.chatID, li.firstName, li.lastName, li.userName, li.time)

	l.log.Println(logMsg)
}