package two_pointers

import "testing"

func TestTwoAreas(t *testing.T) {
	var items = []struct {
		name  string
		input []int
	}{
		{"11 input 48", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		{"11 input 1", []int{1, 1}},
		{"11 input 16", []int{4, 3, 2, 1, 4}},
		{"11 input 2", []int{1, 2, 1}},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			t.Log(MaxArea(item.input))
		})
	}

}
