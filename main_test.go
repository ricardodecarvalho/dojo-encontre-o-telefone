package main

import (
	"testing"
)

func TestFindFone(t *testing.T) {

	tests := []struct {
		about    string
		input    string
		expected string
		err      bool
	}{
		{
			about:    "Print the phone number correctly 1",
			input:    "1-HOME-SWEET-HOME",
			expected: "1-4663-79338-4663",
		},
		{
			about:    "Print the phone number correctly 2",
			input:    "MY-MISERABLE-JOB",
			expected: "69-647372253-562",
		},
		{
			about:    "Print the phone number correctly 3",
			input:    "MY LOVE",
			expected: "69 5683",
		},
		{
			about:    "Print the phone number correctly 4",
			input:    "MY TEST-0",
			expected: "69 8378-0",
		},
		{
			about:    "Print the phone number correctly 5",
			input:    "test-lowercase",
			expected: "8378-569372273",
		},
		{
			about:    "Fail print the phone number correctly",
			input:    "!@#$%*()_+=",
			expected: "invalid character",
		},
	}

	for _, test := range tests {
		t.Run(test.about, func(t *testing.T) {
			fone := FindFone(test.input)
			if fone != test.expected {
				t.Errorf("Expected (%s) response (%s)", test.expected, fone)
			}
		})
	}
}
