package binarysearch

import (
	"testing"
)

// Assuming the signature of your binary search function is:
// func binarySearch(data []int, value int) int
// It should return the index of the value in the slice if found, or -1 if not found.

func TestBinarySearchEmptySlice(t *testing.T) {
	data := []int{}
	result := binarySearch(data, 5)
	if result != -1 {
		t.Errorf("Expected -1, but got %d", result)
	}
}

func TestBinarySearchValuePresent(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(data, 5)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

func TestBinarySearchValueNotPresent(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(data, 11)
	if result != -1 {
		t.Errorf("Expected -1, but got %d", result)
	}
}

func TestBinarySearchMultipleOccurrences(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 5, 5, 8, 9, 10}
	result := binarySearch(data, 5)
	if result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}
}

var data = []struct {
	name      string
	followers int
}{
	{"John", 5},
	{"Jane", 10},
	{"James", 15},
	{"Judy", 20},
	{"Jenny", 25},
	{"Jasper", 30},
	{"Jack", 35},
	{"Jill", 40},
	{"Bob", 45},
	{"Borice", 50},
	{"Boris", 55},
	{"Donald", 60},
	{"Doris", 65},
	{"Derek", 70},
	{"Diana", 75},
	{"Dennis", 80},
	{"Daisy", 85},
	{"Duke", 90},
	{"Duke", 95},
	{"Duke", 100},
}

func TestBinarySearchStruct(t *testing.T) {
	toSearch := []struct {
		followers int
		name      string
	}{{50, "Borice"}, {40, "Jill"}, {30, "Jasper"}, {20, "Judy"}, {5000, ""}}
	for _, tt := range toSearch {
		if ans := BinarySearchStruct(data, tt.followers); ans != tt.name {
			t.Errorf("expected %s, got %s", tt.name, ans)
		}
	}
}
