//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
  name string
  price int
}

func getTotalNumberOfItems(shoppingList [4]Product) int {
  var totalItems int

  for i := 0; i < len(shoppingList); i++ {
    item := shoppingList[i]

    if item.name != "" {
      totalItems += 1
    }
  }

  return totalItems
}

func getTotalCostOfItems(shoppingList [4]Product) int {
  totalCost := 0
  for i := 0; i < len(shoppingList); i++ {
    item := shoppingList[i]
    totalCost += item.price
  }
  return totalCost
}

func printShoppinListInfo(shoppingList [4]Product) {
  numberOfItems := getTotalNumberOfItems(shoppingList)
  fmt.Println("The last item on the list", shoppingList[numberOfItems -1])
  fmt.Println("The total number of items", numberOfItems)
  fmt.Println("The total cost of the items", getTotalCostOfItems(shoppingList))
}

func main() {

  myshoppingList := [4]Product {
    {name: "Watter", price: 2000},
    {name: "Butter", price: 1500},
    {name: "Milk", price: 1700},
  }

  printShoppinListInfo(myshoppingList)

  myshoppingList[3] = Product{name: "Beer", price: 1100}

  printShoppinListInfo(myshoppingList)
  
}
