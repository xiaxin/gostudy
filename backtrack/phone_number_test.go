package backtrack

import "testing"

var PhoneNumberTestData = []string{
	"2",
}

func TestPhoneNumber(t *testing.T) {

	for _, input := range PhoneNumberTestData {
		t.Run("Queue", func(t *testing.T) {
			list := PhoneNumberQueue(input)
			t.Log(list)
		})

		t.Run("Backtrack", func(t *testing.T) {
			list := PhoneNumberBacktrack(input)
			t.Log(list)
		})
	}
}
