package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 1. Working with Time
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Custom Formatting (The 1-2-3-4-5-6 rule)
	fmt.Println("Formatted:", now.Format("Monday, Jan 2, 2006"))

	// Adding/Subtracting Time
	tomorrow := now.Add(24 * time.Hour)
	fmt.Printf("Tomorrow will be: %v\n", tomorrow.Format("02-01-2006"))

	// 2. Working with the OS
	// Get current working directory
	dir, _ := os.Getwd()
	fmt.Println("Current Directory:", dir)

	// Environment Variables
	user := os.Getenv("USER") // On Windows use "USERNAME"
	if user == "" {
		user = "Gopher"
	}
	fmt.Printf("Hello, %s!\n", user)

	// Command Line Arguments
	fmt.Println("All Arguments:", os.Args)
}