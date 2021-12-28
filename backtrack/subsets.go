package backtrack

// 78. 子集
func Subsets78(nums []int) [][]int {
	var (
		result = [][]int{nil}
		track  []int
		fn     func(index int)
	)

	fn = func(index int) {

		result = append(result, append([]int{}, track...))

		if index == len(nums)-1 {
			return
		}

		for i := index; i < len(nums); i++ {

			if i <= index {
				continue
			}

			track = append(track, nums[i])
			fn(i)
			track = track[0 : len(track)-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		track = append(track, nums[i])
		fn(i)
		track = track[0 : len(track)-1]
	}

	return result
}

// TODO 90. 子集 II
