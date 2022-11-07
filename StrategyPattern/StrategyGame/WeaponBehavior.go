package strategygame

import "fmt"

const (
	knifeDmg = 2
	bowDmg   = 3
	SwordDmg = 4
	AxeDmg   = 4
)

type WeaponBehavior interface {
	UseWeapon()
	GetWeapDmg() int
}

type KnifeBehavior struct{}
type BowAndArrowBehavior struct{}
type AxeBehavior struct{}
type SwordBehavior struct{}

func (w KnifeBehavior) UseWeapon()       { fmt.Printf("Handle a knife with %d damage!\n", knifeDmg) }
func (w BowAndArrowBehavior) UseWeapon() { fmt.Printf("Handle a bow with %d damage!\n", bowDmg) }
func (w AxeBehavior) UseWeapon()         { fmt.Printf("Handle an Axe with %d damage!\n", AxeDmg) }
func (w SwordBehavior) UseWeapon()       { fmt.Printf("Handle a sword with %d damage!\n", SwordDmg) }

func (KnifeBehavior) GetWeapDmg() int       { return knifeDmg }
func (BowAndArrowBehavior) GetWeapDmg() int { return bowDmg }
func (AxeBehavior) GetWeapDmg() int         { return AxeDmg }
func (SwordBehavior) GetWeapDmg() int       { return SwordDmg }
