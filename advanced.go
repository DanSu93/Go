package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
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
		if isSimple(i) {
			primeArr = append(primeArr, i)
		}
	}

	fmt.Println(primeArr)

}

func isSimple(num int64) bool {
	var i int64
	var sqrtNum = int64(math.Sqrt(float64(num)))
	for i = 2; i < sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
