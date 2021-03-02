// Copyright 2021 Goltsev M. All rights reserved.
package fizzbuzz

import "strconv"

func Fizzbuzz(num uint32) string {
	if num < 1 {
		return "Type a number >= 1"
	} else if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	} else if num%3 == 0 {
		return "Fizz"
	} else if num%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(int(num))
	}
}
