package main

import (
	"strconv"
	"testing"
)

func TestFindGreatestProduct(t *testing.T) {
	expected := 23514624000
	result := findGreatestProduct()
	if result != expected {
		t.Errorf("Test failed: expected %d, got %d", expected, result)
	}
}

func findGreatestProduct() int {
	maxProduct := 0
	for i := 0; i < len(number)-13; i++ {
		product := 1
		for j := i; j < i+13; j++ {
			n, _ := strconv.Atoi(string(number[j]))
			product *= n
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct
}
