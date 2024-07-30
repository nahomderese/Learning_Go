package internal

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "simple palindrome",
			input: "madam",
			want:  true,
		},
		{
			name:  "simple non-palindrome",
			input: "hello",
			want:  false,
		},
		{
			name:  "palindrome with mixed case",
			input: "MadAm",
			want:  true,
		},
		{
			name:  "palindrome with spaces and punctuation",
			input: "A man, a plan, a canal, Panama!",
			want:  true,
		},
		{
			name:  "non-palindrome with spaces and punctuation",
			input: "Hello, World!",
			want:  false,
		},
		{
			name:  "empty string",
			input: "",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
