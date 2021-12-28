package two_pointers

import "testing"

func TestThreeSumClosest(t *testing.T) {
	var items = []struct {
		name   string
		input  []int
		target int
	}{
		// {"16 input [-1,2,1,-4]", []int{-1, 2, 1, -4}, 2},
		// {"16 input [0,0,0]", []int{0, 0, 0}, 0},
		// {"16 input [1,1,1,1]", []int{1, 1, 1, 1}, 0},
		// {"16 input [1,1,1,0]", []int{1, 1, 1, 0}, -100},
		{"16 input [1,1,-1,-1,3]", []int{1, 1, -1, -1, 3}, -1},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := ThreeSumClosest(item.input, item.target)
			t.Log(result)
		})
	}
}
