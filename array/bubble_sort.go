package array

// 数组冒泡排序
func bubbleSort(nums []int) []int {
	// 1个和0个元素不需要排序
	if len(nums) <= 1 {
		return nums
	}
	// 从后网前比较，将最大放置后面

	// for i := len(nums) - 1; i >= 0; i-- {
	// 	for j := i - 1; j >= 0; j-- {
	// 		if nums[i] < nums[j] {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 		}
	// 	}
	// }

	// 从前往后比较，将最小放置最前
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
