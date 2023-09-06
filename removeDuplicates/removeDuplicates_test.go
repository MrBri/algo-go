package removeduplicates

import (
	"reflect"
	"testing"
)

// Assuming the signature of your remove duplicates function is:
// func removeDuplicates(data []int) []int

func TestRemoveDuplicatesEmptySlice(t *testing.T) {
	data := []int{}
	expected := []int{}
	result := removeDuplicates(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestRemoveDuplicatesNoDuplicates(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	result := removeDuplicates(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestRemoveDuplicatesWithDuplicates(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 1, 2, 3}
	expected := []int{1, 2, 3, 4, 5}
	result := removeDuplicates(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestRemoveDuplicatesAllIdentical(t *testing.T) {
	data := []int{1, 1, 1, 1, 1}
	expected := []int{1}
	result := removeDuplicates(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
