// - Sumary: Used a Map to represent a shopping list and do
// Map operation over the Map

package main

import "fmt"

func main() {
	// Creating an empty Map
	shoppingList := make(map[string]int)

	// Adding elements to a Map
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1
	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	// Deleting something on a map
	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	// Checking the existance of a key on a Map
	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup,", cereal, "boxes")
	}

	// Iterating over a Map
	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "on the shopping list")
}
