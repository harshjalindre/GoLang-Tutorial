package main

import "fmt"

func main() {
	// 1. Create a map of inventory (Product Name -> Quantity)
	inventory := make(map[string]int)

	// 2. Add items
	inventory["Apples"] = 10
	inventory["Bananas"] = 25
	inventory["Oranges"] = 15

	// 3. Update an item
	inventory["Apples"] = 12

	// 4. Delete an item
	delete(inventory, "Oranges")

	// 5. The "Comma OK" check
	val, ok := inventory["Grapes"]
	if ok {
		fmt.Println("Grapes in stock:", val)
	} else {
		fmt.Println("Grapes are out of stock!")
	}

	fmt.Println("Final Inventory:", inventory)
}