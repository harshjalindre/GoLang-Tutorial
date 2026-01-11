package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. Advanced Formatting
	name := "Harsh"
	age := 15
	fmt.Printf("Default: %v, Type: %T, Quoted: %q\n", name, name, name)
	fmt.Printf("Age in binary: %b\n", age)

	// 2. Reading from a String (Simulated Input)
	reader := strings.NewReader("Go is awesome!")
	
	// 3. Writing to standard output (Console)
	// os.Stdout implements the io.Writer interface
	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		fmt.Println("\nError reading:", err)
	}

	// 4. Writing to a file
	content := "Learning Day 18: fmt and io"
	err = os.WriteFile("day18.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("File error:", err)
	}
	fmt.Println("\nFile 'day18.txt' created successfully!")
}