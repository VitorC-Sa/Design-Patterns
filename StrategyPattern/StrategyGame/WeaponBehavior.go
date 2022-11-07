package strategygame

import "fmt"

type WeaponBehavior interface {
	UseWeapon()
}

type KnifeBehavior struct{}
type BowAndArrowBehavior struct{}
type AxeBehavior struct{}
type SwordBehavior struct{}

func (KnifeBehavior) UseWeapon()       { fmt.Println("Handle a knife!") }
func (BowAndArrowBehavior) UseWeapon() { fmt.Println("Handle a bow and arrow!") }
func (AxeBehavior) UseWeapon()         { fmt.Println("Handle an Axe!") }
func (SwordBehavior) UseWeapon()       { fmt.Println("Handle a sword!") }
