package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Hi Prabhas! Are you ready for tomorrow?    ",
			expected: []string{"hi", "prabhas!", "are", "you", "ready", "for", "tomorrow?"},
		},
		{
			input:    "   aLL excited,           fingers CROSSED    ",
			expected: []string{"all", "excited,", "fingers", "crossed"},
		},
		{
			input:    "   let's DO IT!    ",
			expected: []string{"let's", "do", "it!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of strings do not match, expected: %v, got %v", c.expected, actual)
		}

		for i := range actual {
			a_word := actual[i]
			c_word := c.expected[i]

			if a_word != c_word {
				t.Errorf("Strings do not match, expected: %v, got %v", a_word, c_word)
			}
		}
	}
}
