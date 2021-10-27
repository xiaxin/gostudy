package array

func selectSort(nums []int) []int {
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		min := i
		for j := i + 1; j < numsLen; j++ {
			if nums[i] > nums[j] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
