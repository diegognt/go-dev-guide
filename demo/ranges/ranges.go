// - Summary: A program that print out the glyph 
// representation of the string in a slice
// to do so, lets use the "range" keyword

package main

import "fmt"

func main() {
  slice := []string{"Hello", "world", "!"}

  for index, element := range slice {
    fmt.Println(index, element, ":")

    for _, character := range element {
      fmt.Printf("  %q\n", character)
    }
  }
}
