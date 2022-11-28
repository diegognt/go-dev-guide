//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
  var myFavoriteColor string = "Blue"
  fmt.Println("My favorite color is:", myFavoriteColor)

  birthYear, ageInYears := 1990, 32
  fmt.Println("My birth year is:", birthYear)
  fmt.Println("Age in years:", ageInYears)

  var (
    firstInitial = "D"
    lastIntial = "N"
  )
  fmt.Println("The first initial is", firstInitial)
  fmt.Println("The last initial is", lastIntial)

  var ageInDays int
  ageInDays = ageInYears * 365
  fmt.Println("The age in days =", ageInDays)
}
