package corpus

type Dict interface {
	Size() int
	Add(word *Word) error
	Del(word string) error
	Find(word string) *Word
	SumFreq() float64
}

func NewHashDict() Dict {
	return &hashDict{
		words:   make(map[string]*Word),
		size:    0,
		sumFreq: 0,
	}
}

type hashDict struct {
	words   map[string]*Word
	size    int
	sumFreq float64
}

func (d *hashDict) Size() int {
	return d.size
}

func (d *hashDict) Add(word *Word) error {
	d.words[word.str] = word
	d.sumFreq += word.freq
	return nil
}

func (d *hashDict) Del(word string) error {
	if w, ok := d.words[word]; ok {
		delete(d.words, w.str)
		d.sumFreq -= w.freq
	}
	return nil
}

func (d *hashDict) Find(word string) *Word {
	if w, ok := d.words[word]; ok {
		return w
	}
	return nil
}

func (d *hashDict) SumFreq() float64 {
	return d.sumFreq
}
