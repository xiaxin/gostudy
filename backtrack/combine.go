package backtrack

// 77. 组合（中等）
func Combine(n int, k int) [][]int {
	var (
		result [][]int
		nums   []int
		track  []int
		fn     func(nums []int, start int, track []int)
	)
	// 构造树结构数据（一维数组）
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	fn = func(nums []int, start int, track []int) {
		if len(track) == k {
			result = append(result, append([]int{}, track...))
			return
		}

		for i := start; i < len(nums); i++ {
			// 剪枝：track 长度加上区间 [nums, n] 的长度小于 k，不可能构造出长度为 k 的 track
			if len(track)+(n-nums[i]+1) < k {
				return
			}
			track = append(track, nums[i])
			fn(nums, i+1, track)
			track = track[0 : len(track)-1]
		}
	}

	fn(nums, 0, track)

	return result
}
