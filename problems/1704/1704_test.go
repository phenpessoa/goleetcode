package main

import "testing"

func Test1704(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want bool
	}{
		{"book", true},
		{"textbook", false},
	} {
		if got := halvesAreAlike(tc.in); got != tc.want {
			t.Errorf("\ninvalid result\nin: %s\nwant: %v\ngot: %v", tc.in, tc.want, got)
		}
	}
}
