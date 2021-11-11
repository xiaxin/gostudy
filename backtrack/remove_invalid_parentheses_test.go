package backtrack

import (
	"testing"
)

func TestRemoveInvalidParentheses(t *testing.T) {

	var items = []struct {
		name  string
		input string
	}{
		// {"301 input ()())()", "()())()"},
		// {"301 input (a)())()", "(a)())()"},
		// {"301 input )(", ")("},
		// {"301 input ''", ""},
		// {"301 input n", "n"},
		// {"301 input ()", "()"},
		// {"301 input (((", "((("},
		{"301 input )d))", ")d))"},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			result := RemoveInvalidParentheses(item.input)

			t.Logf("[%d] %v", len(result), result)

			for _, item := range result {
				t.Logf("ITEM[%v]", []rune(item))
			}
		})
	}
}

func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func helper(ans *[]string, str string, start, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lremove+rremove > len(str)-i {
			return
		}
		// 尝试去掉一个左括号
		if lremove > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, lremove-1, rremove)
		}
		// 尝试去掉一个右括号
		if rremove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lremove, rremove-1)
		}
	}
}

func removeInvalidParentheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	helper(&ans, s, 0, lremove, rremove)
	return
}
