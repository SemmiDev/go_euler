package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	before := time.Now()
	n := 600851475143
	largestPrimeFactor := largest(primeFactors(n))
	fmt.Println("Nilai faktorisasi prima terbesar : ", largestPrimeFactor)
	after := time.Now()
	fmt.Println("Waktu eksekusi", after.Nanosecond()-before.Nanosecond(), "nano second")
}

func largest(factors []int) int {
	max := factors[0]
	for _, factor := range factors {
		if factor > max {
			max = factor
		}
	}
	return max
}
func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			n /= i
			i--
		}
	}
	if n > 0 {
		factors = append(factors, n)
	}
	return factors
}