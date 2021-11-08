package backtrack

import "testing"

/**
	例子
		1 2 3
		2 3 1
		3 1 2
**/
func TestSudoku(t *testing.T) {
	{
		nums := [][]int{
			{1, 2, 3},
			{2, 3, 1},
			{1, 2, 0},
		}

		t.Log(Sudoku(nums, 3))
	}

	{
		nums := [][]int{
			{3, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}

		t.Log(Sudoku(nums, 4))
	}
}
