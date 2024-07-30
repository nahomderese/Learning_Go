package internal

import (
	"strings"
	"testing"
)

func TestIntInput(t *testing.T) {
	tests := []struct {
		name   string
		prompt string
		input  string
		want   int
	}{
		{
			name:   "valid input",
			prompt: "Enter a number: ",
			input:  "42\n",
			want:   42,
		},
		{
			name:   "invalid then valid input",
			prompt: "Enter a number: ",
			input:  "abc\n42\n",
			want:   42,
		},
		{
			name:   "negative number",
			prompt: "Enter a number: ",
			input:  "-42\n-10\n42",
			want:   42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got := intInput(tt.prompt, reader)
			if got != tt.want {
				t.Errorf("intInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatInput(t *testing.T) {
	tests := []struct {
		name   string
		prompt string
		input  string
		want   float64
	}{
		{
			name:   "valid input",
			prompt: "Enter a number: ",
			input:  "42.5\n",
			want:   42.5,
		},
		{
			name:   "invalid then valid input",
			prompt: "Enter a number: ",
			input:  "abc\n42.5\n",
			want:   42.5,
		},
		{
			name:   "negative number",
			prompt: "Enter a number: ",
			input:  "-42.5\n-10.2\n2.5",
			want:   2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got := floatInput(tt.prompt, reader)
			if got != tt.want {
				t.Errorf("floatInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInput(t *testing.T) {
	tests := []struct {
		name   string
		prompt string
		input  string
		want   string
	}{
		{
			name:   "valid input",
			prompt: "Enter a string: ",
			input:  "hello\n",
			want:   "hello",
		},
		{
			name:   "invalid then valid input",
			prompt: "Enter a string: ",
			input:  "\nhello\n",
			want:   "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got := stringInput(tt.prompt, reader)
			if got != tt.want {
				t.Errorf("stringInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
