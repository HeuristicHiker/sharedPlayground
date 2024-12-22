package main

import (
	"testing"
)

func TestRichestCustomer(t *testing.T) {
	testCase := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	result := richestCustomer(testCase)
	soln := 1

	if result != soln {
		t.Errorf("Expected %v but got %v", soln, result)
	}

}
