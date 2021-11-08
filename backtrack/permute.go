package backtrack

import (
	"sort"
)

// 46. 全排列
func Permute46(nums []int) [][]int {
	var (
		result [][]int
		track  []int
		fn     func(*[]int, *[]int, *map[int]bool)

		visited = make(map[int]bool)
	)

	fn = func(nums *[]int, track *[]int, visited *map[int]bool) {
		if len(*nums) == len(*track) {
			result = append(result, append([]int{}, *track...))
			return
		}
		for _, n := range *nums {
			if _, ok := (*visited)[n]; ok {
				continue
			}

			(*visited)[n] = true

			*track = append(*track, n)

			fn(nums, track, visited)

			*track = (*track)[0 : len(*track)-1]

			delete(*visited, n)
		}
	}
	fn(&nums, &track, &visited)
	return result
}

// 47. 全排列 II
func PermuteUnique(nums []int) [][]int {

	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return nil
	}

	var (
		result [][]int
		track  []int
		fn     func([]int)

		visited = make(map[int]bool)
	)

	fn = func(nums []int) {
		if len(nums) == len(track) {
			result = append(result, append([]int{}, track...))
			return
		}
		for i, n := range nums {

			// TODO !visited[i-1] 作用
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			if _, ok := visited[i]; ok {
				continue
			}

			visited[i] = true

			track = append(track, n)

			fn(nums)

			track = (track)[0 : len(track)-1]

			delete(visited, i)
		}
	}
	fn(nums)
	return result
}
