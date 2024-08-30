package reverse_string

import "testing"

func TestReverse(t *testing.T) {
	for _, test := range testCases {
		if actual := Reverse(test.input); actual != test.expected {
			t.Errorf("Reverse(%q) = %q, want %q", test.input, actual, test.expected)
		}
	}
}

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "an empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "a word",
		input:       "robot",
		expected:    "tobor",
	},
	{
		description: "a capitalized word",
		input:       "Ramen",
		expected:    "nemaR",
	},
	{
		description: "a sentence with punctuation",
		input:       "I'm hungry!",
		expected:    "!yrgnuh m'I",
	},
	{
		description: "a palindrome",
		input:       "racecar",
		expected:    "racecar",
	},
	{
		description: "an even-sized word",
		input:       "drawer",
		expected:    "reward",
	},
}
