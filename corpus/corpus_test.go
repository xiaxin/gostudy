package corpus

import (
	"log"
	"testing"
)

var c *corpus

func init() {

	c = New()

	c.AddWord("你好", 20, "n")
	c.AddWord("世界", 20, "n")
	c.AddWord("如果", 20)
	c.AddWord("你好世", 20)
	c.AddWord("亲爱的", 20)

}

func TestSplit(t *testing.T) {

	// t.Log(c.splitDag("你如果"))
	t.Log(c.splitDag("你好,世界"))

}

func TestToSlice(t *testing.T) {
	words := c.ToSlice("你好,世界")

	for i, word := range words {
		log.Printf("index[%d] word:%s po:%s", i, word, word.pos)
	}
}
