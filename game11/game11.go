package game11

import (
	"bufio"
	"fmt"
	"os"
)

func Game11_1() {
	ReadSpace(1)
	fmt.Println(Distance())
}

func Game11_2() {
	ReadSpace(1000000 - 1)
	fmt.Println(Distance())
}

func Distance() (d int) {
	for i := range space {
		for j := i + 1; j < len(space); j++ {
			d += space[i].distance(space[j])

		}
	}
	return
}

type Space []*Galaxy

type Galaxy struct {
	x, y int
}

var space Space
var columns = map[int]bool{}
var rows = map[int]bool{}

var wide int
var height int

func ReadSpace(expension int) {
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
	y := 0
	for fileScanner.Scan() {
		x := 0
		line := fileScanner.Text()
		wide = len(line)
		for _, b := range line {
			if b == '#' {
				space = append(space, &Galaxy{x, y})
				columns[x] = true
				rows[y] = true
			}
			x++
		}
		y++
	}
	height = y
	addX := make([]int, wide)
	v := 0
	for i := 0; i < wide; i++ {
		if !columns[i] {
			v += expension
		}
		addX[i] = v
	}
	addY := make([]int, height)
	v = 0
	for i := 0; i < height; i++ {
		if !rows[i] {
			v += expension
		}
		addY[i] = v
	}
	for _, g := range space {
		g.x += addX[g.x]
		g.y += addY[g.y]
	}
	// for _, g := range space {
	// 	fmt.Printf("%#v\n", *g)
	// }
}

func (g *Galaxy) distance(other *Galaxy) int {
	s := 0
	x0, y0 := g.x, g.y
	x1, y1 := other.x, other.y

	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}

	return y1 - y0 + x1 - x0
	// for !(x0 == x1 && y0 == y1) {
	// 	s++
	// 	if x1-x0 > y1-y0 {
	// 		x0++
	// 	} else {
	// 		y0++
	// 	}
	// }
	return s
}
