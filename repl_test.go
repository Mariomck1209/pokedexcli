package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "  hello  WorlD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			// if they don't match, use t.Errorf to print an error message
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			// and fail the test
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				// if they don't match, use t.Errorf to print an error message
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
				// and fail the test
			}
		}

	}

}
