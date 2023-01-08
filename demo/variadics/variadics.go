// -- Sumary:
// This demo includes several functions that shows how to use "Variadics"

package main

import "fmt"

// summatory of any number of integers as parameter.
func summatory(numbers ...int) int {
  summatory := 0

  for _, number := range numbers {
    summatory += number
  }

  return summatory
}

func main() {
  a := []int{1, 2, 3}
  b := []int{4, 5, 6}

  all := append(a, b...)

  answer := summatory(all...)
  fmt.Println("The sumatory of all number is:", answer)
}
