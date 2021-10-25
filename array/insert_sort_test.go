package array

import "testing"

func TestInsertSort(t *testing.T) {
	var nums []int
	nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	t.Log(insertSort(nums))

	nums = []int{}
	t.Log(insertSort(nums))

	nums = []int{3, 3}
	t.Log(insertSort(nums))

	nums = []int{2, 1}
	t.Log(insertSort(nums))
}
