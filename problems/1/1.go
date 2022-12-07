package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if idx, ok := m[n]; ok {
			return []int{idx, i}
		}
		m[target-n] = i
	}
	return nil
}
