package ch3

import (
	"strconv"
	"testing"
)

var testsMax = []struct {
	nums []int
	max  int
}{
	{[]int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 2343243},
	{[]int{12, 12, 12}, 12},
	{[]int{10, 200, 3000, 5000, 4}, 5000},
}

func TestFindMax(t *testing.T) {
	for _, tt := range testsMax {
		t.Run(strconv.Itoa(tt.max), func(t *testing.T) {
			ans := FindMax(tt.nums)
			if ans != tt.max {
				t.Errorf("got %d, want %d", ans, tt.max)
			}
		})
	}
}
