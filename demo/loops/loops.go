package main

import "fmt"

func main() {
  sumatory := 0
  fmt.Println("Sumatory is:", sumatory)

  // Traditional loop
  for i := 1; i <= 10; i++ {
    sumatory += i
    fmt.Println("Sumatory is:", sumatory)
  }

  // While loop style
  for sumatory > 10 {
    sumatory -= 5
    fmt.Println("Decrement. Sumatory is:", sumatory)
  }
}
