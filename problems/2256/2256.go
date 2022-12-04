package main

import (
	"fmt"
)

func minimumAverageDifference(nums []int) int {
	var (
		smallest int = 1<<63 - 1
		idx      int
		xs       = make([]int, len(nums))
	)

	for i, n := range nums {
		if i == 0 {
			xs[i] = n
			continue
		}
		xs[i] = n + xs[i-1]
	}

	for i := range nums {
		a := len(nums) - i - 1
		x := xs[i]
		y := i + 1
		z := xs[len(nums)-1] - xs[i]

		var avg int
		if a > 0 {
			avg = (x / y) - (z / a)
		} else {
			avg = (x / y) - z
		}

		if avg < 0 {
			avg *= -1
		}

		if avg < smallest {
			idx = i
			smallest = avg
		}
	}

	return idx
}

func main() {
	fmt.Println(minimumAverageDifference([]int{2, 5, 3, 9, 5, 3}))           // 3
	fmt.Println(minimumAverageDifference([]int{0}))                          // 0
	fmt.Println(minimumAverageDifference([]int{0, 1, 0, 1, 0, 1}))           // 0
	fmt.Println(minimumAverageDifference([]int{1, 2, 3, 4, 5}))              // 0
	fmt.Println(minimumAverageDifference([]int{2, 5, 5, 4}))                 // 2
	fmt.Println(minimumAverageDifference([]int{5, 4, 3, 2, 1}))              // 1
	fmt.Println(minimumAverageDifference([]int{8, 2, 6, 18, 10, 6, 20, 15})) // 0
	fmt.Println(minimumAverageDifference([]int{4, 2, 0}))                    // 2
}
