package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	expectedSum := 233168
	sum := sumMultiples()
	if sum != expectedSum {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expectedSum)
	}
}
