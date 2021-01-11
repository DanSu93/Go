package primenum

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func findPrimeNumbers() {
	var n string
	var in, i int64
	var primeArr []int64

	fmt.Print("Введите  число N: ")
	fmt.Scanln(&n) // явная передача указателя
	in, err := strconv.ParseInt(n, 0, 64)
	if err != nil {
		fmt.Println("Вы ввели не число!")
		os.Exit(1)
	}

	for i = 0; i <= in; i++ {
		if isPrime(i) {
			primeArr = append(primeArr, i)
		}
	}

	fmt.Println(primeArr)

}

func isPrime(num int64) bool {
	for i := int64(2); i <= int64(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
