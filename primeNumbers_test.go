package main

import "testing"

// Test_primeNumbers defines the functional testing of various usecases of PrimeNumbers function in primeNumbers.go file
func Test_primeNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{name: "3 as input", input: 3, output: 2},
		{name: "0 as input", input: 0, output: 0},
		{name: "1 as input", input: 1, output: 0},
		{name: "4 as input", input: 4, output: 2},
	}
	for _, test := range tests {
		res := PrimeNumbers(test.input)
		if res != test.output {
			t.Errorf("Test %s failed", test.name)
		}
	}
}

// Benchmark Testing for PrimeNumbers Function
func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeNumbers(5)
	}
}
