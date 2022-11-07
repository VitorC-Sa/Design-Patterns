package strategygame

import (
	"fmt"
)

type King struct {
	character
}

func (King) Display() { fmt.Println("Fight like a King!") }

func NewKing() *King {
	return &King{
		character: character{
			weapon:   nil,
			charName: "King",
			Life:     55,
			DmgSec:   4,
		},
	}
}
