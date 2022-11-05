package ducksgoodimplementation

import "fmt"

type mallardDuck struct {
	duck
	//Display()
}

func (mallardDuck) Display() { fmt.Println("I'm a real Mallard duck!") }

func NewMallardDuck() *mallardDuck {
	return &mallardDuck{
		duck: duck{
			flyBehavior:   FlyWithWings{},
			quackBehavior: Quack{},
		},
	}
}
