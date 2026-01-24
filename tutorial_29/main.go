package main

import "fmt"

// --- 1. Generics Example ---
// This function works for any type that supports the + operator
type Number interface {
	int | int64 | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

// --- 2. Composition Example ---
type Author struct {
	Name string
	Bio  string
}

type Post struct {
	Title string
	Content string
	Author // Embedding: Post now HAS a Name and Bio
}

func main() {
	// Testing Generics
	fmt.Println("Sum Ints:", Sum(10, 20))
	fmt.Println("Sum Floats:", Sum(1.5, 2.5))

	// Testing Composition
	myPost := Post{
		Title: "Day 29 of Go",
		Content: "Learning about embedding!",
		Author: Author{
			Name: "Harsh",
			Bio:  "A lover of clean code.",
		},
	}

	// We can access Author fields directly on the Post!
	fmt.Printf("Post by: %s\n", myPost.Name) 
}