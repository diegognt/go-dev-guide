package main

import "fmt"

func average (a, b, c int) float32 {
  return float32(a+b+c) / 3
}

func main() {
  quizOne, quizTwo, quizThree := 9, 7, 8

  if(quizOne > quizTwo) {
    fmt.Println("The quiz one scored higher than quiz two.")
  } else if quizOne < quizTwo {
    fmt.Println("The quiz two scored higher than quiz one.")
  } else {
    fmt.Println("Quiz one and quiz two have the same score.")
  }

  if average(quizOne, quizTwo, quizThree) > 7 {
    fmt.Println("Acceptable grade.")
  }
}
