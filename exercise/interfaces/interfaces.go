//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

// * The shop has lifts for multiple vehicle sizes/types:
//   - Motorcycles: small lifts
//   - Cars: standard lifts
//   - Trucks: large lifts
const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name

// NOTE: If an operand implements method String()
// string, that method will be invoked to convert the object to a string,
// which will then be formatted as required by the verb (if any).
// To avoid recursion in cases use "string()"

func (motorcycle Motorcycle) String() string {
	return fmt.Sprintf("\"%s\" motorcycle", string(motorcycle))
}

func (car Car) String() string {
	return fmt.Sprintf("\"%s\" car", string(car))
}

func (truck Truck) String() string {
	return fmt.Sprintf("\"%s\" truck", string(truck))
}

// Implementing the PickLift interface on each vehicle.

func (motorcycle Motorcycle) PickLift() Lift {
	return SmallLift
}

func (car Car) PickLift() Lift {
	return StandardLift
}

func (truck Truck) PickLift() Lift {
	return LargeLift
}

// Sends a vehicle to the correct lifter based on its type
//
// Works with vehicle that implements the "LiftPicker" interface.
//
// Takes an argument:
//   * vehicle (LiftPicker): A vehicle to work with. 
func sendToLift(vehicle LiftPicker) {
	switch vehicle.PickLift() {
	case SmallLift:
		fmt.Printf("Sending the %v to a small lift\n", vehicle)
	case StandardLift:
		fmt.Printf("Sending the %v to a standard lift\n", vehicle)
	case LargeLift:
		fmt.Printf("sending the %v to a large lift\n", vehicle)
	}
}

func main() {
	car := Car("Sport")
	motorcycle := Motorcycle("Explorer")
	truck := Truck("Self-Driving")

	sendToLift(car)
	sendToLift(motorcycle)
	sendToLift(truck)
}
