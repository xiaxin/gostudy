package backtrack

// TODO 37. 解数独
func Sudoku(nums [][]int, wh int) (bool, [][]int) {
	return sudoku(nums, wh, 0, 0)
}

func sudoku(nums [][]int, wh, i, j int) (bool, [][]int) {
	if j == wh {
		return sudoku(nums, wh, i+1, 0)
	}

	if i == wh {
		return true, nums
	}

	if nums[i][j] > 0 {
		return sudoku(nums, wh, i, j+1)
	}

	for c := 1; c <= wh; c++ {
		if !sudoKuIsValid(nums, wh, i, j, c) {
			continue
		}
		nums[i][j] = c
		if ok, _ := sudoku(nums, wh, i, j+1); ok {
			return true, nums
		}
		nums[i][j] = 0
	}

	return false, nil
}

func sudoKuIsValid(nums [][]int, wh, i, j, n int) bool {
	for x := 0; x < wh; x++ {
		if nums[i][x] == n {
			return false
		}
		if nums[x][j] == n {
			return false
		}
	}
	return true
}
