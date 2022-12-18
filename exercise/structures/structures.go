//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y float32
}

type Rectangle struct {
	a Coordinate // bottom left coordinates
	b Coordinate // bottom right coordinates
}

func rectangleWidth(rectangle Rectangle) float32 {
	return rectangle.b.x - rectangle.a.x
}

func rectangleLength(rectangle Rectangle) float32 {
	return rectangle.a.y - rectangle.b.y
}

func rectangleArea(rectangle Rectangle) float32 {
	return rectangleLength(rectangle) * rectangleWidth(rectangle)
}

func rectanglePerimeter(rectangle Rectangle) float32 {
	return (rectangleLength(rectangle) * 2) + (rectangleWidth(rectangle) * 2)
}

func doubleRectangleSize(rectangle Rectangle) Rectangle {
	rectangle.a.y *= 2
	rectangle.b.x *= 2

	return rectangle
}
func main() {
	rectangle := Rectangle{
    a: Coordinate{0.5, 14.6}, 
    b: Coordinate{3.4, 6.9},
  }

	fmt.Println("The area of the rectangle is", rectangleArea(rectangle))
	fmt.Println("The parameter of the rectangle is:", rectanglePerimeter(rectangle))

	rectangle = doubleRectangleSize(rectangle)

	fmt.Println("The area of the doubled rectangle is", rectangleArea(rectangle))
	fmt.Println("The parameter of the doubled rectangle is:", rectanglePerimeter(rectangle))
}
