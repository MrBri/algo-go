package main

import (
	"fmt"
	"math"
)

func FindMin(nums []int) int {
	min := math.MaxInt
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func FindSum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func FindAverage(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}

func FindMedian(nums []int) int {
	mid := int(float64(len(nums) / 2))
	median := nums[mid]
	if len(nums)%2 == 0 {
		median = (nums[mid] + nums[mid+1]) / 2
	}
	return median
}

func FindSpread(nums []int) float64 {
	sum := 0.0
	for _, n := range nums {
		sum += float64(n)
	}
	avg := sum / float64(len(nums))
	return avg * (math.Pow(float64(len(nums)), 1.2))
}

func main() {
	fmt.Println("hello world!")
}
