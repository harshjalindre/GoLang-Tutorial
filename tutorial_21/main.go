package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup when this function finishes

	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", text, i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Starting Goroutines...")

	// We are starting TWO tasks
	wg.Add(2)

	go printMessage("Task A", &wg)
	go printMessage("Task B", &wg)

	// Wait for all tasks to hit wg.Done()
	wg.Wait()

	fmt.Println("All tasks finished!")
}