package strategygame

import (
	"fmt"
)

type Knight struct {
	character
}

func (Knight) Display() { fmt.Println("Fight like a Knight!") }

func NewKnight() *Knight {
	return &Knight{
		character: character{
			weapon:   nil,
			charName: "Knight",
			Life:     60,
			DmgSec:   3,
		},
	}
}
