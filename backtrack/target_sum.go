package backtrack

import (
	"gostudy/common"
)

/**
	pos 指向 nums 的位置

	pos == len(nums) 退出  如果 track 里的值是 target 加入 result

	track 里的 size() == 5

	循环 i := range nums
		i 分两种 +i 和 -i 都对应一套 pop  push
**/

func TargetSum(nums []int, target int) []string {
	var result []string
	stack := common.NewStack()

	targetSum(nums, 0, 0, target, stack, &result)

	return result
}

func targetSum(nums []int, pos int, sum int, target int, stack *common.Stack, result *[]string) {

	// fmt.Println(sum, stack)
	if pos == len(nums) {
		if sum == target {
			tmp := stack.String()
			*result = append(*result, tmp)
			return
		}
		return
	}

	stack.PushNode(1, nums[pos])
	targetSum(nums, pos+1, sum+nums[pos], target, stack, result)
	stack.Pop()

	stack.PushNode(-1, -nums[pos])
	targetSum(nums, pos+1, sum-nums[pos], target, stack, result)
	stack.Pop()
}
