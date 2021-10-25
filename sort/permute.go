package sort

// TODO package 没放对

/**
46. 全排列（https://leetcode-cn.com/problems/permutations/）

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
	输入：nums = [0,1]
	输出：[[0,1],[1,0]]

示例 3：
	输入：nums = [1]
	输出：[[1]]


提示：
	1 <= nums.length <= 6
	-10 <= nums[i] <= 10
	nums 中的所有整数 互不相同
	**/

// 全排列 TODO 待优化
func permute(nums []int) [][]int {
	var (
		trackMap map[int]bool = make(map[int]bool)
		result   [][]int
		track    []int
		fn       func(nums, track []int, trackMap map[int]bool)
	)
	fn = func(nums, track []int, trackMap map[int]bool) {
		if len(nums) == len(trackMap) {
			res := make([]int, len(nums))
			copy(res, track)
			result = append(result, res)
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := trackMap[i]; ok {
				continue
			}

			trackMap[i] = true
			track = append(track, nums[i])

			fn(nums, track, trackMap)

			delete(trackMap, i)
			track = track[0 : len(track)-1]
		}

	}
	fn(nums, track, trackMap)
	return result
}
