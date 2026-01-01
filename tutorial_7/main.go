package main

import "fmt"

func main() {
	// 1. Arrays (Fixed size)
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "Array"
	fmt.Println("Array:", arr)

	// 2. Slices (Dynamic size)
	slice := []int{10, 20, 30}
	fmt.Println("Initial Slice:", slice)

	// 3. Adding to a slice
	slice = append(slice, 40, 50)
	fmt.Println("After Append:", slice)

	// 4. Slicing a slice (Getting a sub-section)
	// includes index 1, excludes index 3
	subSlice := slice[1:3] 
	fmt.Println("Sub-slice [1:3]:", subSlice)
	
	// 5. Length and Capacity
	fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))
}