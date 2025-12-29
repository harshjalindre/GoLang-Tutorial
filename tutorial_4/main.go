package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. If statement 
	if num := 10; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is a single digit")
	} else {
		fmt.Println(num, "is 10 or greater")
	}

	// 2. Standard Switch (No 'break' needed!)
	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	case "Monday":
		fmt.Println("Back to work.")
	default:
		fmt.Println("Just another weekday.")
	}

	// 3. Conditionless Switch (Perfect for complex logic)
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// 4. The 'fallthrough' keyword
	// Occasionally, you WANT the next case to run.
	val := 1
	switch val {
	case 1:
		fmt.Println("Case 1 matches.")
		fallthrough
	case 2:
		fmt.Println("This also executes because of fallthrough!")
	}
}