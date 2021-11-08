package backtrack

import (
	"sort"
)

// 39. 组合总和
func CombinationSum39(candidates []int, target int) [][]int {
	var (
		result = make([][]int, 0)
		track  = make([]int, 0)

		fn func(candidates []int, target int, start int)
	)

	sort.Ints(candidates)

	fn = func(candidates []int, target int, start int) {
		for i := start; i < len(candidates); i++ {
			if target < candidates[i] {
				continue
			}
			if target == candidates[i] {
				// 不用copy 直接append 节省内存
				result = append(result, append([]int{candidates[i]}, track...))
				return
			}
			track = append(track, candidates[i])
			fn(candidates, target-candidates[i], i)
			track = track[0 : len(track)-1]
		}
	}

	fn(candidates, target, 0)

	return result
}

// 40. 组合总和 II
func CombinationSum40(candidates []int, target int) [][]int {
	var (
		result [][]int
		fn     func(candidates []int, target int, index int)
		track  []int
	)

	sort.Ints(candidates)

	fn = func(candidates []int, target int, start int) {

		for i := start; i < len(candidates); i++ {
			candidate := candidates[i]

			if target-candidate == 0 {
				result = append(result, append([]int{candidates[i]}, track...))
				return
			}

			if target-candidate < 0 {
				continue
			}

			// 过滤重复数字
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			track = append(track, candidate)
			fn(candidates, target-candidate, i+1)
			track = track[0 : len(track)-1]

		}
	}
	fn(candidates, target, 0)

	return result
}

// TODO 216. 组合总和 III（困难）
