package main

import (
	"fmt"
	"math"
	"sort"
)

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
		if num != val {
			if idx != j {
				nums[idx], nums[j] = nums[j], nums[idx]
			}
			j++
		}
	}

	return j
}

// clamping
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n <= 2 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		redundantNumber := i > 0 && nums[i] == nums[i-1]

		if redundantNumber {
			continue
		}

		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			fmt.Println(nums[i], nums[j], nums[k])
			if abs(sum-target) < diff {
				res, diff = sum, abs(sum-target)
			}
			if sum == target {
				return res
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return res
}

func threeSumClosest2(nums []int, target int) int {
	n, res, dist := len(nums), 0, math.MaxInt32

	if n < 3 {
		return res
	}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			newDist := abs(sum - target)

			// can't do better than perfect
			if sum == target {
				return sum
			}
			// if distance is shorter, set the new result and difference
			if newDist < dist {
				res, dist = sum, abs(sum-target)
			} else if sum > target {
				k--
			} else {
				j++
			}

		}
	}

	return res

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 2}, 2},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 7, 8, 10}, 3},
		{[]int{-1, 2, 1, -4}, 1},
	}

	currentTest := tests[3]

	fmt.Println(threeSumClosest2(currentTest.nums, 2))

}
