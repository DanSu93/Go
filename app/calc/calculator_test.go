// Package calc provides basic mathematical operations on numbers.
// Support operations: addition, subtraction, multiplication, division, raising to a power, taking a percentage.
// This package does not guarantee bit-identical results across architectures.
package calc

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type testCalculate struct {
	x         float64
	y         float64
	op        rune
	want      float64
	expectErr error
}

var testCases = []testCalculate{
	{
		x:         32.35,
		y:         -17.2,
		op:        '*',
		want:      -556.42,
		expectErr: nil,
	},
	{
		x:         32.35,
		y:         -17.2,
		op:        '/',
		want:      -1.8808139534883723,
		expectErr: nil,
	},
	{
		x:         -2,
		y:         4,
		op:        '^',
		want:      16,
		expectErr: nil,
	},
	{
		x:         12,
		y:         0,
		op:        '/',
		want:      0,
		expectErr: errors.New("cannot be divided by 0"),
	},
	{
		x:         14,
		y:         8,
		op:        '&',
		want:      0,
		expectErr: errors.New("operation specified incorrectly"),
	},
}

// Testing the Calculate Function
func TestCalculate(t *testing.T) {
	for _, tc := range testCases {
		got, err := Calculate(tc.x, tc.y, tc.op)
		if !reflect.DeepEqual(tc.want, got) || !reflect.DeepEqual(tc.expectErr, err) {
			t.Fatalf("%s: expected: %v, got: %v", string(tc.op), tc.want, got)
		}
	}
}

// Example using the Calculate function
func ExampleCalculate() {
	res, _ := Calculate(12, 8, '+')
	fmt.Println(res)
}
