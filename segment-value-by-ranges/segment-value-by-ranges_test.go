package main

import (
	//"reflect"
	"testing"
)

func TestSegmentarValorPorRangosFixed(t *testing.T) {
	testRuns := []struct {
		title    string
		value    int
		expected struct {
			s1, s2, s3, s4, s5 int
		}
	}{
		{
			title: "x = 0",
			value: 0,
			expected: struct {
				s1, s2, s3, s4, s5 int
			}{
				s1: 0,
				s2: 0,
				s3: 0,
				s4: 0,
				s5: 0,
			},
		},
		{
			title: "x = 1500",
			value: 1500,
			expected: struct {
				s1, s2, s3, s4, s5 int
			}{
				s1: 50,
				s2: 50,
				s3: 600,
				s4: 800,
				s5: 0,
			},
		},
	}
	for _, testRun := range testRuns {
		t.Logf("\n=====Running unit test: %s=====\n", testRun.title)
		s1, s2, s3, s4, s5 := SegmentarValorPorRangos(testRun.value)
		if testRun.expected.s1 != s1 ||
			testRun.expected.s2 != s2 ||
			testRun.expected.s3 != s3 ||
			testRun.expected.s4 != s4 ||
			testRun.expected.s5 != s5 {
			t.Errorf("\nGenerated:{s1:%d s2:%d s3:%d s4:%d s5:%d}\nExpected:  %+v", s1, s2, s3, s4, s5, testRun.expected)
		}
	}
}
