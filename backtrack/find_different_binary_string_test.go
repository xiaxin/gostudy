package backtrack

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {

	var items = []struct {
		name  string
		input []string
	}{
		{"1980 input [\"01\",\"10\"]", []string{"01", "10"}},
		{"1980 input [\"10\",\"01\"]", []string{"10", "01"}},
		{"1980 input [\"00\",\"01\"]", []string{"00", "01"}},
		{"1980 input [\"00\",\"01\"]", []string{"11", "00"}},
		{"1980 input [\"111\",\"011\",\"001\"]", []string{"111", "011", "001"}},
		{"1980 input [\"0\"]", []string{"0"}},
		{"1980 input [\"0\"]", []string{"00"}},
	}

	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {

			t.Run("FindDifferentBinaryString/B", func(t *testing.T) {
				result := findDifferentBinaryStringByLeetcode(item.input)
				t.Log(result)
			})

			t.Run("FindDifferentBinaryString/D", func(t *testing.T) {
				result := findDifferentBinaryString(item.input)

				t.Log(result)
			})
		})
	}
}
