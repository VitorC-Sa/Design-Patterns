package ducksgoodimplementation

import "fmt"

type modelDuck struct {
	duck
	//Display()
}

func (modelDuck) Display() { fmt.Println("I'm a real Model duck!") }

func NewModelDuck() *modelDuck {
	return &modelDuck{
		duck: duck{
			flyBehavior:   FlyNoWay{},
			quackBehavior: Quack{},
		},
	}
}
