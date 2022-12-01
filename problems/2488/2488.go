package main

import (
	"fmt"
)

func countSubarrays(nums []int, k int) int {
	var idx int
	for i, n := range nums {
		if n == k {
			idx = i
			break
		}
	}

	left := make([]int, idx+1)
	for i := idx - 1; i > -1; i-- {
		if nums[i] > k {
			left[i] = left[i+1] + 1
			continue
		}
		left[i] = left[i+1] - 1
	}

	m := make(map[int]int, len(nums)-idx-1)
	var x int
	for i := idx + 1; i < len(nums); i++ {
		switch i {
		case idx + 1:
			if nums[i] > k {
				x = 1
				break
			}
			x = -1
		default:
			if nums[i] > k {
				x++
				break
			}
			x--
		}
		m[x]++
	}

	var count int
	for i := 0; i < idx+1; i++ {
		if left[i] == 0 || left[i] == 1 {
			count++
		}

		count += m[-1*left[i]]
		count += m[-1*left[i]+1]
	}

	return count
}

func main() {
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))    // 3
	fmt.Println(countSubarrays([]int{2, 3, 1}, 3))          // 1
	fmt.Println(countSubarrays([]int{2, 5, 1, 4, 3, 6}, 1)) // 3
	fmt.Println(countSubarrays([]int{1}, 1))                // 1
	fmt.Println(countSubarrays([]int{2, 1, 4, 3}, 2))       // 3
}
