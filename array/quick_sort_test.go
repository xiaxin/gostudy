package array

import "testing"

func TestQuickSort(t *testing.T) {
	var nums []int
	nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	t.Log(quickSort(nums, 0, len(nums)-1))

	// nums = []int{}
	// t.Log(quickSort(nums, 0, len(nums)-1))

	// nums = []int{3, 3}
	// t.Log(quickSort(nums, 0, len(nums)-1))

	// nums = []int{2, 1}
	// t.Log(quickSort(nums, 0, len(nums)-1))
}
