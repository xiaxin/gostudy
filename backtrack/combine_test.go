package backtrack

import "testing"

func TestCombine(t *testing.T) {
	var items = []struct {
		name string
		n, k int
	}{
		{"77 input n=4 k=2", 4, 2},
		{"77 input n=1 k=1", 1, 1},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := Combine(item.n, item.k)
			for _, v := range result {
				t.Log(v)
			}
		})
	}
}
