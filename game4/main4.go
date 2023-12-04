package game4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Game4_1() {
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

	var cards []*Card
	for fileScanner.Scan() {
		cards = append(cards, DecodeCard(fileScanner.Text()))
	}

	sum := 0
	for _, c := range cards {
		c.SearchMatch()
		fmt.Printf("%v => %d\n", c.WinnersMatch, c.Points())
		sum = sum + c.Points()
	}
	fmt.Println(sum)
}

func Game4_2() {
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

	var cards []*Card
	for fileScanner.Scan() {
		cards = append(cards, DecodeCard(fileScanner.Text()))
	}

	for _, c := range cards {
		c.SearchMatch()
	}
	for i, c := range cards {
		m := len(c.WinnersMatch)
		for j := i + 1; j < len(cards) && m > 0; j++ {
			cc := cards[j]
			cc.InstanceCount += c.InstanceCount
			m--
		}
	}
	sum := 0
	for _, c := range cards {
		sum = sum + c.InstanceCount
		fmt.Println(c.InstanceCount)
	}
	fmt.Println("----")
	fmt.Println(sum)
}

type Card struct {
	Numbers       map[int]struct{}
	Winners       map[int]struct{}
	WinnersMatch  []int
	InstanceCount int
}

func (c *Card) SearchMatch() {
	for w := range c.Winners {
		if _, ok := c.Numbers[w]; ok {
			c.WinnersMatch = append(c.WinnersMatch, w)
		}
	}
}

func (c *Card) Points() int {
	if len(c.WinnersMatch) == 0 {
		return 0
	}
	return int(math.Pow(2, float64(len(c.WinnersMatch)-1)))
}

func DecodeCard(line string) *Card {
	c := &Card{
		Winners:       map[int]struct{}{},
		Numbers:       map[int]struct{}{},
		InstanceCount: 1,
	}
	fields := strings.Split(strings.Split(line, ":")[1], "|")
	numbers := strings.Split(strings.TrimSpace(fields[0]), " ")
	for _, str := range numbers {
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		c.Numbers[n] = struct{}{}
	}
	winners := strings.Split(strings.TrimSpace(fields[1]), " ")
	for _, str := range winners {
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}

		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		c.Winners[n] = struct{}{}
	}
	return c
}
