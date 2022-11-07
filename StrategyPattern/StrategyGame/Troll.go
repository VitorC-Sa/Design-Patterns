package strategygame

import (
	"fmt"
)

type Troll struct {
	character
}

func (c Troll) Display() {
	fmt.Println("Fight like a Troll!")
}

func NewTroll() *Troll {
	return &Troll{
		character: character{
			weapon:   nil,
			charName: "Troll",
			Life:     70,
			DmgSec:   2,
		},
	}
}
