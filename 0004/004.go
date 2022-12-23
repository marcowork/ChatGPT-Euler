package main

import "fmt"

func isPalindrome(n int) bool {
	// Convert the number to a string and check if it is a palindrome
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	maxPalindrome := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	fmt.Println(maxPalindrome)
}
