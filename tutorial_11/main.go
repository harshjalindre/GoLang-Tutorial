package main
import "fmt"

func updateValue(val int) {
    val = 100 // Only changes the copy
}

func updateValuePointer(ptr *int) {
    *ptr = 100 // Changes the value at the memory address
}

func main() {
    x := 10
    
    updateValue(x)
    fmt.Println("After copy update:", x) // Still 10

    updateValuePointer(&x) // Pass the address
    fmt.Println("After pointer update:", x) // Now 100!
}