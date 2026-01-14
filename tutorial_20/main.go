package main

import (
	"fmt"
	"os"
)

// 1. Define the Interface
type WeatherProvider interface {
	GetWeather(city string) WeatherData
}

// 2. Define the Data Struct
type WeatherData struct {
	City        string
	Temperature float64
	Condition   string
}

// 3. Method to format output
func (w WeatherData) Display() {
	fmt.Printf("☁️ Weather in %s:\n", w.City)
	fmt.Printf("- Temp: %.1f°C\n", w.Temperature)
	fmt.Printf("- Info: %s\n", w.Condition)
}

// 4. Concrete Provider (Mock API)
type MockWeatherAPI struct{}

func (m MockWeatherAPI) GetWeather(city string) WeatherData {
	// In a real app, this would be an HTTP call
	return WeatherData{
		City:        city,
		Temperature: 22.5,
		Condition:   "Partly Cloudy",
	}
}

func main() {
	// Check for CLI arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [city-name]")
		return
	}
	city := os.Args[1]

	// Use the Interface
	var provider WeatherProvider = MockWeatherAPI{}
	data := provider.GetWeather(city)

	data.Display()
}