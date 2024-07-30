package internal

import (
	"testing"
)

func TestAverageCalculator(t *testing.T) {
	tests := []struct {
		name   string
		scores map[string]float64
		want   float64
	}{
		{
			name: "average of three scores",
			scores: map[string]float64{
				"math":    90,
				"science": 80,
				"english": 70,
			},
			want: 80,
		},
		{
			name: "average of one score",
			scores: map[string]float64{
				"math": 100,
			},
			want: 100,
		},
		{
			name:   "average of no scores",
			scores: map[string]float64{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AverageCalculator(tt.scores)
			if got != tt.want {
				t.Errorf("AverageCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
