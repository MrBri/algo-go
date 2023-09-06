package binarysearch

import "math"

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := int(math.Round(float64((l + r) / 2)))
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
