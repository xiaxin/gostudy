package backtrack

import (
	"gostudy/common"
)

func Parentheses(n int) []string {
	var result []string
	stack := common.NewStack()

	parentheses(n, n, stack, &result)

	return result
}

func parentheses(left, right int, stack *common.Stack, result *[]string) {

	if right < left {
		return
	}
	if left < 0 || right < 0 {
		return
	}

	if left == 0 && right == 0 {
		tmp := stack.ValString()
		*result = append(*result, tmp)
		return
	}

	stack.PushNode(0, "(")

	parentheses(left-1, right, stack, result)
	stack.Pop()

	stack.PushNode(1, ")")
	parentheses(left, right-1, stack, result)
	stack.Pop()
}
