package main

import "fmt"
import "math"

// NthPrime returns the prime no. at the nth position.
// If nth must be greater than 0 othwerwise error is returned.
func NthPrime(nth uint) (uint, error) {
	if nth == 0 {
		return 0, fmt.Errorf("please enter nth value > 0")
	} else if nth == 1 {
		return 2, nil
	}

	number, counter := uint(3), uint(1)
	for {
		isPrime := true
		maxDiv := uint(math.Ceil(math.Sqrt(float64(number))))
		for divisor := uint(3); divisor <= maxDiv; divisor++ {
			if number%divisor == 0 {
				isPrime = false
			}
		}

		if isPrime {
			counter++
			if counter == nth {
				return number, nil
			}
		}

		number += 2
	}
}
