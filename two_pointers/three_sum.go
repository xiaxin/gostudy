package two_pointers

import (
	"sort"
)

// 15. 三数之和
func ThreeSum(nums []int) [][]int {
	var (
		n      = len(nums)
		result [][]int
	)

	sort.Ints(nums)

	for a := 0; a < n; a++ {
		if nums[a] > 0 {
			break
		}
		c := n - 1
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for b < c && nums[a]+nums[b]+nums[c] > 0 {
				c--
			}

			if b == c {
				break
			}

			if nums[a]+nums[b]+nums[c] == 0 {
				result = append(result, []int{nums[a], nums[b], nums[c]})
			}

		}
	}

	return result
}

func ThreeSumTP(nums []int) [][]int {

	var (
		n      = len(nums)
		result [][]int
	)

	sort.Ints(nums)

	for a := 0; a < n; a++ {
		if nums[a] > 0 {
			break
		}
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		target := -1 * nums[a]

		for left, right := a+1, n-1; left < right; {
			if left > a+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[a], nums[left], nums[right]})
				left++
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
