package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var x, y, op string
	var fx, fy, result float64

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&x)
	fx = stringToFloat64(x)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&y)
	fy = stringToFloat64(y)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, %): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		result = fx + fy
	case "-":
		result = fx - fy
	case "*":
		result = fx * fy

	case "/":
		if fy == 0 {
			fmt.Println("На 0 делить нельзя!")
			os.Exit(1)
		}
		result = fx / fy
	case "^":
		result = math.Pow(fx, fy)

	case "%":
		result = fx * (fy / 100)

	default:
		fmt.Println("Операция выбрана неверно!")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", result)
}
func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Вы ввели не число!")
		os.Exit(1)
	}
	return f
}
