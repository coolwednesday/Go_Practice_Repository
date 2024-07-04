package main

import "testing"

func Test_perimeterRectangle(t *testing.T) {
	type dimensions struct {
		width, height float64
	}
	tests := []struct {
		name   string
		input  dimensions
		output float64
	}{
		{name: "length and breadth are negative", input: dimensions{-1, -1}, output: -1},
		{name: "length is negative", input: dimensions{-1, 4}, output: -1},
		{name: "breadth is negative", input: dimensions{1, -4}, output: -1},
		{name: "Length and breadth are 0", input: dimensions{0, 0}, output: 0},
		{name: "Length and breadth are positive", input: dimensions{3, 2}, output: 10},
	}
	for _, test := range tests {
		res := PerimeterRectangle(test.input.width, test.input.height)
		if res != test.output {
			t.Errorf("Test %s failed", test.name)
		}
	}
}

func Test_perimeterCircle(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		output float64
	}{
		{name: "Radius is negative", input: -1, output: -1},
		{name: "Radius is 0", input: 0, output: 0},
		{name: "Radius is positive", input: 4, output: 25.12},
	}
	for _, test := range tests {
		res := PerimeterCircle(test.input)
		if res != test.output {
			t.Errorf("Test %s failed", test.name)
		}
	}
}

// Benchmark Testing for sum Function
func BenchmarkPerimeterCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PerimeterCircle(5)
	}
}

// Benchmark Testing for sum Function
func BenchmarkPerimeterRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PerimeterRectangle(5, 6)
	}
}
