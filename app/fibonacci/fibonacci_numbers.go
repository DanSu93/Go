package fibonacci

func fibonacci(num uint32) uint32 {
	if num <= 1 {
		return num
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func fibSeriesMemoization(num uint32) []uint32 {
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
