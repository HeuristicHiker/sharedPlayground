package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1000000, 2000000, 3000000}, 5000000, []int{1, 2}},
		{[]int{1, 9, 5, 7}, 8, []int{0, 3}},
		{[]int{1, 2}, 3, []int{0, 1}},
		{[]int{-5, -3, -10, -1}, -8, []int{0, 1}},
		{[]int{-10, -20, -30, -40}, -50, []int{1, 2}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums=%v and target=%d, expected %v but got %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
