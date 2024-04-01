package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("Add dressing")
}

func prepareDishes(dishes []Preparer) {
	for _, dish := range dishes {
		fmt.Printf("-- Dish: %v --\n", dish)
		dish.PrepareDish()
		fmt.Println()
	}
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken wings"),
		Salad("Roman salad"),
	}
	prepareDishes(dishes)
}
