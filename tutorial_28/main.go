package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) {
	fmt.Println("Slow operation started...")
	
	select {
	case <-time.After(5 * time.Second):
		// Simulate a long task
		fmt.Println("Operation finished successfully!")
	case <-ctx.Done():
		// If the context is cancelled or times out
		fmt.Println("Operation cancelled:", ctx.Err())
	}
}

func main() {
	// 1. Create a context that times out after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Important: Always clean up the context!

	// 2. Pass context to the goroutine
	go slowOperation(ctx)

	// Wait long enough to see the result
	time.Sleep(3 * time.Second)
	fmt.Println("Main program finished.")
}