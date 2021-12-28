package backtrack

import "testing"

func TestMakeSquare(t *testing.T) {
	var items = []struct {
		name   string
		input  []int
		result bool
	}{
		{"473 result true", []int{1, 1, 2, 2, 2}, true},
		{"473 result false", []int{3, 3, 3, 3, 4}, false},
		{"473 result true", []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}, true},
		{"473 result ????", []int{12, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48, 52, 56, 60}, false},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := Makesquare(item.input)
			t.Log(result)
		})
	}
}
