package main

import (
	"testing"
)

func Test_sumPairsMethodsDoTheSame(t *testing.T) {
	testRuns := []struct {
		n int
	}{
		{
			n: 0,
		},
		{
			n: 1,
		},
		{
			n: 7,
		},
		{
			n: 21,
		},
		{
			n: 55,
		},
	}
	for _, testRun := range testRuns {
		val1 := sumPairsMethod1(testRun.n)
		val2 := sumPairsMethod2(testRun.n)
		val3 := sumPairsMethod3(testRun.n)
		if !(val1 == val2 && val2 == val3) {
			t.Errorf("Test with %d does not return all same values, they were v1=%d,  v2=%d,  v3=%d", testRun.n, val1, val2, val3)
		}
	}
}

func Benchmark_sumPairsMethod1_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumPairsMethod1(10)
	}
}

func Benchmark_sumPairsMethod2_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumPairsMethod2(10)
	}
}

func Benchmark_sumPairsMethod3_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumPairsMethod3(10)
	}
}
