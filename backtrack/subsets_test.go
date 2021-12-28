package backtrack

import "testing"

func TestSubsets(t *testing.T) {
	var items = []struct {
		name  string
		input []int
	}{
		{"78 input [1,2,3]", []int{1, 2, 3}},
		{"78 input [0]", []int{0}},
		{"78 input [1,2,2]", []int{1, 2, 2}},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			// 78
			{
				result := Subsets78(item.input)
				t.Log(result)
			}
		})
	}
}
