package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Oops:", err)
	} else {
		fmt.Println("Result:", result)
	}
}