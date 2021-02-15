package primenum

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  bool
	}{
		{name: "prime number", input: 7, want: true},
		{name: "prime number", input: 11, want: true},
		{name: "prime number", input: 13, want: true},
		{name: "prime number", input: 17, want: true},
		{name: "prime number", input: 31, want: true},
		{name: "not a prime number", input: 1, want: false},
		{name: "not a prime number", input: 8, want: false},
		{name: "not a prime number", input: 14, want: false},
		{name: "not a prime number", input: 16, want: false},
		{name: "not a prime number", input: 100, want: false},
	}

	for _, tc := range tests {
		got := isPrime(tc.input)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got %v,", tc.name, tc.want, got)
		}
	}
}

func TestFindPrimeNumbers(t *testing.T) {
	tests := []struct {
		input int64
		want  []int64
	}{
		{input: 5, want: []int64{2, 3, 5}},
		{input: 24, want: []int64{2, 3, 5, 7, 11, 13, 17, 19, 23}},
	}

	for _, tc := range tests {
		got := findPrimeNumbers(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
