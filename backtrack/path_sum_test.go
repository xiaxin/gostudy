package backtrack

import "testing"

func TestPathSum(t *testing.T) {
	var items = []struct {
		name   string
		input  []int
		target int
	}{
		{"113 input [5,4,8,11,null,13,4,7,2,null,null,5,1]", []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 9, 5, 1}, 22},
		{"113 input []", []int{}, 1},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			tree := BuildTreeBySliceInt(item.input)
			result := PathSum113(tree, item.target)

			for _, item := range result {
				t.Log(item)
			}
		})
	}
}
