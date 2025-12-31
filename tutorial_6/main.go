package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Using the 'math' package
	fmt.Println("Pi is:", math.Pi) // Pi is capitalized, so it's exported!
	fmt.Println("Square root of 16:", math.Sqrt(16))

	// Using the 'strings' package
	message := "i love golang"
	fmt.Println(strings.ToUpper(message))
	fmt.Println("Contains 'go'?", strings.Contains(message, "go"))

	// Concept check:
	// If you tried to call math.pi (lowercase), it would fail.
}