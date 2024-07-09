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
		{name: "Rectangle", input: rectangle{length: 10, breadth: 8}, output: "Shape Details => Area: 80.0 Perimeter: 36.0"},
		{name: "Circle", input: circle{radius: 10}, output: "Shape Details => Area: 314.0 Perimeter: 62.8"},
		{name: "Triangle", input: triangle{base: 4, height: 3}, output: "Shape Details => Area:6.0  Perimeter: 12.0"},
		{name: "Square", input: square{length: 10}, output: "Shape Details => Area: 100.0 Perimeter: 40.0"},
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
