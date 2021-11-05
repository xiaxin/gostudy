package backtrack

import "testing"

func TestTargetSum(t *testing.T) {
	nums := []int{1, 1, 1, 1, 2}

	for _, v := range TargetSum(nums, 2) {
		t.Log(v)
	}
}
