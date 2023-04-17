package resolution_home_made

import (
	//"reflect"
	"testing"
)

func TestSplitMsInDaysHoursMinutesAndSeconds(t *testing.T) {

	testRuns := []struct {
		title    string
		seconds  int
		expected Result
	}{
		{
			title:   "seconds = 0",
			seconds: 0,
			expected: Result{
				days:    0,
				hours:   0,
				minutes: 0,
				seconds: 0,
			},
		},
		{
			title:   "seconds = 1030",
			seconds: 1030,
			expected: Result{
				days:    0,
				hours:   0,
				minutes: 17,
				seconds: 10,
			},
		},
		{
			title:   "seconds = 12045",
			seconds: 12045,
			expected: Result{
				days:    0,
				hours:   3,
				minutes: 20,
				seconds: 45,
			},
		},
		{
			title:   "seconds = 176520",
			seconds: 176520,
			expected: Result{
				days:    2,
				hours:   1,
				minutes: 2,
				seconds: 0,
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
