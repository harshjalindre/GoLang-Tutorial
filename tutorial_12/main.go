package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	// Create an instance
	book1 := Book{
		Title:  "The Go Programming Language",
		Author: "Harsh J",
		Pages:  380,
	}

	// Structs inside a Slice
	library := []Book{
		book1,
		{Title: "Clean Code", Author: "H Jalindre", Pages: 464},
	}

	fmt.Println("My Library:")
	for _, book := range library {
		fmt.Printf("- %s by %s (%d pages)\n", book.Title, book.Author, book.Pages)
	}
}