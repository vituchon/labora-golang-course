package slices

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRotareRightWorks(t *testing.T) {
	str := "ABC"
	expected := "CAB"
	generated := RotateRight(str)
	if generated != expected {
		t.Errorf("No son iguales, generated: %s, expected: %s", generated, expected)
	}

	str = "ABCD"
	expected = "DABC"
	generated = RotateRight(str)
	if generated != expected {
		t.Errorf("No son iguales, generated: %s, expected: %s", generated, expected)
	}

	str = "MNO"
	expected = "OMN"
	generated = RotateRight(str)
	fmt.Println(str, generated)
	if generated != expected {
		t.Errorf("No son iguales, generated: %s, expected: %s", generated, expected)
	}
}

func TestRotareRightWorksTDT(t *testing.T) {
	type Test struct {
		input    string
		expected string
	}

	var tests []Test = []Test{
		{
			input:    "ABC",
			expected: "CAB",
		},
		{
			input:    "MNO",
			expected: "OMN",
		},
		{
			input:    "xyza",
			expected: "axyz",
		},
	}

	for _, test := range tests {
		generated := RotateRight(test.input)
		if generated != test.expected {
			t.Errorf("No son iguales, generated: %s, expected: %s", generated, test.expected)
		}
	}
}

//go test -v ./slices
/*
func TestDesignedForFail(t *testing.T) {
	t.Errorf("Estoy fallandoo")
}*/
func TestSumsWork(t *testing.T) {

	testRuns := []struct {
		title    string
		values   []int
		expected int
	}{
		{
			title:    "[1,2,3]",
			values:   []int{1, 2, 3},
			expected: 6,
		},
		{
			title:    "[1,2,3,4]",
			values:   []int{1, 2, 3, 4},
			expected: 10,
		},
	}
	for _, testRun := range testRuns {
		t.Logf("\n=====Running unit test: %s=====\n", testRun.title)
		generated := SumUsingChannel(testRun.values)
		if testRun.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, testRun.expected)
		}
		generated = SumNotUsingChannel(testRun.values)
		if testRun.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, testRun.expected)
		}

	}
}

// para correr los benchmark
// $slices/go test -benchmem -bench . github.com/vituchon/labora-golang-course/meeting-concurrency/slices
// $slices/go test -bench .  ./slices // todos los benchs
// $slices/go test -bench=BenchmarkSumUsingChannel ./slices // los que sigan el patron
// se puede agregar flag -run=none (para que no ejecute tests!) => slices/go test -run=none  -bench .  ./slices
// https://apuntes.de/golang/pruebas-benchmark/#gsc.tab=0
// https://www.golinuxcloud.com/golang-benchmark/

var slice = rand.Perm(10000000)

func BenchmarkSumUsingChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumUsingChannel(slice)
	}
}

func BenchmarkSumNotUsingChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumNotUsingChannel(slice)
	}
}
