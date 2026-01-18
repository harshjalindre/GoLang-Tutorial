package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Activity struct {
	Type     string `json:"type"`
	Activity string `json:"activity"`
}

func main() {
	url := "https://bored-api.appbrewery.com/random"

	// 1. Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	// 2. ALWAYS close the body
	defer resp.Body.Close()

	// 3. Read the raw body
	body, _ := io.ReadAll(resp.Body)

	// 4. Unmarshal JSON into our struct
	var suggestion Activity
	json.Unmarshal(body, &suggestion)

	fmt.Printf("Bored? Try this %s activity: %s\n", suggestion.Type, suggestion.Activity)
}