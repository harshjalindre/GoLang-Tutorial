package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func riskyBusiness() {
	defer handlePanic() // Setting up the safety net
	fmt.Println("Executing risky business...")
	panic("Something went horribly wrong!")
	fmt.Println("This line will never run.")
}

func main() {
	fmt.Println("Starting program.")
	
	// Example of Defer for cleanup
	defer fmt.Println("Cleanup complete.")
	
	riskyBusiness()
	
	fmt.Println("Program finished successfully.")
}