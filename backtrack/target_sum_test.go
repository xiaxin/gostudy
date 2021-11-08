package backtrack

import "testing"

func TestTargetSum(t *testing.T) {
	var items = []struct {
		name   string
		input  []int
		target int
		result int
	}{
		{"494 input [1, 1, 1, 1, 1]", []int{1, 1, 1, 1, 1}, 3, 5},
		{"494 input [1, 1, 1, 1, 2]", []int{1, 1, 1, 1, 2}, 2, 7},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := TargetSum(item.input, item.target)

			t.Log(result)
		})
	}
}
