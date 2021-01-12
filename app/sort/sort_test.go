package sort

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	name  string
	input []int
	want  []int
}

var testCases = []testCase{
	{name: "Sorted ascending", input: []int{-7, -3, 0, 1, 3, 5, 7, 9, 11, 23, 44}, want: []int{-7, -3, 0, 1, 3, 5, 7, 9, 11, 23, 44}},
	{name: "Sorted descending", input: []int{44, 23, 11, 9, 7, 5, 3, 1, 0, -3, -7}, want: []int{-7, -3, 0, 1, 3, 5, 7, 9, 11, 23, 44}},
	{name: "Not sorted slice", input: []int{-7, 7, 0, 1, 3, 5, -3, 9, 11, 23, 44}, want: []int{-7, -3, 0, 1, 3, 5, 7, 9, 11, 23, 44}},
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		c := make([]int, len(tc.input))
		copy(c, tc.input)
		bubbleSort(c)
		if !reflect.DeepEqual(c, tc.want) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, c)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		c := make([]int, len(tc.input))
		copy(c, tc.input)
		insertionSort(c)
		if !reflect.DeepEqual(c, tc.want) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, c)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			c := make([]int, len(tc.input))
			copy(c, tc.input)
			bubbleSort(c)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			c := make([]int, len(tc.input))
			copy(c, tc.input)
			insertionSort(c)
		}
	}
}

func ExampleBubbleSort() {
	input := []int{-7, 7, 0, 1, 3, 5, -3, 9, 11, 44, 23}
	fmt.Println("Not sorted slice: ", input)
	bubbleSort(input)
	fmt.Println("Sorted slice: ", input)
}
