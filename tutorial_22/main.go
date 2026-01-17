package main

import "fmt"

func sendData(ch chan string) {
	// Send a message into the channel
	ch <- "Hello from the background! 🚀"
}

func main() {
	// 1. Create a channel
	myChannel := make(chan string)

	// 2. Start a goroutine
	go sendData(myChannel)

	// 3. Receive the message (This line waits until data is sent)
	message := <-myChannel

	fmt.Println("Received:", message)
}