package array

import "testing"

func TestBubbleSort(t *testing.T) {
	{
		nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

		t.Log(bubbleSort(nums))
	}

	{
		nums := []int{9, 8, 7, 6, 5, 10, 4, 3, 2, 1, 0}

		t.Log(bubbleSort(nums))
	}
}
