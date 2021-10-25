package array

import "testing"

func TestMergeSort(t *testing.T) {
	var nums []int
	nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	t.Log(mergeSort(nums))

	nums = []int{}
	t.Log(mergeSort(nums))

	nums = []int{3, 3}
	t.Log(mergeSort(nums))

	nums = []int{2, 1}
	t.Log(mergeSort(nums))
}
