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

func removeDup(nums []int) int {
	// assume sorted
	for idx := 1; idx <= len(nums)-1; idx++ {
		fmt.Println(idx)
	}
	return 0
}

func main() {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 1, 2}, 2},
	}

	currentTest := tests[0]

	fmt.Println(removeDup(currentTest.nums), 2)

}
