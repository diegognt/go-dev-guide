package main

import "fmt"

func doubleIt(x int) int {
  return x + x
}

func addIt(leftHandside, rightHandSide int) int {
  return leftHandside + rightHandSide
}

func greet() {
  fmt.Println("Hello from the greet function!")
}

func main () {
  greet()

  dozen := doubleIt(6)
  fmt.Println("A dozen is", dozen)

  bakerDozen := addIt(dozen, 1)
  fmt.Println("A baker's dozen is", bakerDozen)

  anotherBakerDozen := addIt(doubleIt(6), 1)
  fmt.Println("Another baker's dozen is", anotherBakerDozen)
}
