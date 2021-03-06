package backtrack

import "testing"

func TestGenerateParentheses(t *testing.T) {

	var items = []struct {
		name  string
		input int
	}{
		{"22 Input 3", 3},
	}

	for _, item := range items {

		t.Run(item.name, func(t *testing.T) {

			result := GenerateParentheses(item.input)

			for _, v := range result {
				t.Log(v)
			}
		})
	}

}
