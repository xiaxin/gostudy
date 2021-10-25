package array

import "testing"

func TestIntArray(t *testing.T) {
	array := NewIntArray([]int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	t.Log(array.Search(5))
	t.Log(array.Search(0))
	t.Log(array.Search(-1))
	t.Log(array.Search(10))
	t.Log(array.Search(100))
}

// Slice 索引测试
func TestSlice(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5}

	t.Log(nums[3:4]) // 索引范围 [1, 2) = [1, 1]

	t.Log(nums[:6]) // 索引范围 [0, 6) = [0, 5]
	t.Log(nums[1:])
}
