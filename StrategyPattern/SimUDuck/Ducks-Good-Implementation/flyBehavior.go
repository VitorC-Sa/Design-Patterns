package ducksgoodimplementation

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}
type FlyNoWay struct{}
type FlyRocketPowered struct{}

func (FlyWithWings) Fly()     { fmt.Println("I Can Fly with Wings!") }
func (FlyNoWay) Fly()         { fmt.Println("I Can't Fly!") }
func (FlyRocketPowered) Fly() { fmt.Println("I'm flying with a rocket!") }
