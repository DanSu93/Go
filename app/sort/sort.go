package sort

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
func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
		}
	}
}
