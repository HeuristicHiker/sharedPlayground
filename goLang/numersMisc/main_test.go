package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	in  []int
	out []int
}

func TestPlusOne(t *testing.T) {
	testCases := []testCase{
		{
			in:  []int{1, 2, 3},
			out: []int{1, 2, 4},
		},
		{
			in:  []int{4, 3, 2, 1},
			out: []int{4, 3, 2, 2},
		},
		{
			in:  []int{9},
			out: []int{1, 0},
		},
	}

	for i, testCase := range testCases {
		fmt.Println("--- Test ", i, " ---")
		result := plusOne(testCase.in)
		fmt.Println(result)
		if !reflect.DeepEqual(result, testCase.out) {
			t.Errorf("Test case: %d Expected: %v but got: %v", i, testCase.out, result)
		}
	}
}
