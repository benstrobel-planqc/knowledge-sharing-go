package example

import (
	"math"
	"testing"
)

func TestOneCase(t *testing.T) {
	v := math.Abs(-1)
	w := 1.0
	if v != w {
		t.Errorf("Abs(-1) = %v; want %v", v, w)
	}
}

func TestMultipleCases(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
		{-2.5, 2.5},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			v := math.Abs(test.input)
			if v != test.expected {
				t.Errorf("Abs(%v) = %v; want %v", test.input, v, test.expected)
			}
		})
	}
}
