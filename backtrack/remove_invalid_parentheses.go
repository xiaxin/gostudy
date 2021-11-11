package backtrack

// TODO 301. 删除无效的括号 第一次开发思路错误
func RemoveInvalidParentheses(s string) []string {
	var (
		result []string
		track  []byte

		visited = make(map[int]bool)
		hash    = make(map[string]bool)

		input                = []byte(s)
		oldLeft, left, right = 0, 0, 0
	)

	// if len(input) > 0 && input[0] != ')' {
	// 	track = append(track, input[0])

	// 	if input[0] == '(' {
	// 		left++
	// 	}

	// 	oldLeft = left
	// }

	if len(input) == 0 {
		result = append(result, "")
		return result
	}

	if len(input) == 2 && input[0] == '(' && input[1] == ')' {
		result = append(result, "()")
		return result
	}

	for i := 0; i < len(input); i++ {

		if input[i] == '(' || input[i] == ')' {
			visited[i] = true
		}

		for j := 0; j < len(input); j++ {
			var (
				char  = input[j]
				index = j
			)
			if _, ok := visited[index]; ok {
				continue
			}

			if index == 0 && char == ')' {
				continue
			}

			if index == len(input)-1 && char == '(' {
				continue
			}

			if char == '(' {
				left++
			}

			if char == ')' {
				right++
			}
			track = append(track, char)
		}
		if input[i] == '(' || input[i] == ')' {
			delete(visited, i)
		}

		if left == right && len(track) >= 0 {

			if len(track) > 0 && track[len(track)-1] == '(' {
				continue
			}

			var key string
			if len(track) > 0 && track[len(track)-1] != '(' {
				key = string(track)
			}

			if len(track) == 0 {
				key = ""
			}

			if _, ok := hash[key]; !ok {
				hash[key] = true
				result = append(result, key)
			}
		}
		track = track[0:0]
		left, right = oldLeft, 0
	}

	if len(result) == 0 {

		result = append(result, "")
	}

	return result
}
