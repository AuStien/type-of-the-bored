package words

import (
	"fmt"

	"github.com/austien/type-of-the-bored/ansi"
)

type Letter interface{}

type letter struct {
	Character    rune
	Color        ansi.CODE
	HasUnderline bool
}

func NewLetter(char rune) *letter {
	return &letter{
		Character:    char,
		HasUnderline: false,
		Color:        ansi.NONE,
	}
}

func (l *letter) CompareCharacter(char rune) {
	if char == l.Character || (char == '\x00' && l.Character == ' ') {
		l.Color = ansi.GREEN
	} else {
		l.Color = ansi.RED
	}
}

func (l *letter) ToString() string {
	underline := ansi.NONE
	if l.HasUnderline {
		underline = ansi.UNDERLINE
	}

	color := l.Color.Value

	if l.Character == ' ' && l.Color.Name == ansi.RED.Name {
		color = "\033[41m"
	}

	return fmt.Sprintf("%s%s%s%s%s", underline, color, string(l.Character), ansi.RESET_UNDERLINE, ansi.RESET)
}
