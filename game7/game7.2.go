package game7

import (
	"sort"
)

func Game7_2() {

	Game7(DecodeHand2)

}

func DecodeHand2(s [5]byte) *Hand {
	h := &Hand{}
	indexed := map[byte]byte{}
	for i := 0; i < 5; i++ {
		c := s[i]
		if c >= '2' && c <= '9' {
			v := c - '2' + 2
			h.Cards[i] = v
			indexed[v] = indexed[v] + 1
			continue
		}
		var v byte
		switch s[i] {
		case 'T':
			v = 10
		case 'J':
			v = 0
		case 'Q':
			v = 12
		case 'K':
			v = 13
		case 'A':
			v = 14
		default:
			panic("unknow card ")
		}
		h.Cards[i] = v
		indexed[v] = indexed[v] + 1
	}

	if len(indexed) == 1 && indexed[0] == 5 {
		indexed[14] = 5
		delete(indexed, 0)
	}

	var ccs []*CardCount
	var jCount byte
	for s, c := range indexed {
		if s == 0 {
			jCount = c
			continue
		}
		ccs = append(ccs, &CardCount{Count: c, Symbol: s})
	}

	sort.Sort((SortCardCount)(ccs))

	indexed[ccs[0].Symbol] = indexed[ccs[0].Symbol] + jCount
	delete(indexed, 0)

	var counts []byte
	for _, v := range indexed {
		counts = append(counts, v+'0')
	}
	sort.Slice(counts, func(i int, j int) bool { return counts[i] > counts[j] })
	h.Type = string(counts)
	return h
}

type SortCardCount []*CardCount

func (a SortCardCount) Len() int      { return len(a) }
func (a SortCardCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortCardCount) Less(i, j int) bool {
	if a[i].Count != a[j].Count {
		return a[i].Count > a[j].Count
	}
	return int(a[i].Symbol) > int(a[j].Symbol)
}

type CardCount struct {
	Count  byte
	Symbol byte
}
