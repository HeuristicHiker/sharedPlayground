package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for idx, num := range nums {
		want := target - num

		if val, ok := seen[want]; ok {
			return []int{val, idx}
		} else {
			seen[num] = idx
		}
	}

	return nil
}

func main() {
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

	currentTest := tests[0]

	fmt.Println(twoSum(currentTest.nums, currentTest.target))

}
