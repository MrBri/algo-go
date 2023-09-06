package ch2

import (
	"math"
)

func GetFollowerPrediction(followerCount int, influencerType string, numMonths int) int {
	switch influencerType {
	case "fitness":
		return followerCount * int(math.Pow(4, float64(numMonths)))
	case "cosmetic":
		return followerCount * int(math.Pow(3, float64(numMonths)))
	default:
		return followerCount * int(math.Pow(2, float64(numMonths)))
	}
}

func GetInfluencerScore(numFollowers int, avgEngagementPerc float64) int {
	return int(math.Round(avgEngagementPerc * math.Log2(float64(numFollowers))))
}

func Factorial(num int) int {
	total := 1
	for i := 1; i <= num; i++ {
		total *= i
	}
	return total
}
