package main

import (
	"reflect"
	"testing"
)

func TestDetermineDivisorsWorks(t *testing.T) {
	type Test struct {
		input    int
		expected []int
	}

	tests := []Test{
		Test{
			input:    0,
			expected: []int{},
		},
		Test{
			input:    1,
			expected: []int{},
		},
		Test{
			input:    2,
			expected: []int{1},
		},
		Test{
			input:    6,
			expected: []int{1, 2, 3},
		},
		Test{
			input:    10,
			expected: []int{1, 2, 5},
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
		input    int
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
		generated := isPerfect(test.input)
		if test.expected != generated {
			t.Errorf("isPerfect Fails. Input=%v, Generated=%v, Expected=%v", test.input, generated, test.expected)
		}
	}
}

func TestAreFriendsWorks(t *testing.T) {
	type Input struct {
		x int
		y int
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
