package main

import "testing"

// Test_Sum defines the functional testing of Sum function in sum.go file
func Test_Sum(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{name: "0 as input", input: 0, output: 0},
		{name: "-1 as input", input: -1, output: 0},
		{name: "4 as input", input: 4, output: 10},
	}
	for i, test := range tests {
		res := Sum(test.input)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %d, want:%d", i+1, test.name, res, test.output)
		}
	}
}

// Benchmark Testing for sum Function
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(5)
	}
}
