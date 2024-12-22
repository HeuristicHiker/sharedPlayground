package main

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	testCase := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result := diagonalSum(testCase)
	soln := 25

	if result != soln {
		t.Errorf("Expected %v but got %v", soln, result)
	}

}
