package backtrack

import (
	"gostudy/common"
)

var PhoneNumberKeymap = map[rune][]rune{

	'1': {},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

//  基于回溯的解法
func PhoneNumberBacktrack(input string) []string {

	var result []string

	phoneNumberBacktrack("", []rune(input), &result)

	return result
}

func phoneNumberBacktrack(base string, input []rune, result *[]string) {

	if len(input) == 0 {
		*result = append(*result, base)
		return
	}

	for _, c := range PhoneNumberKeymap[input[0]] {
		phoneNumberBacktrack(base+string(c), input[1:], result)
	}

}

// 基于队列的解法 TODO 优化代码
func PhoneNumberQueue(input string) []string {

	queue := common.NewList()

	inputRune := []rune(input)

	for i := 0; i < len(inputRune); i++ {
		if inputRune[i] == '1' {
			continue
		}

		phoneNumberQueue(PhoneNumberKeymap[inputRune[i]], queue)
	}

	var result []string
	for queue.Size() > 0 {
		node := queue.Pop()
		result = append(result, node.ValString())
	}

	return result
}

func phoneNumberQueue(list []rune, queue *common.List) []string {
	var (
		result []string
		bases  []string
	)

	for queue.Size() > 0 {
		node := queue.Pop()
		bases = append(bases, node.ValString())

	}
	if len(bases) == 0 {
		bases = append(bases, "")
	}

	for _, base := range bases {
		for _, v := range list {
			queue.PushNode(0, base+string(v))
		}
	}

	return result
}
