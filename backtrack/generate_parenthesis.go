package backtrack

import (
	"strings"
)

// 22. 括号生成
func GenerateParentheses(n int) []string {
	var (
		result []string
		track  []string
		fn     func(left int, right int)
	)

	// left/right 左右括号的次数
	fn = func(left int, right int) {
		if right < left {
			return
		}

		if left < 0 || right < 0 {
			return
		}

		if left == 0 && right == 0 {
			//  添加结果
			result = append(result, strings.Join(track, ""))
			return
		}

		track = append(track, "(")
		fn(left-1, right)
		track = track[0 : len(track)-1]

		track = append(track, ")")
		fn(left, right-1)
		track = track[0 : len(track)-1]
	}

	fn(n, n)

	return result
}
