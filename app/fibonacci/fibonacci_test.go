package fibonacci

import "testing"

type testFibonacci struct {
	input uint32
	want  uint32
}

var result uint32
var testCases = []testFibonacci{
	{
		input: 0,
		want:  0,
	},
	{
		input: 1,
		want:  1,
	},
	{
		input: 2,
		want:  1,
	},
	{
		input: 3,
		want:  2,
	},
	{
		input: 4,
		want:  3,
	},
	{
		input: 8,
		want:  21,
	},
	{
		input: 20,
		want:  6765,
	},
}

func TestF(t *testing.T) {
	for _, tc := range testCases {
		got := F(tc.input)
		if got != tc.want {
			t.Fatalf("input: %v: expected: %v, got: %v", tc.input, tc.want, got)
		}
	}
}

func TestFibonacci(t *testing.T) {
	for _, tc := range testCases {
		got := fibonacci(tc.input)
		if got != tc.want {
			t.Fatalf("input: %v: expected: %v, got: %v", tc.input, tc.want, got)
		}
	}
}

func BenchmarkF(b *testing.B) {
	var res uint32
	for i := 0; i < b.N; i++ {
		res = F(20)
	}
	result = res
}

func BenchmarkFibonacci(b *testing.B) {
	var res uint32
	for i := 0; i < b.N; i++ {
		res = fibonacci(20)
	}
	result = res
}
