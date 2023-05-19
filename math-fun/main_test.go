package main

import (
	"reflect"
	"testing"
)

func TestDetermineDivisorsWorks(t *testing.T) {
	type Test struct {
		input    uint
		expected []uint
	}

	tests := []Test{
		Test{
			input:    0,
			expected: []uint{},
		},
		Test{
			input:    1,
			expected: []uint{},
		},
		Test{
			input:    2,
			expected: []uint{1},
		},
		Test{
			input:    6,
			expected: []uint{1, 2, 3},
		},
		Test{
			input:    10,
			expected: []uint{1, 2, 5},
		},
	}
	for _, test := range tests {
		generated := DetermineDivisors(test.input)
		if !reflect.DeepEqual(test.expected, generated) {
			t.Errorf("DetermineDivisors Fails. Input=%v, Generated=%v, Expected=%v", test.input, generated, test.expected)
		}
	}
}

func TestIfPerfectWorks(t *testing.T) {
	type Test struct {
		input    uint
		expected bool
	}

	tests := []Test{
		Test{
			input:    0,
			expected: false,
		},
		Test{
			input:    1,
			expected: false,
		},
		Test{
			input:    2,
			expected: false,
		},
		Test{
			input:    6,
			expected: true,
		},
		Test{
			input:    28,
			expected: true,
		},
		Test{
			input:    496,
			expected: true,
		},
		Test{
			input:    8128,
			expected: true,
		},
	}
	for _, test := range tests {
		generated := IsPerfect(test.input)
		if test.expected != generated {
			t.Errorf("isPerfect Fails. Input=%v, Generated=%v, Expected=%v", test.input, generated, test.expected)
		}
	}
}

func TestAreFriendsWorks(t *testing.T) {
	type Input struct {
		x uint
		y uint
	}
	type Test struct {
		input    Input
		expected bool
	}

	tests := []Test{
		Test{
			input: Input{
				x: 0,
				y: 0,
			},
			expected: false,
		},
		Test{
			input: Input{
				x: 220,
				y: 284,
			},
			expected: true,
		},
		Test{
			input: Input{
				x: 1184,
				y: 1210,
			},
			expected: true,
		},
		Test{
			input: Input{
				x: 1000,
				y: 2000,
			},
			expected: false,
		},
		Test{
			input: Input{
				x: 2620,
				y: 2924,
			},
			expected: true,
		},
		Test{
			input: Input{
				x: 66928,
				y: 66992,
			},
			expected: true,
		},
	}
	for _, test := range tests {
		generated := AreFriends(test.input.x, test.input.y)
		if test.expected != generated {
			t.Errorf("AreFriends Fails. Input=%v, Generated=%v, Expected=%v", test.input, generated, test.expected)
		}
	}
}

func TestAbsWorks(t *testing.T) {
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
