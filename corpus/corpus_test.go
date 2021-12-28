package corpus

import "testing"

func TestSplit(t *testing.T) {
	c := New()

	c.AddWord("你好", 20)
	c.AddWord("世界", 20)
	c.AddWord("你好世", 20)
	c.AddWord("亲爱的", 20)

	t.Log(c.splitDag("你好世界, Hello World, 亲爱的"))

}
