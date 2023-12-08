package game7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Game7_1() {
	Game7(DecodeHand)
}

func Game7(decoder func([5]byte) *Hand) {

	if len(os.Args) < 2 {
		panic("missing input file as second parameter")
	}
	filePath := os.Args[2]
	readFile, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var hands []*Hand
	for fileScanner.Scan() {
		token := strings.Split(fileScanner.Text(), " ")
		bid, _ := strconv.Atoi(token[1])
		h := DecodeHandStr(token[0], decoder)
		h.Bid = bid
		h.Str = token[0]
		hands = append(hands, h)
	}

	sort.Sort((SortByScore)(hands))
	sum := 0
	for i, h := range hands {
		fmt.Printf("%#v %s\n", *h, h.Str)
		sum += (i + 1) * h.Bid
	}
	fmt.Println(sum)
}

type SortByScore []*Hand

func (a SortByScore) Len() int      { return len(a) }
func (a SortByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByScore) Less(i, j int) bool {
	if a[i].Type < a[j].Type {
		return true
	}
	if a[j].Type < a[i].Type {
		return false
	}

	for k := 0; k < 5; k++ {
		if a[i].Cards[k] < a[j].Cards[k] {
			return true
		}
		if a[j].Cards[k] < a[i].Cards[k] {
			return false
		}
	}
	return false
}

type Hand struct {
	Cards [5]byte
	Type  string
	Bid   int
	Str   string
}

var tmpHand [5]byte

func DecodeHandStr(str string, decoder func([5]byte) *Hand) *Hand {
	for i := 0; i < 5; i++ {
		tmpHand[i] = str[i]
	}
	return decoder(tmpHand)
}

func DecodeHand(s [5]byte) *Hand {
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
			v = 11
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
	var counts []byte
	for _, v := range indexed {
		counts = append(counts, v+'0')
	}
	sort.Slice(counts, func(i int, j int) bool { return counts[i] > counts[j] })
	h.Type = string(counts)
	return h
}
