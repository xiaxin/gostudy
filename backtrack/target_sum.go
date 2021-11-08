package backtrack

// 494. 目标和 TODO 时间太久
func TargetSum(nums []int, target int) int {
	var (
		result int
		track  []int
		fn     func(nums []int, pos int, sum int, target int)
	)

	fn = func(nums []int, pos int, sum int, target int) {

		if pos == len(nums) {
			if sum == target {
				result += 1
				return
			}
			return
		}

		track = append(track, nums[pos])
		fn(nums, pos+1, sum+nums[pos], target)
		track = track[0 : len(track)-1]

		track = append(track, -nums[pos])
		fn(nums, pos+1, sum-nums[pos], target)
		track = track[0 : len(track)-1]
	}

	fn(nums, 0, 0, target)

	return result
}
