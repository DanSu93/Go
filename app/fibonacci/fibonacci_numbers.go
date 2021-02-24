// The fibonacci package allows you to compute Fibonacci numbers
// There are 3 type of functions for calculation: recursive (Fibonacci),
// optimized recursive with storing the result of the operation (FibSeriesMemoization),
// iteration in a loop(F)
package fibonacci

// Fibonacci return uint32 value of recursive counting of fibonacci numbers
func Fibonacci(num uint32) uint32 {
	if num <= 1 {
		return num
	} else {
		return Fibonacci(num-1) + Fibonacci(num-2)
	}
}

// FibSeriesMemoization return slice of uint32 values get from
// optimized recursive with storing the result of the operation
func FibSeriesMemoization(num uint32) []uint32 {
	fibonacciSlice := make([]uint32, num)
	fibonacciMap := make(map[uint32]uint32)
	for i := uint32(1); i <= num; i++ {
		fibonacciSlice[i-1] = fibMemo(i, fibonacciMap)
	}
	return fibonacciSlice
}

func fibMemo(num uint32, fMap map[uint32]uint32) uint32 {
	if num <= 1 {
		fMap[num] = num
		return num
	}
	_, ok := fMap[num-1]

	if !ok {
		fMap[num-1] = fibMemo(num-1, fMap)
	}
	_, ok = fMap[num-2]

	if !ok {
		fMap[num-2] = fibMemo(num-2, fMap)
	}
	return fMap[num-1] + fMap[num-2]
}

// F return uint32 value of uint32 get from iteration in a loop
func F(n uint32) uint32 {
	if n <= 0 {
		return 0
	}
	var f1, f2 uint32
	f2 = 1
	for i := uint32(0); i < n-1; i++ {
		sum := f1 + f2
		f1 = f2
		f2 = sum
	}
	return f2
}
