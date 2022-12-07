package main

import "testing"

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	for _, tc := range []struct {
		args args
		want int
	}{
		{
			args: args{
				root: newTree(10, 5, 15, 3, 7, 18),
				low:  7,
				high: 15,
			},
			want: 32,
		},
		{
			args: args{
				root: newTree(10, 5, 15, 3, 7, 13, 18, 1, 6),
				low:  6,
				high: 10,
			},
			want: 23,
		},
		{
			args: args{
				root: nil,
				low:  0,
				high: 0,
			},
			want: 0,
		},
	} {
		if got := rangeSumBST(tc.args.root, tc.args.low, tc.args.high); got != tc.want {
			t.Errorf("\nrangeSumBST() = %v, want %v", got, tc.want)
		}
	}
}

func newTree(nums ...int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	out := new(TreeNode)
	for _, n := range nums {
		out.insert(n)
	}
	return out
}
