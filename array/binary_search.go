package array

import "fmt"

func binarySearch(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// left
		if nums[mid] > target {
			right = mid - 1
		}

		// right
		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] == target {
			return mid
		}
	}

	return -1
}

func binarySearchLeftBoard(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		}

		if target < nums[mid] {
			right = mid - 1
		}

		if nums[mid] == target {
			right = mid - 1
		}
	}

	fmt.Println(left, right)

	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func binarySearchRightBoard(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		}

		if target < nums[mid] {
			right = mid - 1
		}

		if nums[mid] == target {
			left = mid + 1
		}
	}

	if right >= 0 && nums[right] == target {
		return right
	}

	return -1
}
