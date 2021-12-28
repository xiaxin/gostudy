package backtrack

// TODO 79. 单词搜索
func WordSearch79(board [][]byte, word string) bool {
	var (
		r, c    int
		row     []byte
		col     byte
		fn      func(row int, col int, index int) bool
		visited = make(map[int]map[int]bool)
		wordLen = len(word)
		rowLen  = len(board)
		colLen  = len(board[0])
	)

	fn = func(row, col, index int) bool {
		if wordLen == index {
			return true
		}
		if index > wordLen ||
			row < 0 || col < 0 ||
			row >= rowLen || col >= colLen ||
			board[row][col] != word[index] {
			return false
		}

		if visited[row] == nil {
			visited[row] = make(map[int]bool)
		}

		if _, ok := visited[row][col]; ok {
			return false
		}

		for _, item := range [][]int{{-1, -0}, {+1, 0}, {0, -1}, {0, +1}} {
			visited[row][col] = true
			if fn(row+item[0], col+item[1], index+1) {
				return true
			}
			delete(visited[row], col)
		}

		return false

	}

	for r, row = range board {
		for c, col = range row {
			if col == word[0] {
				if fn(r, c, 0) {
					return true
				}
			}
		}
	}

	return false
}

// TODO 212. 单词搜索 II
