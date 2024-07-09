package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	for _, test := range tests {
		assert.Equal(t, test.output, Sum(test.input))
	}
}

// Benchmark Testing for sum Function
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(5)
	}
}
