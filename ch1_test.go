package main

import (
	"math"
	"testing"
)

var tests = []struct {
	id     string
	nums   []int
	min    int
	sum    int
	avg    float64
	median int
	spread float64
}{
	{"1", []int{7, 4, 3, 100, 2343243, 343434, 1, 2, 32}, 1, 2686826, 298536.22222222225, 2343243, 4169539.035981},
	{"2", []int{12, 12, 12}, 12, 36, 12.0, 12, 44.846314},
	{"3", []int{10, 200, 3000, 5000, 4}, 4, 8214, 1642.8, 3000, 11333.099439},
}

func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}

func TestFindMin(t *testing.T) {
	for _, tt := range tests {
		t.Run(string(tt.id), func(t *testing.T) {
			ans := FindMin(tt.nums)
			if ans != tt.min {
				t.Errorf("got %d, want %d", ans, tt.min)
			}
		})
	}
}

func TestFindSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(string(tt.id), func(t *testing.T) {
			ans := FindSum(tt.nums)
			if ans != tt.sum {
				t.Errorf("got %d, want %d", ans, tt.sum)
			}
		})
	}
}

func TestFindAverage(t *testing.T) {
	for _, tt := range tests {
		t.Run(string(tt.id), func(t *testing.T) {
			ans := FindAverage(tt.nums)
			if ans != tt.avg {
				t.Errorf("got %f, want %f", ans, tt.avg)
			}
		})
	}
}

func TestFindMedian(t *testing.T) {
	for _, tt := range tests {
		t.Run(string(tt.id), func(t *testing.T) {
			ans := FindMedian(tt.nums)
			if ans != tt.median {
				t.Errorf("got %d, want %d", ans, tt.median)
			}
		})
	}
}

func TestFindSpread(t *testing.T) {
	for _, tt := range tests {
		t.Run(string(tt.id), func(t *testing.T) {
			ans := FindSpread(tt.nums)
			if !withinTolerance(ans, tt.spread, 1e-7) {
				t.Errorf("got %f, want %f", ans, tt.spread)
			}
		})
	}
}
