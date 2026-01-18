package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
	InStock bool  `json:"in_stock"`
}

func main() {
	// 1. Encoding (Struct -> JSON)
	p1 := Product{Name: "Gopher Mug", Price: 12.99, InStock: true}
	
	bytes, _ := json.MarshalIndent(p1, "", "  ") // Pretty print
	fmt.Println("JSON Output:")
	fmt.Println(string(bytes))

	// 2. Decoding (JSON -> Struct)
	jsonInput := `{"product_name": "Go Shirt", "price": 25.00, "in_stock": false}`
	var p2 Product

	err := json.Unmarshal([]byte(jsonInput), &p2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nDecoded Struct: %+v\n", p2)
}