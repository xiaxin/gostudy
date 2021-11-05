package backtrack

func Backtrack(nums []int) [][]int {
	var (
		contains = make(map[int]bool)
		result   [][]int
		cache    []int
		fn       func([]int, []int, map[int]bool)
	)

	fn = func(nums []int, cache []int, contains map[int]bool) {

		if len(nums) == len(contains) {
			var tmp = make([]int, len(contains))

			copy(tmp, cache)

			result = append(result, tmp)
			return
		}

		for _, n := range nums {
			if _, ok := contains[n]; ok {
				continue
			}
			contains[n] = true
			cache = append(cache, n)

			fn(nums, cache, contains)

			delete(contains, n)

			cache = cache[0 : len(cache)-1]
		}
	}

	fn(nums, cache, contains)

	return result
}
