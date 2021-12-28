package two_pointers

import "testing"

func TestThreeSum(t *testing.T) {
	var items = []struct {
		name  string
		input []int
	}{
		{"15 result [[-1,-1,2],[-1,0,1]", []int{-1, 0, 1, 2, -1, -4}},
		{"15 result []", []int{}},
		{"15 result []", []int{0}},
		{"15 result [-1,0,1,2,-1,-4,-2,-3,3,0,4]", []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			{
				result := ThreeSum(item.input)
				t.Log(result)
			}

			{
				result := ThreeSumTP(item.input)
				t.Log(result)
			}
		})
	}
}
