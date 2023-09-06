package ch2

import (
	"strconv"
	"testing"
)

var tests = []struct {
	followerCount  int
	influencerType string
	numMonths      int
	prediction     int
}{
	{10, "fitness", 1, 40},
	{10, "fitness", 2, 160},
	{10, "fitness", 4, 2560},
	{12, "cosmetic", 4, 972},
	{15, "business", 4, 240},
	{10, "fitness", 5, 10240},
	{10, "fitness", 6, 40960},
}

var testInfluencerScore = []struct {
	numFollowers      int
	avgEngagementPerc float64
	score             int
}{
	{40000, 0.3, 5},
	{43000, 0.1, 2},
	{100000, 0.6, 10},
	{200, 0.8, 6},
}

var testFactorial = []struct {
	num       int
	factorial int
}{
	{2, 2},
	{3, 6},
	{5, 120},
	{7, 5040},
	{9, 362880},
	{11, 39916800},
}

func TestGetFollowerPrediction(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.influencerType, func(t *testing.T) {
			ans := GetFollowerPrediction(tt.followerCount, tt.influencerType, tt.numMonths)
			if ans != tt.prediction {
				t.Errorf("got %d, want %d", ans, tt.prediction)
			}
		})
	}
}

func TestGetInfluencerScore(t *testing.T) {
	for _, tt := range testInfluencerScore {
		t.Run(strconv.Itoa(tt.numFollowers), func(t *testing.T) {
			ans := GetInfluencerScore(tt.numFollowers, tt.avgEngagementPerc)
			if ans != tt.score {
				t.Errorf("got %d, want %d", ans, tt.score)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	for _, tt := range testFactorial {
		t.Run(strconv.Itoa(tt.num), func(t *testing.T) {
			ans := Factorial(tt.num)
			if ans != tt.factorial {
				t.Errorf("got %d, want %d", ans, tt.factorial)
			}
		})
	}
}

// - Initial followers: 200

// - Decay rate: 0.5

// - Days: 1

// Final followers: 100

// =====================================

// - Initial followers: 200

// - Decay rate: 0.5

// - Days: 2

// Final followers: 50

// =====================================

// - Initial followers: 200

// - Decay rate: 0.5

// - Days: 3

// Final followers: 25

// =====================================

// - Initial followers: 1000

// - Decay rate: 0.005

// - Days: 2

// Final followers: 990

// =====================================

// - Initial followers: 1000

// - Decay rate: 0.05

// - Days: 3

// Final followers: 857

// =====================================

// - Initial followers: 1200

// - Decay rate: 0.55

// - Days: 8

// Final followers: 2

// =====================================

// - Initial followers: 1200

// - Decay rate: 0.09

// - Days: 16

// Final followers: 265

// =====================================

// - Initial followers: 200

// - Decay rate: 0.5

// - Days: 1

// Final followers: 100

// =====================================

// - Initial followers: 200

// - Decay rate: 0.5

// - Days: 2

// Final followers: 50

// =====================================

// - Initial followers: 200

// - Decay rate: 0.5

// - Days: 3

// Final followers: 25

// =====================================

// - Initial followers: 1000

// - Decay rate: 0.005

// - Days: 2

// Final followers: 0

// =====================================

// - Initial followers: 1000

// - Decay rate: 0.05

// - Days: 3

// Final followers: 0

// =====================================

// - Initial followers: 1200

// - Decay rate: 0.55

// - Days: 8

// Final followers: 10

// =====================================

// - Initial followers: 1200

// - Decay rate: 0.09

// - Days: 16

// Final followers: 0

// =====================================
