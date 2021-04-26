package main

import "testing"

func TestLuhn(t *testing.T) {

	cases := []struct {
		num      string
		expected bool
	}{
		{"0", true},
		{"1", false},
		{"42", true},
		{"79927398713", true},
		{"49927398716", true},
		{"79927398710", false},
		{"49927398716", true},
		{"49927398717", false},
		{"1234567812345678", false},
		{"1234567812345670", true},
		{"89610195012344000018", false},
		{"89610195012344000013", true},
		{"123four567890", false},
	}

	for _, c := range cases {
		result := luhn(c.num)
		if result != c.expected {
			t.Errorf("Did not expect %v to be %v", c.num, result)
		}
	}

}
