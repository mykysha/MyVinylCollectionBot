package conversationer

import (
	"github.com/nndergunov/tgBot/bot/pkg/db"
	"github.com/nndergunov/tgBot/bot/pkg/domain/answerer"
)

const (
	StartKeyboardKey        = "startKeyboard"
	EditKeyboardKey         = "editKeyboard"
	ViewKeyboardKey         = "viewKeyboard"
	ExpandedViewKeyboardKey = "expandedView"
	EditViewKeyboardKey     = "editView"
)

const timeFormat = "02 Jan 06 15:04 MST"

type Conversationer struct {
	keyboards map[string][][]string
	answers   answerer.Answers
	database  *db.Database
}

func NewConver(database *db.Database, answers answerer.Answers, keyboards map[string][][]string) *Conversationer {
	return &Conversationer{
		keyboards: keyboards,
		answers:   answers,
		database:  database,
	}
}
