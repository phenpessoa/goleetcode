package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	for _, tc := range []struct {
		want []int
		args args
	}{
		{
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	} {
		if got := twoSum(tc.args.nums, tc.args.target); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("\ntwoSum() = %v, want %v", got, tc.want)
		}
	}
}
