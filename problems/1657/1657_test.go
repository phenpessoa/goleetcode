package main

import "testing"

func Test1657(t *testing.T) {
	for _, tc := range []struct {
		word1 string
		word2 string
		want  bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"cabbba", "aabbss", false},
		{"uau", "ssx", false},
		{"abbzzca", "babzzcz", false},
	} {
		if got := closeStrings(tc.word1, tc.word2); got != tc.want {
			t.Errorf("\ninvalid result\nword1: %s\nword2: %s\nwant: %v\ngot: %v",
				tc.word1, tc.word2, tc.want, got,
			)
		}
	}
}
