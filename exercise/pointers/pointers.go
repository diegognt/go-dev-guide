//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func stats(items []Item) {
	fmt.Println("\n----- List of items -----")
	for _, item := range items {
		tag := "inactivate"

		if item.tag {
			tag = "activate"
		}

		fmt.Println("The product", item.name, "is:", tag)
	}
}

func main() {
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	items := []Item{
		{name: "Milk", tag: Active},
		{name: "Flour", tag: Active},
		{name: "Bottle of water", tag: Active},
		{name: "Coffee", tag: Active},
	}
	stats(items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[2].tag)
	stats(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	stats(items)
}
