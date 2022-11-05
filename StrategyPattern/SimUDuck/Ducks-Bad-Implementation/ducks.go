package ducksbadimplementation

import "fmt"

type IDuck interface {
	Quack()
	Swim()
	Display()
	Fly()
}

type Duck struct {
	//Quack()
	//Swim()
	//Display() -> abstract (child class implements)
	//Fly()
}

func (Duck) Quack() { fmt.Println("Quack Quack!") }               //implements Quack superclass method
func (Duck) Swim()  { fmt.Println("A beautiful Duck Swimming!") } //implements Swim superclass method
func (Duck) Fly()   { fmt.Println("I Can Fly!") }                 //implements Fly superclass method

type MallardDuck struct {
	Duck
	//Display()
}

func (MallardDuck) Display() { fmt.Println("Mallard Duck Noise!") } //implements Display method

type RedheadDuck struct {
	Duck
	//Display()
}

func (RedheadDuck) Display() { fmt.Println("Redhead Duck Noise!") } //implements Display method

type RubberDuck struct {
	Duck
	//Quack() -> Squack()
	//Display()
	//Fly() -> Do Nothing!
}

func (RubberDuck) Quack()   { fmt.Println("Squeak Squeak!") }     //implements and change Quack method to Squeak method
func (RubberDuck) Display() { fmt.Println("Rubber Duck Noise!") } //implements Display method
func (RubberDuck) Fly()     {}                                    //implements and change Fly method to Do Nothing!

func GoDuck(duck IDuck) {
	duck.Quack()
	duck.Swim()
	duck.Display()
	duck.Fly()
	fmt.Println()
}
