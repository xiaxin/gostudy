package two_pointers

// 11. 盛最多水的容器
func MaxArea(nums []int) int {
	var (
		i, height, maxNum int
	)
	j := len(nums) - 1

	for i < j {

		if nums[i] < nums[j] {
			height = nums[i]
			i++
		} else {
			height = nums[j]
			j--
		}

		if tmp := height * (j - i + 1); tmp > maxNum {
			maxNum = tmp
		}
	}

	return maxNum
}
