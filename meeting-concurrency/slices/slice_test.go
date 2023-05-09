package slices

import (
	"math/rand"
	"testing"
)

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
// $slices/go test -benchmem -run=^$ -bench=BenchmarkSum*  github.com/vituchon/labora-golang-course/meeting-concurrency/slices

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
