package main

import "testing"

// Functional test for remaining oven time (Cannot take a negative number as describes the time spent in oven)
func Test_RemainingOvenTime(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{name: "0 as input", input: 0, output: 40},
		{name: "Positive Number as input", input: 6, output: 34},
	}
	for i, test := range tests {
		res := RemainingOvenTime(test.input)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %d, want:%d", i+1, test.name, res, test.output)
		}
	}
}

// Functional test for PreparationTime
func Test_PreparationTime(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{name: "0 as input", input: 0, output: 0},
		{name: "Negative number as input", input: -4, output: -1},
		{name: "Positive Number as input", input: 4, output: 8},
	}
	for i, test := range tests {
		res := PreparationTime(test.input)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %d, want:%d", i+1, test.name, res, test.output)
		}
	}
}
func Test_ElapsedTime(t *testing.T) {
	type totalTime struct {
		numberOfLayers      int
		actualMinutesInOven int
	}
	tests := []struct {
		name   string
		input  totalTime
		output int
	}{
		{name: "totalTime is 0", input: totalTime{0, 0}, output: 0},
		{name: "numberOfLayers is negative", input: totalTime{-1, 2}, output: -1},
		{name: "actualMinutesInOven is negative", input: totalTime{2, -3}, output: -1},
		{name: "Both actualMinutesInOven & numberOfLayers are negative", input: totalTime{-2, -5}, output: -1},
		{name: "actualMinutesInOven & numberOfLayers are positive numbers", input: totalTime{2, 5}, output: 9},
	}
	for i, test := range tests {
		res := ElapsedTime(test.input.numberOfLayers, test.input.actualMinutesInOven)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %d, want:%d", i+1, test.name, res, test.output)
		}
	}

}

// Benchmark Testing for ElapsedTime Function
func BenchmarkElapsedTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ElapsedTime(2, 5)
	}
}

// Benchmark Testing for  Function
func BenchmarkPreparationTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreparationTime(2)
	}
}

// Benchmark Testing for sum Function
func BenchmarkRemainingOvenTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemainingOvenTime(6)
	}
}
