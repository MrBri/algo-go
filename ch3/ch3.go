package ch3

func FindMax(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return max
}
