package main

import (
	"fmt"
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{600851475143, 6857},
		{13195, 29},
		{10, 5},
		{7, 7},
	}

	for _, test := range tests {
		result := largestPrimeFactor(test.n)
		if result != test.expected {
			t.Errorf("largestPrimeFactor(%d) = %d, want %d", test.n, result, test.expected)
		}
	}
}
