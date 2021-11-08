package backtrack

import "testing"

func TestPermute46(t *testing.T) {
	var items = []struct {
		name  string
		input []int
	}{
		{"46 Input: [1, 2, 3]", []int{1, 2, 3}},
		{"46 Input: [0, 1]", []int{0, 1}},
		{"46 Input: [1]", []int{1}},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := Permute46(item.input)

			for _, res := range result {
				t.Log(res)
			}
		})
	}
}

func TestPermute47(t *testing.T) {
	var items = []struct {
		name  string
		input []int
	}{
		{"47 Input: [1, 1, 2]", []int{1, 1, 2}},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := PermuteUnique(item.input)

			for _, res := range result {
				t.Log(res)
			}
		})
	}
}
