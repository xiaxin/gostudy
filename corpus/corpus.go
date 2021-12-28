package corpus

import (
	"math"
	"regexp"
	"strings"
)

const (
	// RatioWord ratio words and letters
	RatioWord float32 = 1.5
	ToLower           = true
)

var reEng = regexp.MustCompile(`[[:alnum:]]`)

type corpus struct {
	dict Dict
}

func New() *corpus {
	return &corpus{
		dict: NewHashDict(),
	}
}

func (c *corpus) AddWord(word string, freq int) error {
	return c.dict.Add(word, freq)
}

func (c *corpus) DelWord(word string) error {
	return c.dict.Del(word)
}

func (c *corpus) FindWord(word string) (Word, int) {
	freq := c.dict.Find(word)
	if freq == 0 {
		return NewWord(word), 0
	}

	return NewWord(word), freq

}

func (c *corpus) makeDag(text string) map[int][]int {
	words := NewWord(text)
	runes := words.ToRunes()

	dag := make(map[int][]int)

	n := words.Len()

	for i := 0; i < n; i++ {
		dag[i] = make([]int, 0)

		j := i

		temp := runes[i : j+1]

		for {
			_, freq := c.FindWord(string(temp))

			if freq != 0 {

				// log.Printf("[make-dag] i[%d] j[%d] %s %d \r\n", i, j, string(temp), freq)

				if freq > 0 {
					dag[i] = append(dag[i], j)
				}
			}

			j++

			if j >= n {
				break
			}
			temp = runes[i : j+1]
		}

		if len(dag[i]) == 0 {
			dag[i] = append(dag[i], i)
		}
	}

	return dag
}

func (c *corpus) calcDag(text string, dag map[int][]int) map[int]DagRoute {

	words := NewWord(text)
	runes := words.ToRunes()

	n := len(runes)
	rs := make(map[int]DagRoute)

	rs[n] = DagRoute{freq: 0.0, index: 0}
	var r DagRoute

	// 总频率的对数
	logT := math.Log(float64(c.dict.SumFreq()))
	// log.Println(c.dict.SumFreq(), logT)

	for idx := n - 1; idx >= 0; idx-- {
		for _, i := range dag[idx] {
			temp := runes[idx : i+1]
			freq := float64(c.dict.Find(string(temp)))

			// log.Printf("[calc-dag] idx[%d] i[%d] freq:%f  runes[idx : i+1]:%s", idx, i, freq, string(runes[idx:i+1]))

			if freq > 0 {
				f := math.Log(freq) - logT + rs[i+1].freq
				r = DagRoute{freq: f, index: i}
			} else {
				f := math.Log(1.0) - logT + rs[i+1].freq
				r = DagRoute{freq: f, index: i}
			}

			// log.Printf("[calc-dag] idx[%d] i[%d] r.freq:%f  runes[idx : i+1]:%s", idx, i, r.freq, string(runes[idx:i+1]))

			if v, ok := rs[idx]; !ok {
				rs[idx] = r
			} else {
				f := v.freq == r.freq && v.index < r.index
				if v.freq < r.freq || f {
					rs[idx] = r
				}
			}
			// log.Printf("[calc-dag] idx[%d] i[%d] route:%f %d  runes[idx : i+1]:%s", idx, i, r.freq, r.index, string(runes[idx:i+1]))
		}
	}

	return rs
}

// 使用DAG分词
func (c *corpus) splitDag(text string) []string {

	mLen := int(float32(len(text))/RatioWord) + 1
	result := make([]string, 0, mLen)

	if ToLower {
		// 英语字母小写
		text = strings.ToLower(text)
	}
	runes := []rune(text)

	dag := c.makeDag(text)

	routes := c.calcDag(text, dag)
	// 字符长度
	length := len(runes)

	var y int
	var buf []rune

	for x := 0; x < length; {
		y = routes[x].index + 1
		frag := runes[x:y]

		// log.Printf("[splite dag] frag:%s", string(frag))

		if reEng.MatchString(string(frag)) && len(frag) == 1 {
			buf = append(buf, frag...)
			x = y
			continue
		}

		if len(buf) > 0 {
			result = append(result, string(buf))
			buf = make([]rune, 0)
		}

		result = append(result, string(frag))
		x = y
	}

	if len(buf) > 0 {
		result = append(result, string(buf))
	}

	return result

}
