package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var size int

	fmt.Print("\nNeed to generate data for slice. Enter the slice size: ")
	_, err := fmt.Scanln(&size)
	if err != nil {
		fmt.Println("It's not a number!")
		return
	}
	slice := rand.Perm(size) // source slice

	bubbleSortedSlice := bubbleSort(slice)       // new slice after Bubble Sort
	insertionSortedSlice := insertionSort(slice) // new slice after Insertion Sort

	fmt.Println("Source slice: ", slice)
	fmt.Println("Bubble Sort: ", bubbleSortedSlice)
	fmt.Println("Insertion Sort: ", insertionSortedSlice)
}

//Function to sort an array using Bubble Sort
func bubbleSort(slice []int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if sliceCopy[j] > sliceCopy[j+1] {
				sliceCopy[j], sliceCopy[j+1] = sliceCopy[j+1], sliceCopy[j]
			}
		}
	}
	return sliceCopy
}

//Function to sort an array using Insertion Sort
func insertionSort(slice []int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j--  {
			if sliceCopy[j-1] > sliceCopy[j] {
				sliceCopy[j-1], sliceCopy[j] = sliceCopy[j], sliceCopy[j-1]
			}
		}
	}
	return sliceCopy
}
