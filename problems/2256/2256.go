package main

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
