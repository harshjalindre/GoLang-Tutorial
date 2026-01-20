package main

import "testing"

func TestAdd(t *testing.T) {
	// Table of test cases
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -1, -1, -2},
		{"Zero value", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}