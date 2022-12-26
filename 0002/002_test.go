package main

import (
	"fmt"
	"testing"
)

func TestSumEvenFibonacci(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"test1", 44},
		{"test2", 4613732},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenFibonacci(); got != tt.want {
				t.Errorf("sumEvenFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
