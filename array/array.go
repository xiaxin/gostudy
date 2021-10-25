package array

type IntArray struct {
	val []int
}

func NewIntArray(val []int) *IntArray {
	// TODO 增加快速排序
	return &IntArray{val}
}

func (i IntArray) IsExist(num int) bool {
	return i.search(num, 0, len(i.val)) > 0
}

func (i IntArray) Search(num int) int {
	return i.search(num, 0, len(i.val)-1)
}

func (i IntArray) search(num int, left, right int) int {
	// FUNC A

	// if left > right {
	// 	return -1
	// }
	// mid := left + (right-left)/2
	// value := i.val[mid]

	// if num == value {
	// 	return mid
	// } else if num < value {
	// 	return i.search(num, left, mid-1)
	// } else if num > value {
	// 	return i.search(num, mid+1, right)
	// }

	// FUNC B

	for left <= right {
		mid := left + (right-left)/2
		value := i.val[mid]
		if num < value {
			right = mid - 1
		} else if num > value {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1

}
