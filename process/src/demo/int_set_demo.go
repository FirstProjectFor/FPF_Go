package main

//IntSet innner type use int64
type IntSet struct {
	words []uint64
}

// Has Set exsists or not
func (s IntSet) Has(v int) bool {
	word, bit := wordAndBit(v)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//Add add value to set
func (s IntSet) Add(v int) {
	word, bit := wordAndBit(v)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func wordAndBit(v int) (int, uint) {
	return v / 64, uint(v % 64)
}

func main() {

}
