package ducksgoodimplementation

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}
type Squeak struct{}
type MuteQuack struct{}

func (Quack) Quack()     { fmt.Println("Quack Quack!") }
func (Squeak) Quack()    { fmt.Println("Squeak Squeak!") }
func (MuteQuack) Quack() { fmt.Println("... ...") }
