package main

import "testing"

func TestHey(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{name: "Empty String", input: "", output: "Fine. Be that way!"},
		{name: "Multiple WhiteSpace Characters", input: "  ", output: "Fine. Be that way!"},
		{name: "Capital Letters with question mark", input: "WHY?", output: "Calm down, I know what I'm doing!"},
		{name: "Capital Letters with no question mark", input: "YOU", output: "Whoa, chill out!"},
		{name: "Numbers and question mark", input: "4?", output: "Sure."},
		{name: "Gibberish", input: "hello there", output: "Whatever."},
		{name: "Only Numbers", input: "12234", output: "Whatever."},
	}

	for i, test := range tests {
		result := Hey(test.input)

		if result != test.output {
			t.Errorf("Test Failed[%d] \n desc: %v got = %q, want %q", i+1, test.name, result, test.output)
		}
	}
}
