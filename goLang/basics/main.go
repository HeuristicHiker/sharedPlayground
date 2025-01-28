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

func removeDup(nums []int) []int {
	// assume sorted

	finalLen := len(nums)

	if finalLen <= 1 {
		return nums
	}

	last, finder := 0, 0

	for last <= len(nums) {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				ans := nums[:last+1]
				fmt.Println("final nums: ", ans)
				return nums
			}
		}
		nums[last+1] = nums[finder]
		last++
	}

	return nums
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0 // need this to track the length of the array we care about

	for idx, num := range nums {
		// fmt.Println(num, val, " ", idx, j)
		if num != val {
			if idx != j {
				nums[idx], nums[j] = nums[j], nums[idx]
			}
			j++
		}
		fmt.Println("j ", j)
		fmt.Println(nums)
	}

	return j
}

func main() {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 2}, 2},
	}

	currentTest := tests[1]

	fmt.Println(removeElement(currentTest.nums, 2))

}
