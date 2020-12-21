package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var size int

	fmt.Print("\nNeed to generate data for slice. Enter the slice size: ")
	_, err := fmt.Scanln(&size) // явная передача указателя
	if err != nil {
		fmt.Println("It's not a number!")
		return
	}
	slice := rand.Perm(size) // source slice
	fmt.Println("Source slice: ", slice)

	bubbleSort(slice) // new slice after Bubble Sort
	fmt.Println("Bubble Sort: ", slice)
	insertionSort(slice) // new slice after Insertion Sort
	fmt.Println("Insertion Sort: ", slice)
}

//имеем дело с указателями при передаче параметров, хотя явно их не указываем
//Function to sort an array using Bubble Sort
func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

//Function to sort an array using Insertion Sort
//имеем дело с указателями при передаче параметров, хотя явно их не указываем
func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
		}
	}
}
