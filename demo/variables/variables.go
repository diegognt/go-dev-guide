package main

import "fmt"

func main() {
  // A string variable
  var myName = "Jhon Doe"
  fmt.Println("My name is:", myName)

  // Types annotation
  var name string = "Kathy"
  fmt.Println("A name:", name)

  // Create a assign
  userName := "Admin"
  fmt.Println("A userName:", userName)

  // An uninitialize variable
  var sum int
  fmt.Println("The sum is =", sum)

  // A compound assigment
  partOne, other := 1, 5
  fmt.Println("The part one is", partOne, ", the other part is", other)

  // OK error idiom
  partTwo, other := 2, 0
  fmt.Println("The part two is", partTwo, ", the other part is", other)

  sum = partOne + partTwo
  fmt.Println("The sum is:", sum)

  // Block assigment
  var (
    lessonName = "Variables"
    lessonType = "Demo"
  )
  fmt.Println("The lesson name is:", lessonName)
  fmt.Println("The lesson type is:", lessonType)


  // A compound ignoring a variable
  wordOne, wordTwo, _ := "Hello", "world", "!"
  fmt.Println(wordOne, wordTwo)
}

