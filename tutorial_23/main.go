package main

import (
	"fmt"
	"time"
)

func main() {
	server1 := make(chan string)
	server2 := make(chan string)

	// Simulate Server 1
	go func() {
		time.Sleep(2 * time.Second)
		server1 <- "Response from Server 1 ✅"
	}()

	// Simulate Server 2
	go func() {
		time.Sleep(1 * time.Second)
		server2 <- "Response from Server 2 ✅"
	}()

	// Use select to wait for the FASTEST response
	select {
	case msg1 := <-server1:
		fmt.Println(msg1)
	case msg2 := <-server2:
		fmt.Println(msg2)
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("❌ Error: Operations timed out!")
	}
}