package strategygame

import (
	"fmt"
	"time"
)

type iCharacter interface {
	Display()
	TakeAHit(hit int)
	GetName() string
	GetHp() int
	GetDmgSec() int
	UseWeapon()
	GetWeaponDmg() int
}

type character struct {
	weapon   WeaponBehavior
	charName string
	Life     int
	DmgSec   int
}

func (c *character) SetWeapon(w WeaponBehavior) { c.weapon = w }
func (c *character) TakeAHit(hit int)           { c.Life -= hit }
func (c character) GetName() string             { return c.charName }
func (c character) GetHp() int                  { return c.Life }
func (c character) GetDmgSec() int              { return c.DmgSec }
func (c character) UseWeapon()                  { c.weapon.UseWeapon() }
func (c character) GetWeaponDmg() int           { return c.weapon.GetWeapDmg() }

func Fight(c1, c2 iCharacter) {
	fmt.Println(c1.GetName())
	c1.Display()
	fmt.Println("Versus")
	fmt.Println(c2.GetName())
	c2.Display()

	hit := func(atk, def iCharacter) int {
		fmt.Printf("%s's turn\n", atk.GetName())
		atk.UseWeapon()
		totalDmg := atk.GetDmgSec() + atk.GetWeaponDmg()
		fmt.Printf("%s's atk points: %d!\n", atk.GetName(), atk.GetDmgSec())
		fmt.Printf("%s hits a %d total damage!\n", atk.GetName(), totalDmg)
		def.TakeAHit(totalDmg)
		fmt.Printf("%s has %d lifepoints now!\n", def.GetName(), def.GetHp())
		return def.GetHp()
	}

	for i := 1; ; i++ {
		time.Sleep(time.Second)
		fmt.Printf("\nTurn N°:%d\n", i)

		if hit(c1, c2) <= 0 {
			fmt.Printf("%s Wins!\n", c1.GetName())
			break
		}

		if hit(c2, c1) <= 0 {
			fmt.Printf("%s Wins!\n", c2.GetName())
			break
		}
	}
}
