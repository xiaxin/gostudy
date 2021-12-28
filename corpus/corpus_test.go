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

	// 动词
	c.AddWord("来到", 20)
	// 地名
	c.AddWord("北京", 20)
	// 中华人民共和国
	c.AddWord("中华", 20)
	c.AddWord("人民", 20)
	c.AddWord("共和国", 20)
	c.AddWord("中华人民共和国", 20)
	// 清华大学
	c.AddWord("清华", 20)
	c.AddWord("华大", 20)
	c.AddWord("大学", 20)
	c.AddWord("清华大学", 20)
	c.AddWord("北京清华大学", 20)

}

func TestToSlice(t *testing.T) {
	words := c.ToSlice("你好,世界")

	for i, word := range words {
		log.Printf("index[%d] word:%s po:%s", i, word, word.pos)
	}
}

func TestCut(t *testing.T) {

	text := "我来到北京清华大学"
	t.Log(c.Cut(text))

	{
		text := "他来到了网易杭研大厦"
		t.Log(c.Cut(text))
	}

}

//  全模式测试
func TestCutAll(t *testing.T) {

	text := "我来到北京清华大学"

	result := c.CutAll(text)

	log.Println(result)
}
