package primenum

import (
	"math"
)

func findPrimeNumbers(num int64) []int64 {
	var primeArr []int64

	for i := int64(2); i <= num; i++ {
		if isPrime(i) {
			primeArr = append(primeArr, i)
		}
	}
	return primeArr
}

func isPrime(num int64) bool {
	if num < 2 {
		return false
	}
	for i := int64(2); i <= int64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
