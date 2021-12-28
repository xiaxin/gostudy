package corpus

type Word struct {
	str   string
	len   int
	runes []rune
}

func NewWord(str string) Word {
	runes := []rune(str)

	return Word{
		str:   str,
		len:   len(runes),
		runes: runes,
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
