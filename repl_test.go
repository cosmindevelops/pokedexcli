package main

import "testing"

func TestCleanInput(t *testing.T) {
	// this test is to make sure that the cleanInput function is working as expected
	// cleanInput is a function that takes a string as input and returns an array of strings
	// this test is checking that the function is working correctly by passing in a variety of test cases
	// and verifying that the output is what is expected

	cases := []struct {
		input    string // the input to the cleanInput function
		expected []string
	}{
		{
			input:    "hello world",              // input string
			expected: []string{"hello", "world"}, // expected output
		},
		{
			input:    "HELLO world",
			expected: []string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input) // call the cleanInput function with the input string
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v and %v", len(actual), len(cs.expected)) // if the lengths are not equal
			continue                                                                        // print an error message and move on to the next test case
		}
		for i := range actual {
			actualWord := actual[i] // get the i-th word of the actual output
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v is not equal to %v", actualWord, expectedWord) // if the word is not equal to the expected word
			}
		}
	}
}
