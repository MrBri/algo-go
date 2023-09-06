package fib

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	nums := []int{0, 1}
	for i := 2; i <= n; i++ {
		nums[0], nums[1] = nums[1], (nums[0] + nums[1])
	}
	return nums[1]
}
