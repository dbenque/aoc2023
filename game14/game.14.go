package game14

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Space []Column
type Index map[int]map[*Rock]struct{}
type Column Index
type Row Index
type Rock struct {
	x, y    int
	rolling bool
}

var columns = map[int]map[*Rock]struct{}{}
var rows = map[int]map[*Rock]struct{}{}

type Rocks struct {
	slice   []*Rock
	coord   func(*Rock) *int
	inverse bool
}

func (i Index) Print(n int) {
	for r := range i[n] {
		fmt.Printf("%#v\n", *r)
	}
}

// Len implements sort.Interface.
func (rs Rocks) Len() int {
	return len(rs.slice)
}

// Less implements sort.Interface.
func (rs Rocks) Less(i int, j int) bool {
	if rs.inverse {
		return *rs.coord(rs.slice[i]) > *rs.coord(rs.slice[j])
	}
	return *rs.coord(rs.slice[i]) < *rs.coord(rs.slice[j])
}

// Swap implements sort.Interface.
func (rs Rocks) Swap(i int, j int) {
	rs.slice[i], rs.slice[j] = rs.slice[j], rs.slice[i]
}

func (i Index) Move(f coordFunc, r *Rock, to int) {

	// cpR := *r
	source := i[*f(r)]
	destination := i[to]
	if destination == nil {
		destination = map[*Rock]struct{}{}
		i[to] = destination
	}
	delete(source, r)
	*f(r) = to
	destination[r] = struct{}{}

	//fmt.Printf("%#v  =(%d)=>  %#v\n", cpR, to, *r)
}

var space Space

func Game14_1() {
	ReadSpace()
	Score()
}

var height = 0
var length = 0

func ReadSpace() {
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
		height++
		line := fileScanner.Text()
		length = len(line)
		fmt.Printf("%s L=%d\n", line, length)
		for i, b := range line {
			if len(space) <= i {
				space = append(space, Column{})
			}
			if b == '.' {
				continue
			}
			r := &Rock{y: height - 1, x: i, rolling: b == 'O'}

			m, ok := columns[r.x]
			if !ok {
				m = map[*Rock]struct{}{}
			}
			m[r] = struct{}{}
			columns[r.x] = m

			m, ok = rows[r.y]
			if !ok {
				m = map[*Rock]struct{}{}
			}
			m[r] = struct{}{}
			rows[r.y] = m
		}
	}
}

type coordFunc func(*Rock) *int

func X(r *Rock) *int {
	return &r.x
}
func Y(r *Rock) *int {
	return &r.y
}

func PushNorth() {
	Push(columns, Y, 0, rows)
}
func PushWest() {
	Push(rows, X, 0, columns)
}
func PushSouth() {
	Push(columns, Y, height-1, rows)
}
func PushEast() {
	Push(rows, X, length-1, columns)
}

func Cycle() {
	PushNorth()
	PushWest()
	PushSouth()
	PushEast()
}

func Push(index Index, coord coordFunc, borderIndex int, shuffledIndex Index) {
	for _, col := range index {
		rocks := Rocks{coord: coord, inverse: borderIndex != 0}
		for r := range col {
			rocks.slice = append(rocks.slice, r)
		}
		sort.Sort(rocks)

		if borderIndex != 0 {
			border := borderIndex
			for _, r := range rocks.slice {
				if r.rolling && border > *coord(r) {
					shuffledIndex.Move(coord, r, border)
					border--
					continue
				}
				border = *coord(r) - 1
			}
		} else {
			top := 0
			for _, r := range rocks.slice {
				if r.rolling && top < *coord(r) {
					shuffledIndex.Move(coord, r, top)
					top++
					continue
				}
				top = *coord(r) + 1
			}
		}
	}
}

func Score() {
	loop := map[int]int{}
	cycles := 1000000000
	//cycles := 10
loopCycles:
	for i := 0; i < cycles; i++ {
		Cycle()

		sum := 0
		for _, c := range columns {
			for r := range c {
				if r.rolling {
					sum += height - r.y
				}
			}
		}
		fmt.Printf("%d => %d [%d]\n", i, sum, loop[sum])

		loop[sum] += 1

		if loop[sum] > 100 {
			fmt.Println("Looping, probable results:")
			for k, v := range loop {
				if v >= 48 {
					fmt.Println(k)
				}
			}
			break loopCycles
		}
	}
}
