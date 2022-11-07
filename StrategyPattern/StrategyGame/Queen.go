package strategygame

import (
	"fmt"
)

type Queen struct {
	character
}

func (Queen) Display() { fmt.Println("Fight like a Queen!") }

func NewQueen() *Queen {
	return &Queen{
		character: character{
			weapon:   nil,
			charName: "Queen",
			Life:     50,
			DmgSec:   5,
		},
	}
}
