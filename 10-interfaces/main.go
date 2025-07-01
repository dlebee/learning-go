package main

import (
	"errors"
	"fmt"
)

type Player interface {
	Health() int
	PlayerName() string
	Attack(enemy Player, weapon Weapon) error
}

type Weapon interface {
	Damage() int
	WeaponName() string
}

type HumanPlayer struct {
	Name          string
	currentHealth int
}

func (p HumanPlayer) Health() int {
	return p.currentHealth
}

func (p HumanPlayer) PlayerName() string {
	return p.Name
}

func (p *HumanPlayer) Attack(enemy Player, weapon Weapon) error {

	if enemy.Health() <= 0 {
		return errors.New(enemy.PlayerName() + " is dead and cannot attack")
	}

	weaponsDamage := weapon.Damage()
	if p.currentHealth <= weaponsDamage {
		p.currentHealth = 0
	} else {
		p.currentHealth -= weaponsDamage
	}

	fmt.Println(enemy.PlayerName(), "attacked", p.PlayerName(), "with", weapon.WeaponName(), "and made", weaponsDamage, "of damage HP left is", p.currentHealth)
	return nil
}

type Sword struct{}

func (s Sword) WeaponName() string {
	return "Sword"
}

func (s Sword) Damage() int {
	return 10
}

type Arrow struct{}

func (s Arrow) WeaponName() string {
	return "Arrow"
}

func (s Arrow) Damage() int {
	return 50
}

func main() {
	player1 := HumanPlayer{Name: "David", currentHealth: 100}
	player2 := HumanPlayer{Name: "Bhrams", currentHealth: 100}

	sword := Sword{}
	arrow := Arrow{}

	player1.Attack(&player2, sword)
	player2.Attack(&player1, arrow)
	player1.Attack(&player2, sword)
	player1.Attack(&player2, sword)
	player2.Attack(&player1, arrow)

	// this will block..
	if err := player1.Attack(&player2, arrow); err != nil {
		fmt.Println("Error:", err)
	}

}
