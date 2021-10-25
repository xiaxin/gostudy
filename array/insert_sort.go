package array

func insertSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		temp := nums[i]

		j := i - 1

		for j >= 0 && temp < nums[j] {
			nums[j+1] = nums[j]
			j--
		}

		if j+1 != i {
			nums[j+1] = temp
		}
	}
	return nums
}
