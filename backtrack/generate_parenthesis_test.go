package backtrack

import "testing"

func TestParentheses(t *testing.T) {
	result := Parentheses(3)

	for _, v := range result {
		t.Log(v)
	}
}
