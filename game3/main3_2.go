package game3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var engine2 []string

func Game3_2() {
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

	for fileScanner.Scan() {
		engine2 = append(engine2, fileScanner.Text())
	}

	fmt.Printf("%#v\n", engine2)

	readFile.Close()
	sum := int64(0)
	allRows := ParseEngine2()
	for i, v := range allRows {
		var p, n *Row
		if i > 0 {
			p = allRows[i-1]
		}
		if i < len(allRows)-1 {
			n = allRows[i+1]
		}
		v.FindAdjacentNumbers(p, n)

		b, _ := json.MarshalIndent(v, " ", " ")
		fmt.Println(string(b))

		for _, g := range v.Gears {
			sum = sum + g.Compute()
		}
	}
	fmt.Println(sum)
}

type Number struct {
	Begin, End int
	Value      int64
}
type Gear struct {
	Symbol         byte
	Position       int
	ConnectedValue []int64
}

func (g *Gear) Compute() int64 {
	if g.Symbol == '*' && len(g.ConnectedValue) == 2 {
		p := int64(1)
		for _, v := range g.ConnectedValue {
			p = p * v
		}
		return p
	}
	// s := int64(0)
	// for _, v := range g.ConnectedValue {
	// 	s = s + v
	// }
	// return s
	return 0
}

type Row struct {
	Numbers []*Number
	Gears   []*Gear
}

var LineLength int

func (r *Row) AddNumber(b, e int, v int64) {
	r.Numbers = append(r.Numbers, &Number{Begin: b, End: e, Value: v})
}

func (r *Row) AddGear(p int, g byte) {
	r.Gears = append(r.Gears, &Gear{Position: p, Symbol: g})
}
func ParseEngine2() []*Row {
	var rows []*Row
	LineLength = len(engine2[0])
	for l, line := range engine2 {
		rows = append(rows, &Row{})
		b := -1
		processNumber := func(b, e int) {
			v, err := strconv.ParseInt(engine2[l][b:e], 10, 0)
			if err != nil {
				panic(err)
			}
			rows[l].AddNumber(b, e, v)
		}
		for i, c := range []byte(line) {
			if c >= '0' && c <= '9' {
				if b == -1 {
					b = i
				}
				continue
			}
			if b != -1 {
				processNumber(b, i)
				b = -1
			}
			if c != '.' {
				rows[l].AddGear(i, c)
			}
		}
		if b != -1 {
			processNumber(b, len(line))
		}
	}
	return rows
}

func (r *Row) FindAdjacentNumbers(previousRow, nextRow *Row) {
	for _, g := range r.Gears {
		g.ConnectedValue = append(g.ConnectedValue, r.TouchingNumbers(g.Position)...)
		if previousRow != nil {
			g.ConnectedValue = append(g.ConnectedValue, previousRow.TouchingNumbers(g.Position)...)
		}
		if nextRow != nil {
			g.ConnectedValue = append(g.ConnectedValue, nextRow.TouchingNumbers(g.Position)...)
		}
	}
}

func (r *Row) TouchingNumbers(position int) []int64 {
	var values []int64
	for _, n := range r.Numbers {
		b := n.Begin
		if b > 0 {
			b = b - 1
		}
		e := n.End
		if e < LineLength {
			e = e + 1
		}
		if position >= b && position < e {
			values = append(values, n.Value)
		}
	}
	return values
}
