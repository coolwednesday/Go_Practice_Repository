package main

import "testing"

// Functional testing of CanFastAttack Function
func Test_CanFastAttack(t *testing.T) {
	tests := []struct {
		name   string
		input  bool
		output bool
	}{
		{name: "Knight is not awake", input: false, output: true},
		{name: "Knight is awake", input: true, output: false},
	}
	for i, test := range tests {
		res := CanFastAttack(test.input)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %t, want:%t", i+1, test.name, test.output, res)
		}
	}
}

// Functional Testing of CanFreePrisoner Function
func TestCanFreePrisoner(t *testing.T) {
	type status struct {
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		petDogIsPresent bool
	}
	tests := []struct {
		name   string
		input  status
		output bool
	}{
		{name: "None are awake and pet dog is not present", input: status{false, false, false, false}, output: false},
		{name: "Only petDog is Present and none are awake", input: status{false, false, false, true}, output: true},
		{name: "Only Prisoner is awake and pet dog is not present ", input: status{false, false, true, false}, output: true},
		{name: "Only Prisoner is awake and pet dog is present", input: status{false, false, true, true}, output: true},
		{name: "Only archer is awake and dog is not present", input: status{false, true, false, false}, output: false},
		{name: "Only archer is awake and dog is present", input: status{false, true, false, true}, output: false},
		{name: "archer & prisoner are awake and dog is not present", input: status{false, true, true, false}, output: false},
		{name: "archer & prisoner are awake and dog is present", input: status{false, true, true, true}, output: false},
		{name: "Only Knight is awake and dog is not present", input: status{true, false, false, false}, output: false},
		{name: "Knight is awake and dog is present", input: status{true, false, false, true}, output: true},
		{name: "Knight & prisoner are awake and dog is not present", input: status{true, false, true, false}, output: false},
		{name: "Knight & prisoner are awake and dog is present", input: status{true, false, true, true}, output: true},
		{name: "Knight & archer are awake and dog is not present", input: status{true, true, false, false}, output: false},
		{name: "Knight & archer are awake and dog is present", input: status{true, true, false, true}, output: false},
		{name: "All are awake & dog is not present", input: status{true, true, true, false}, output: false},
		{name: "all are awake & dog is present", input: status{true, true, true, true}, output: false},
	}
	for i, test := range tests {
		res := CanFreePrisoner(test.input.knightIsAwake, test.input.archerIsAwake, test.input.prisonerIsAwake, test.input.petDogIsPresent)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %t, want:%t", i+1, test.name, test.output, res)
		}
	}
}

// Functional Testing of CanSpy function
func TestCanSpy(t *testing.T) {
	type status struct {
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
	}
	tests := []struct {
		name   string
		input  status
		output bool
	}{
		{name: "None are awake", input: status{false, false, false}, output: false},
		{name: "Only Prisoner is awake  ", input: status{false, false, true}, output: true},
		{name: "Only archer is awake", input: status{false, true, false}, output: true},
		{name: "archer & prisoner are awake ", input: status{false, true, true}, output: true},
		{name: "Only Knight is awake", input: status{true, false, false}, output: true},
		{name: "Knight & prisoner are awake", input: status{true, false, true}, output: true},
		{name: "Knight & archer are awake", input: status{true, true, false}, output: true},
		{name: "All are awake", input: status{true, true, true}, output: true},
	}
	for i, test := range tests {
		res := CanSpy(test.input.knightIsAwake, test.input.archerIsAwake, test.input.prisonerIsAwake)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %t, want:%t", i+1, test.name, test.output, res)
		}
	}
}

// Functional test for CanSignalPrisoner
func TestCanSignalPrisoner(t *testing.T) {
	type status struct {
		prisonerIsAwake bool
		archerIsAwake   bool
	}
	tests := []struct {
		name   string
		input  status
		output bool
	}{
		{name: "None are awake", input: status{false, false}, output: false},
		{name: "Only archer is awake", input: status{false, true}, output: false},
		{name: "Only prisoner is awake", input: status{true, false}, output: true},
		{name: "Knight & prisoner are awake", input: status{true, true}, output: false},
	}
	for i, test := range tests {
		res := CanSignalPrisoner(test.input.archerIsAwake, test.input.prisonerIsAwake)
		if res != test.output {
			t.Errorf("Test [%d] failed , desc: %v, got: %t, want:%t", i+1, test.name, res, test.output)
		}
	}
}
