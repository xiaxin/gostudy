package array

/**
	1. 将数据拆解拆成类二叉树结构
	2. 递归拆解结束以后，自底向上进行合并（叶子节点为一个节点）
**/
func mergeSort(s []int) []int {
	return splitNums(s)
}

func splitNums(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	mid := len / 2

	leftNums := splitNums(nums[0:mid])
	rightNums := splitNums(nums[mid:len])

	return mergeNums(leftNums, rightNums)
}

func mergeNums(leftNums, rightNums []int) []int {
	var result []int
	leftLen, rightLen := len(leftNums), len(rightNums)
	leftIndex, rightIndex := 0, 0

	for leftIndex < leftLen && rightIndex < rightLen {
		leftVal := leftNums[leftIndex]
		rightVal := rightNums[rightIndex]

		if leftVal < rightVal {
			result = append(result, leftNums[leftIndex])
			leftIndex++
		} else {
			result = append(result, rightNums[rightIndex])
			rightIndex++
		}
	}

	if leftIndex < leftLen {
		result = append(result, leftNums[leftIndex:]...)
	}

	if rightIndex < rightLen {
		result = append(result, rightNums[rightIndex:]...)
	}

	return result
}
