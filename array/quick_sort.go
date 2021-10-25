package array

/**
  TODO
**/
func quickSort(nums []int, left, right int) []int {

	if left >= right {
		return nums
	}
	key := nums[right]

	i, j := left, right

	for {
		if nums[i] < key {
			i++
		}

		if key < nums[j] {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	quickSort(nums, left, i-1)
	quickSort(nums, j+1, right)

	return nums
}
