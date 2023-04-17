package home

import (
	//"reflect"
	"testing"

	"github.com/vituchon/labora-golang-course/milis-to-time-units/resolution"
)

func TestSplitMsInDaysHoursMinutesAndSeconds(t *testing.T) {

	testRuns := []struct {
		title    string
		seconds  int
		expected resolution.Result
	}{
		{
			title:   "seconds = 0",
			seconds: 0,
			expected: resolution.Result{
				Days:    0,
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			},
		},
		{
			title:   "seconds = 1030",
			seconds: 1030,
			expected: resolution.Result{
				Days:    0,
				Hours:   0,
				Minutes: 17,
				Seconds: 10,
			},
		},
		{
			title:   "seconds = 12045",
			seconds: 12045,
			expected: resolution.Result{
				Days:    0,
				Hours:   3,
				Minutes: 20,
				Seconds: 45,
			},
		},
		{
			title:   "seconds = 176520",
			seconds: 176520,
			expected: resolution.Result{
				Days:    2,
				Hours:   1,
				Minutes: 2,
				Seconds: 0,
			},
		},
	}
	for _, testRun := range testRuns {
		t.Logf("\n=====Running unit test: %s=====\n", testRun.title)
		generated := SplitSecondsInDaysHoursMinutesAndSeconds(testRun.seconds)
		if testRun.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, testRun.expected)
		}
	}
}
