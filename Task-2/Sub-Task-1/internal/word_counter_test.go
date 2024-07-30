package internal

import (
	"testing"
)

type frequencyTest struct {
	input    string
	expected map[string]int64
}

func TestWordCounter(t *testing.T) {

	tests := []frequencyTest{
		{
			input: "Hello World",
			expected: map[string]int64{
				"hello": 1,
				"world": 1,
			},
		},
		{
			input: "Hello World hello",
			expected: map[string]int64{
				"hello": 2,
				"world": 1,
			},
		},
		{
			input: "Hello, Worl.d Hel.lo Worl.d",
			expected: map[string]int64{
				"hello": 2,
				"world": 2,
			},
		},
		{
			input: "Hello, Worl.d Hel.lo Worl.d",
			expected: map[string]int64{
				"hello": 2,
				"world": 2,
			},
		},
		{
			input: "Hello, world World hello",
			expected: map[string]int64{
				"hello": 2,
				"world": 2,
			},
		},
	}

	for _, test := range tests {
		actual := WordCounter(test.input)
		if !compareMaps(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}

}

func compareMaps(a, b map[string]int64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bVal, ok := b[k]; !ok || bVal != v {
			return false
		}
	}
	return true
}
