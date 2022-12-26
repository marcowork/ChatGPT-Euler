package main

import "fmt"

func main() {
	// Set up a slice to store the prime numbers
	primes := make([]int, 0)

	// Start with the first prime number, 2
	n := 2
	for len(primes) < 10001 {
		// Check if n is prime by checking if it is divisible by any of the prime numbers we have found so far
		isPrime := true
		for _, prime := range primes {
			if n%prime == 0 {
				isPrime = false
				break
			}
		}

		// If n is prime, add it to the slice of prime numbers
		if isPrime {
			primes = append(primes, n)
		}

		// Increment n and continue the search
		n++
	}

	// The 10001st prime number is the last element in the slice
	fmt.Println(primes[len(primes)-1])
}
