package main

import "fmt"

func main() {
	// Task: Calculate the sum of a slice
	nums := []int{2, 4, 6, 8, 10}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of slice:", sum)

	// Task: Find a specific key in a map
	capitals := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
	}

	for country, city := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, city)
	}

	// Fun fact: Range also works on strings!
	for i, char := range "Go" {
		fmt.Printf("Index %d: %c\n", i, char)
	}
}