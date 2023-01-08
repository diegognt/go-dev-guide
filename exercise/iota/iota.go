//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type MathOperation int

const (
	Addition MathOperation = iota
	Subtract
	Multiply
	Divide
)
const (
	add = MathOperation(Addition)
	sub = MathOperation(Subtract)
	mul = MathOperation(Multiply)
	div = MathOperation(Divide)
)

func (operation MathOperation) calculate(firstOperand, secondOperand int) int {
	return []int{
		firstOperand + secondOperand,
		firstOperand - secondOperand,
		firstOperand * secondOperand,
		firstOperand / secondOperand,
	}[operation]
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
