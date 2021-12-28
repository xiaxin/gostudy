package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// 语料库
type corpus struct {
	words  map[string]int
	maxlen int
	total  float64
}

type prob struct {
	prob_k float64
	k      int
}

// newCorpus returns a new corpus struct initialized with nil values.
func newCorpus() *corpus {
	return &corpus{}
}

// newProb takes a prob_k and a k and returns a pointer to a prob struct.
func newProb(pk float64, k int) *prob {
	return &prob{
		prob_k: pk,
		k:      k,
	}
}

// max finds the maximum of two integers.
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// wordFreq populates the words map in a corpus struct with word frequencies.
// 计算词频
func (c *corpus) wordFreq(w []string) {
	out := make(map[string]int)

	var total int
	for _, x := range w {
		out[x] += 1
		total++
	}

	c.words = out
	c.total = float64(total)
}

// 添加词
func (c *corpus) addWord(w string, f int) {
	if len(w) > c.maxlen {
		c.maxlen = len(w)
	}
	if nil == c.words {
		c.words = make(map[string]int)
	}
	c.words[w] += f
	c.total += float64(f)
}

// wordProb calculates how probable word is in context of corpus c.
func (c *corpus) wordProb(word string) float64 {
	// log.Printf("word[%s] %d", word, c.words[word])
	return float64(c.words[word]) / c.total
}

// maxProb finds the largest prob_k value from a slice of prob structs and
// returns the prob_k, k.
func maxProb(ps []*prob) (float64, int) {
	var (
		prob_k float64
		k      int
	)

	for _, z := range ps {
		if z.prob_k > prob_k {
			prob_k = z.prob_k
			k = z.k
		}
	}

	return prob_k, k
}

// reverse returns slice s in reverse.
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// String specifies how a prob struct is to be represented as a string.
func (p prob) String() string {
	return fmt.Sprintf("(%v), %d", p.prob_k, p.k)
}

func (c *corpus) loadFiles(f string) {

	re := regexp.MustCompile("[A-Za-z]+")
	b, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	wb := re.FindAll(b, -1)
	if wb == nil {
		log.Fatal("Failed to extract words from corpus")
	}

	var (
		words  []string
		maxlen int
	)
	for _, w := range wb {
		if len(w) > maxlen {
			maxlen = len(w)
		}
		words = append(words, strings.ToLower(string(w)))
	}

	c.wordFreq(words)
	c.maxlen = maxlen
}

func (c *corpus) loadWords(f string) {

	re := regexp.MustCompile("[A-Za-z]+")

	wb := re.FindAllString(f, -1)
	if wb == nil {
		log.Fatal("Failed to extract words from corpus")
	}

	var (
		words  []string
		maxlen int
	)
	for _, w := range wb {
		if len(w) > maxlen {
			maxlen = len(w)
		}
		words = append(words, strings.ToLower(string(w)))
	}

	c.wordFreq(words)
	c.maxlen = maxlen
}

// viterbi
func (c *corpus) viterbi(text string) []string {
	var (
		probs = []float64{1.0}
		lasts = []int{0}
	)

	log.Printf("maxlen:%d", c.maxlen)

	for i := 1; i < len(text)+1; i++ {
		var y []*prob
		for j := max(0, i-c.maxlen); j < i; j++ {
			prob := newProb(probs[j]*c.wordProb(text[j:i]), j)
			y = append(y, prob)

			log.Printf("i[%d] j[%d] y[%d] %s, %d pre_prob:%f prob:%f, %d", i, j, len(y), text[j:i], j, probs[j], prob.prob_k, prob.k)
		}

		// 找到最大的 prob.prob_k 和 prob.k
		prob_k, k := maxProb(y)

		// log.Printf("prob %f %d", prob_k, k)

		probs = append(probs, prob_k)
		lasts = append(lasts, k)
		// log.Printf("lasts %v", lasts)
	}
	log.Printf("lasts %v", lasts)

	var (
		words []string
		i     = len(text)
	)

	for 0 < i {
		words = append(words, text[lasts[i]:i])
		i = lasts[i]
	}

	log.Printf("words %v", words)

	return reverse(words)
}

func main() {
	c := newCorpus()
	// c.loadWords("eng_news_2015_1M-sentences.txt")
	c.addWord("aa", 1)

	actual := c.viterbi("aacentraleuropeandebtprogrammeisbecauseofthebank")

	log.Println(actual)

	fmt.Println(c.viterbi("aaelevenpickleswentintotownandfoundapubopen"))
}
