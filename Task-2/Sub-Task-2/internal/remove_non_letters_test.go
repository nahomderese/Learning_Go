package internal

import (
	"testing"
)

type testLetters struct {
	input    string
	expected string
}

func TestRemoveNonLetters(t *testing.T) {

	tests := []testLetters{
		{
			input:    "Hello.",
			expected: "Hello",
		},
		{
			input:    "Hello,",
			expected: "Hello",
		},
		{
			input:    "Hello, World",
			expected: "HelloWorld",
		},
	}

	for _, test := range tests {
		actual := removeNonLetters(test.input)
		if actual != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}
