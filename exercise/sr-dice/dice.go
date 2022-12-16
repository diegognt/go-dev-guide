//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollAnalysis(dice int, sumatory int) string {
  analysis := ""

	switch sumatory := sumatory; {
	case dice == 2 && sumatory == 1:
		analysis = "Snake eyes"
	case sumatory == 7:
		analysis = "Lucky 7"
	case sumatory%2 == 0:
		analysis = "Even"
	case sumatory%2 == 1:
		analysis = "Odd"
	}

  return analysis
}

func rollDice(sides int) int {
  return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rolls := 5
	dices, dieSides := 2, 12
	rolledTime := 1

	for rolledTime <= rolls {
		fmt.Println("\nRoll #", rolledTime)
		sumatoryOfDice := 0

		for i := 1; i <= dices; i++ {
      dieValue := rollDice(dieSides)
			fmt.Println("The Die", i, "rolled with a", dieValue)
			sumatoryOfDice += dieValue
		}

		fmt.Println("Dice value", sumatoryOfDice, "and is a:", rollAnalysis(dices, sumatoryOfDice) )

		rolledTime++
	}
}
