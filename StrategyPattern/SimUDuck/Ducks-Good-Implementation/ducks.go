package ducksgoodimplementation

import "fmt"

//Duck Superclass
type duck struct {
	//Swim()
	//Display() -> abstract (child class implements)
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (duck) Swim()                                { fmt.Println("A beautiful Duck Swimming!") } //implements Swim superclass method
func (d duck) PerformFly()                        { d.flyBehavior.Fly() }
func (d duck) PerformQuack()                      { d.quackBehavior.Quack() }
func (d *duck) SetFlyBehavior(fb FlyBehavior)     { d.flyBehavior = fb }
func (d *duck) SetQuackBehavior(qb QuackBehavior) { d.quackBehavior = qb }

/*
	Main Implementation Example:

	import (
		dg "GoDesignPatterns/SimUDuck/Ducks-Good-Implementation"
	)

	func main() {
		duck := dg.NewMallardDuck()
		duck.Swim()
		duck.Display()
		duck.PerformFly()
		duck.PerformQuack()

		duck.SetFlyBehavior(dg.FlyNoWay{})
		duck.SetQuackBehavior(dg.Squeak{})
		duck.PerformFly()
		duck.PerformQuack()
}
*/
