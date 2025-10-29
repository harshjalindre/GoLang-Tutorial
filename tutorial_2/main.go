package main
import "fmt"

func main() {
	var intNum int = 42
	var floatNum float64 = 3.14
	var str string = "GoLang"
	var isActive bool = true
	
	// Print values
	fmt.Println("Integer:", intNum)
	fmt.Println("Float:", floatNum)
	fmt.Println("String:", str)
	fmt.Println("Boolean:", isActive)

	intNum2 := 100  // Short declaration
	fmt.Println("Another Integer:", intNum2)

	var int1, int2, int3 int = 1, 2, 3 // Multiple variable declaration
	fmt.Println("Multiple Integers:", int1, int2, int3)

	var1, var2, var3 := "A", "B", "C" // Multiple short declaration
	fmt.Println("Multiple Strings:", var1, var2, var3)

}