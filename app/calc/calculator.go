// Package calc provides basic mathematical operations on numbers.
// Support operations: addition, subtraction, multiplication, division, raising to a power, taking a percentage.
// This package does not guarantee bit-identical results across architectures.
package calc

import (
	"errors"
	"math"
)

// Calculate returns the result of an operation on two numbers
// as a floating point value and information about the error, if any.
func Calculate(x, y float64, op rune) (float64, error) {
	var res float64
	switch op {
	case '+':
		res = x + y
	case '-':
		res = x - y
	case '*':
		res = x * y

	case '/':
		if y == 0 {
			return 0, errors.New("cannot be divided by 0")
		}
		res = x / y
	case '^':
		res = math.Pow(x, y)

	case '%':
		res = x * (y / 100)

	default:
		return 0, errors.New("operation specified incorrectly")

	}

	return res, nil
}
