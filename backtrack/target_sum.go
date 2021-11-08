package backtrack

import (
	"gostudy/common"
)

// TODO 494. 目标和
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
