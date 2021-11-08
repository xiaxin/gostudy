package backtrack

import (
	"fmt"
	"testing"
)

func TestCombinationSum39(t *testing.T) {
	var datas = []struct {
		input  []int
		target int
	}{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
		{[]int{2}, 1},
		{[]int{1}, 1},
		{[]int{1}, 2},
		{[]int{2, 7, 6, 3, 5, 1}, 9},
	}

	for i, data := range datas {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {

			result := CombinationSum39(data.input, data.target)

			for _, v := range result {
				t.Log(data.input, v)
			}
		})
	}

}

func TestCombinationSum40(t *testing.T) {
	var datas = []struct {
		input  []int
		target int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
		{[]int{2, 5, 2, 1, 2}, 5},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 27},
	}

	for i, data := range datas {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {

			result := CombinationSum40(data.input, data.target)

			for _, v := range result {
				t.Log(data.input, v)
			}
		})
	}
}
