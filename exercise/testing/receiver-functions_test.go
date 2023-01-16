//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
  return Player{
    name: "testPlayer",
    health: 100,
    maximumHealth: 100,
    energy: 500,
    maximumEnergy: 500,
  }
}

func TestHealth(test *testing.T) {
  player := newPlayer()
  player.addHealth(999)

  if player.health > player.maximumHealth {
    test.Fatalf("Health went over the limit, %v, want: %v", player.health, player.maximumHealth)
  }

  player.applyDamage(player.maximumHealth + 1)
  if player.health < 0 {
    test.Fatalf("Health: %v. Minimum: 0", player.health)
  }
  if player.health > player.maximumHealth {
    test.Fatalf("Health: %v. Maximum: %v", player.health, player.maximumHealth)
  }
}

func TestEnergy(test *testing.T) {
  player := newPlayer()
  player.addEnergy(999)

  if player.energy > player.maximumEnergy {
    test.Fatalf("Health went over the limit, %v, want: %v", player.energy, player.maximumEnergy)
  }

  player.consumeEnergy(player.maximumEnergy + 1)
  if player.energy < 0 {
    test.Fatalf("Health: %v. Minimum: 0", player.energy)
  }
  if player.energy > player.maximumEnergy {
    test.Fatalf("Health: %v. Maximum: %v", player.energy, player.maximumEnergy)
  }
}
