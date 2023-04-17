package reversor

import (
	"testing"
)

func TestReversesWork(t *testing.T) {

	testRuns := []struct {
		title    string
		text     string
		expected string
	}{
		{
			title:    "(empty string) => ''",
			text:     "",
			expected: "",
		},
		{
			title:    "'a' => 'a'",
			text:     "a",
			expected: "a",
		},
		{
			title:    "'ab' => 'ba'",
			text:     "ab",
			expected: "ba",
		},
		{
			title:    "abc' => 'cba'",
			text:     "abc",
			expected: "cba",
		},
	}
	for _, testRun := range testRuns {
		t.Logf("\n=====Running unit test: %s=====\n", testRun.title)
		generated := ReverseRunes(testRun.text)
		if testRun.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, testRun.expected)
		}
		generated = Reverse(testRun.text)
		if testRun.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, testRun.expected)
		}

		generated = InvertWords(testRun.text)
		if testRun.expected != generated {
			t.Errorf("\nGenerated: '%+v' vs Expected: '%+v'", generated, testRun.expected)
		}
	}
}

func BenchmarkReverseRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseRunes("abc")
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abc")
	}
}

func BenchmarkInvertWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InvertWords("abc")
	}
}
