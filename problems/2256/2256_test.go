package main

import "testing"

func Test2256(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{2, 5, 3, 9, 5, 3}, 3},
		{[]int{0}, 0},
		{[]int{0, 1, 0, 1, 0, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{2, 5, 5, 4}, 2},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{8, 2, 6, 18, 10, 6, 20, 15}, 0},
		{[]int{4, 2, 0}, 2},
	} {
		if out := minimumAverageDifference(tc.in); out != tc.want {
			t.Errorf("\ninvalid result\ninput: %v\nwant: %d\ngot: %d", tc.in, tc.want, out)
		}
	}
}
