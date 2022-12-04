package main

import "testing"

func Test2488(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 4, 5}, 4, 3},
		{[]int{2, 3, 1}, 3, 1},
		{[]int{2, 5, 1, 4, 3, 6}, 1, 3},
		{[]int{1}, 1, 1},
		{[]int{2, 1, 4, 3}, 2, 3},
	} {
		if out := countSubarrays(tc.in, tc.k); out != tc.want {
			t.Errorf("\ninvalid result\ninput: %v | k: %d\nwant: %d\ngot: %d", tc.in, tc.k, tc.want, out)
		}
	}
}
