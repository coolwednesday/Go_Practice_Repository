package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_SayHello defines the functional testing of SayHello function in airport_robot_test.go file
func Test_SayHello(t *testing.T) {
	type parameter struct {
		name string
		t    greeter
	}
	tests := []struct {
		name   string
		input  parameter
		output string
	}{
		{name: "Empty String & Italian Type", input: parameter{"", Italian{}}, output: "I can speak Italian: Ciao !"},
		{name: "Full Name & Italian Type", input: parameter{"Divya Darshana", Italian{}}, output: "I can speak Italian: Ciao Divya Darshana!"},
		{name: "Only First Name & Italian Type", input: parameter{"Divya", Italian{}}, output: "I can speak Italian: Ciao Divya!"},
		{name: "Empty & Portuguese Type", input: parameter{"", Portuguese{}}, output: "I can speak Portuguese: Olá !"},
		{name: "Only First Name & Portuguese Type", input: parameter{"Divya", Portuguese{}}, output: "I can speak Portuguese: Olá Divya!"},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, SayHello(test.input.name, test.input.t))
	}
}

// BenchmarkSayHello Benchmark Testing for SayHello Function
func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Divya", Italian{})
	}
}
