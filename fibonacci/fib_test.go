package fib

import (
	"testing"
)

// Assuming the signature of your Fibonacci function is:
// func fibonacci(n int) int

func TestFibonacciFirstNumber(t *testing.T) {
	result := fibonacci(0)
	expected := 0
	if result != expected {
		t.Errorf("For n=0, expected %d but got %d", expected, result)
	}
}

func TestFibonacciFirstFewNumbers(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{20, 6765},
	}

	for _, test := range tests {
		if output := fibonacci(test.n); output != test.result {
			t.Errorf("For n=%d, expected %d but got %d", test.n, test.result, output)
		}
	}
}

func TestFibonacciLargerNumber(t *testing.T) {
	result := fibonacci(30)
	expected := 832040
	if result != expected {
		t.Errorf("For n=30, expected %d but got %d", expected, result)
	}
}
