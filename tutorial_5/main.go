package main
import "fmt"

// Simple function
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Function with multiple returns
func swap(a, b string) (string, string) {
	return b, a
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // returns x and y automatically
}

func main() {
	greet("Harsh")

	// Receiving multiple values
	first, second := swap("World", "Hello")
	fmt.Println(first, second)

	// Using the blank identifier (_) if you want to ignore a return value
	q, _ := split(17)
	fmt.Println("Result q is:", q)
}