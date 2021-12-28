package corpus

type Word struct {
	str   string
	len   int
	runes []rune
	freq  float64
	pos   string
}

func NewWord(str string, freq float64, pos ...string) *Word {
	runes := []rune(str)

	var po string

	if len(pos) > 0 {
		po = pos[0]
	}
	return &Word{
		str:   str,
		len:   len(runes),
		runes: runes,
		freq:  freq,
		pos:   po,
	}
}

func (w Word) Len() int {
	return w.len
}

func (w Word) String() string {
	return w.str
}

func (w Word) ToRunes() []rune {
	return w.runes
}
