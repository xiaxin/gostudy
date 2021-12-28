package corpus

type Dict interface {
	Size() int
	Add(word string, freq int) error
	Del(word string) error
	Find(word string) int
	SumFreq() int
}

func NewHashDict() Dict {
	return &hashDict{
		words: make(map[string]int),
	}
}

type hashDict struct {
	words   map[string]int
	size    int
	sumFreq int
}

func (d *hashDict) Size() int {
	return d.size
}

func (d *hashDict) Add(word string, freq int) error {
	d.words[word] += freq
	d.sumFreq += freq
	return nil
}

func (d *hashDict) Del(word string) error {
	if v, ok := d.words[word]; ok {
		delete(d.words, word)
		d.sumFreq -= v
	}
	return nil
}

func (d *hashDict) Find(word string) int {
	if v, ok := d.words[word]; ok {
		return v
	}
	return 0
}

func (d *hashDict) SumFreq() int {
	return d.sumFreq
}
