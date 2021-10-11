package heap

import "fmt"

type KeyValue struct {
	key string
	val string
}

func (kv KeyValue) String() string {
	return fmt.Sprintf("[%s]:%s", kv.key, kv.val)
}

func (kv KeyValue) Compare(v Value) int {
	if kv.val < v.(KeyValue).val {
		return -1
	} else if kv.val > v.(KeyValue).val {
		return 1
	}
	return 0
}
