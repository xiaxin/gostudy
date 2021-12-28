package backtrack

import (
	"sort"
)

// TODO 473. 火柴拼正方形

func Makesquare(matchsticks []int) bool {
	var (
		sum, svg int
		fn       func(index int, square [4]int) bool

		square   = [4]int{}
		matchLen = len(matchsticks)
	)

	for _, match := range matchsticks {
		sum += match
	}

	if sum%4 != 0 {
		return false
	}

	svg = sum / 4

	sort.Ints(matchsticks)

	fn = func(index int, square [4]int) bool {
		// 退出条件
		if index == -1 {
			if square[0] == svg &&
				square[1] == svg &&
				square[2] == svg &&
				square[3] == svg {
				return true
			}
		}

		for i := 0; i < len(square); i++ {
			// 把 火柴 放在 i 边上 是否大于 svg
			if square[i]+matchsticks[index] > svg ||
				// TODO
				// size[i] == size[i - 1]即上一个分支的值和当前分支的一样，上一个分支没有成功，
				// 说明这个分支也不会成功，直接跳过即可。
				(i > 0 && square[i] == square[i-1]) || (index == matchLen-1 && i != 0) {
				continue
			}

			square[i] += matchsticks[index]
			if fn(index-1, square) {
				return true
			}
			square[i] -= matchsticks[index]
		}
		return false
	}

	return fn(matchLen-1, square)
}
