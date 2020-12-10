package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var size int

	fmt.Print("\nNeed to generate data for slice. Enter the slice size: ")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Println("It's not a number!")
		os.Exit(1)
	}
	slice := rand.Perm(size) // source slice

	bubbleSortedSlice := bubbleSort(slice)       // new slice after Bubble Sort
	insertionSortedSlice := insertionSort(slice) // new slice after Insertion Sort

	fmt.Println("Source slice: ", slice)
	fmt.Println("Bubble Sort: ", bubbleSortedSlice)
	fmt.Println("Insertion Sort: ", insertionSortedSlice)
}

//Function to sort an array using Bubble Sort
func bubbleSort(a []int) []int {
	l := len(a)
	sliceCopy := make([]int, l)
	copy(sliceCopy, a)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if sliceCopy[j] > sliceCopy[j+1] {
				sliceCopy[j], sliceCopy[j+1] = sliceCopy[j+1], sliceCopy[j]
			}
		}
	}
	return sliceCopy
}

//Function to sort an array using Insertion Sort
func insertionSort(a []int) []int {
	l := len(a)
	sliceCopy := make([]int, l)
	copy(sliceCopy, a)

	for i := 1; i < l; i++ {
		j := i
		for j > 0 {
			if sliceCopy[j-1] > sliceCopy[j] {
				sliceCopy[j-1], sliceCopy[j] = sliceCopy[j], sliceCopy[j-1]
			}
			j = j - 1
		}
	}
	return sliceCopy
}
