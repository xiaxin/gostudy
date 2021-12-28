package backtrack

import (
	"fmt"
)

// 1980. 找出不同的二进制字符串
// TODO 没使用回溯方法
func FindDifferentBinaryString(nums []string) string {

	var (
		max    int
		result string

		visited = make(map[byte]bool)
	)

	if len(nums) == 0 {
		return result
	}

	for _, strs := range nums {
		var num byte
		for i := len(strs) - 1; i >= 0; i-- {
			num += (strs[i] - 48) << (len(strs) - i - 1)
		}

		visited[num] = true
	}

	numLen := len(nums[0])

	for i := 0; i < numLen; i++ {
		max += (1 << i)
	}

	for i := 0; i <= max; i++ {
		if _, ok := visited[byte(i)]; !ok {
			result = intToString(numLen, i)
			break
		}
	}

	return result

}

func findDifferentBinaryString(nums []string) string {
	var (
		max    int
		result string

		visited = make(map[byte]bool)
	)

	if len(nums) == 0 {
		return result
	}

	for _, strs := range nums {
		var num byte
		for i := len(strs) - 1; i >= 0; i-- {
			num += (strs[i] - 48) << (len(strs) - i - 1)
		}

		visited[num] = true
	}

	numLen := len(nums[0])

	for i := 0; i < numLen; i++ {
		max += (1 << i)
	}

	for i := 0; i <= max; i++ {
		if _, ok := visited[byte(i)]; !ok {
			result = intToString(numLen, i)
			break
		}
	}

	return result
}

func intToString(x int, n int) string {
	tmp := fmt.Sprintf("%%0%db", x)
	return fmt.Sprintf(tmp, n)
}

func findDifferentBinaryStringByLeetcode(nums []string) string {
	n := len(nums)
	if n == 0 {
		return ""
	}

	var (
		result string
		used   = make(map[string]bool)
		fn     func(index int, base string)
	)

	for _, num := range nums {
		used[num] = true
	}

	fn = func(index int, base string) {

		if index == n {
			if _, ok := used[base]; !ok {
				result = base
			}
			return
		}

		fn(index+1, base+"0")
		fn(index+1, base+"1")
	}

	fn(0, "")

	return result
}
