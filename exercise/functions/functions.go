//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeter(personName string) {
  fmt.Println("Hello ther", personName)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func message() string {
  return "A random message from a inside a function."
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThree(numberOne, numberTwo, numberThree int) int {
  return numberOne + numberTwo + numberThree
}

//* Write a function that returns any number
func numberTwo() int {
  return 2
}

//* Write a function that returns any two numbers
func twoFives()(int, int) {
  return 5, 5
}

func main() {

  greeter("Jhon")
  fmt.Println(message())

  numberA, numberB := twoFives()
  sumOfThreeNumbers := addThree(numberTwo(), numberA, numberB)
  fmt.Println("The sum of three numbers", sumOfThreeNumbers)
}
