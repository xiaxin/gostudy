package dp

import "testing"

func TestBag01(t *testing.T) {
	{
		// nums := [][]int{{1, 1500}, {4, 3000}, {3, 2000}}

		// res := bag01(nums, 4)

		// t.Log(res)
	}

	{
		// nums := [][]int{{1, 1500}, {4, 3000}, {3, 2000}}
		nums := [][]int{{4, 3000}, {3, 2000}, {1, 1500}}

		res := bag01(nums, 4)

		t.Log(res)
	}
}
