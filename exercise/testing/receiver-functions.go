//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name                  string
	health, maximumHealth uint
	energy, maximumEnergy uint
}

func (player *Player) addHealth(amount uint) {
	fmt.Println("\nAdding health to the player", player.name)
	player.health += amount

	if player.health > player.maximumHealth {
		player.health = player.maximumHealth
	}

	fmt.Println("The player", player.name, "now have", player.health, "points of healt.")
}

func (player *Player) applyDamage(amount uint) {
	fmt.Println("\nApplying damage to the player", player.name)

	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}

	fmt.Println("The player", player.name, "now have", player.health, "points of healt.")

}

func (player *Player) addEnergy(amount uint) {
	fmt.Println("\nAdding energy to the player", player.name)
	player.energy += amount

	if player.energy > player.maximumEnergy {
		player.energy = player.maximumEnergy
	}

	fmt.Println("The player", player.name, "now have", player.energy, "points of energy.")
}

func (player *Player) consumeEnergy(amount uint) {
	fmt.Println("\nConsuming player", player.name, "energy")

	if player.energy-amount > player.maximumEnergy {
		player.energy = 0
	} else {
		player.energy -= amount
	}

	fmt.Println("The player", player.name, "now have", player.energy, "points of energy.")
}
func main() {
	playerOne := Player{
		name:          "PlayerOne",
		health:        2,
		maximumHealth: 30,
		energy:        10,
		maximumEnergy: 50,
	}

	playerOne.addEnergy(8)
	playerOne.applyDamage(15)
	playerOne.addEnergy(30)
	playerOne.consumeEnergy(16)

}
