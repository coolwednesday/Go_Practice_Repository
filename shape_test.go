package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetDetails(t *testing.T) {
	tests := []struct {
		name   string
		input  shape
		output string
	}{
		//Test Cases for Rectangle struct
		{name: "Rectangle with positive dimensions", input: rectangle{length: 10, breadth: 8}, output: "Shape Details => Area: 80.0 Perimeter: 36.0"},
		{name: "Rectangle with negative dimensions", input: rectangle{length: -10, breadth: -8}, output: "Shape Details => Area: -1.0 Perimeter: -1.0"},
		{name: "Rectangle with zero values", input: rectangle{length: 0, breadth: 0}, output: "Shape Details => Area: 0.0 Perimeter: 0.0"},
		//Test Cases for Circle struct
		{name: "Circle with positive radius", input: circle{radius: 10}, output: "Shape Details => Area: 314.0 Perimeter: 62.8"},
		{name: "Circle with negative radius", input: circle{radius: -5}, output: "Shape Details => Area: -1.0 Perimeter: -1.0"},
		{name: "Circle with zero radius", input: circle{radius: 0}, output: "Shape Details => Area: 0.0 Perimeter: 0.0"},
		//Test Cases for Triangle struct
		{name: "Triangle with negative dimension", input: triangle{base: -1, height: -1}, output: "Shape Details => Area: -1.0 Perimeter: -1.0"},
		{name: "Triangle with zero values", input: triangle{base: 0, height: 0}, output: "Shape Details => Area: 0.0 Perimeter: 0.0"},
		{name: "Triangle with positive dimensions", input: triangle{base: 4, height: 3}, output: "Shape Details => Area: 6.0 Perimeter: 12.0"},
		//Test Cases for Square struct
		{name: "Square with positive dimensions", input: square{length: 10}, output: "Shape Details => Area: 100.0 Perimeter: 40.0"},
		{name: "Square with negative dimensions", input: square{length: -10}, output: "Shape Details => Area: -1.0 Perimeter: -1.0"},
		{name: "Square with zero values", input: square{length: 0}, output: "Shape Details => Area: 0.0 Perimeter: 0.0"},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, getDetails(test.input))
	}
}

// Benchmark Testing for SayHello Function
func BenchmarkgetDetails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDetails(rectangle{length: 10, breadth: 8})
	}
}
