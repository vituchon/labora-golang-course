package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    2,
			expected: 2,
		},
		{
			input:    -5,
			expected: 5,
		},
	}

	for _, test := range tests {
		generated := Abs(test.input)
		t.Logf("Testing input=%d, generated=%d (expected=%d)", test.input, generated, test.expected)
		if generated != test.expected {
			t.Errorf("Ufa no coinciden generated=%d, expected=%d", generated, test.expected)
		}
	}
}
