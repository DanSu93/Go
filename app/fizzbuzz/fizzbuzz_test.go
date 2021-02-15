package fizzbuzz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  string
	}{
		{name: "Fizz", input: 6, want: "Fizz"},
		{name: "Fizz", input: 9, want: "Fizz"},
		{name: "Buzz", input: 5, want: "Buzz"},
		{name: "Buzz", input: 10, want: "Buzz"},
		{name: "FizzBuzz", input: 15, want: "FizzBuzz"},
		{name: "FizzBuzz", input: 45, want: "FizzBuzz"},
		{name: "Number", input: 7, want: "7"},
		{name: "Number", input: 8, want: "8"},
	}

	for _, tc := range tests {
		got := Fizzbuzz(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func ExampleFizzbuzz() {
	fmt.Printf(Fizzbuzz(20))
	fmt.Println()
	fmt.Printf(Fizzbuzz(45))
}
