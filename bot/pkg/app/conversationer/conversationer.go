package conversationer

import "github.com/nndergunov/tgBot/bot/pkg/domain"

const (
	StartKeyboardKey = "startKeyboard"
	EditKeyboardKey  = "editKeyboard"
	ViewKeyboardKey  = "viewKeyboard"
)

type Conversationer struct {
	answers   domain.Answers
	keyboards map[string][][]string
}

func NewConver(answers domain.Answers, keyboards map[string][][]string) *Conversationer {
	return &Conversationer{
		answers:   answers,
		keyboards: keyboards,
	}
}
