package game10

import (
	"bufio"
	"fmt"
	"os"
)

func Game10_1() {

	ReadPipes()

}

type Pipe struct {
	x, y                     int
	in, out                  *Pipe
	top, bottom, left, right bool
}

var start *Pipe

type Maze [][]*Pipe

var maze Maze

func ReadPipes() {
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
		var r []*Pipe
		l := fileScanner.Text()
		for x, b := range l {
			p := &Pipe{x: x, y: y}
			switch b {
			// | is a vertical pipe connecting north and south.
			// - is a horizontal pipe connecting east and west.
			// L is a 90-degree bend connecting north and east.
			// J is a 90-degree bend connecting north and west.
			// 7 is a 90-degree bend connecting south and west.
			// F is a 90-degree bend connecting south and east.
			// . is ground; there is no pipe in this tile.
			// S is the starting position of the animal; there is a pipe on this
			case '|':
				p.bottom = true
				p.top = true
			case '-':
				p.left = true
				p.right = true
			case 'L':
				p.top = true
				p.right = true
			case 'J':
				p.top = true
				p.left = true
			case '7':
				p.left = true
				p.bottom = true
			case 'F':
				p.right = true
				p.bottom = true
			case 'S':
				p.bottom = true
				p.top = true
				p.left = true
				p.right = true
				start = p
			}
			r = append(r, p)
		}
		y++
		maze = append(maze, r)
	}
	Walk(&maze)

	sum := 0
	for _, r := range maze {
		s := TraverseRow(r)
		sum += s
	}
	fmt.Printf("Inside elements: %d\n", sum)
}

func TraverseRow(r []*Pipe) int {
	insideCount := 0
	inside := false
	top, bottom := false, false
	for _, p := range r {
		if !p.OnLoop() {
			if inside {
				insideCount++
			}
			continue
		}
		if p.top {
			top = !top
		}
		if p.bottom {
			bottom = !bottom
		}
		inside = top && bottom
	}
	return insideCount
}

// OnLoop must be called only after Walk has been completed to detect the loop
func (p *Pipe) OnLoop() bool {
	return p.in != nil && p.out != nil
}

func (m *Maze) Top(p *Pipe) *Pipe {
	if p.y == 0 {
		return nil
	}
	return (*m)[p.y-1][p.x]
}
func (m *Maze) Bottom(p *Pipe) *Pipe {
	if p.y == len(*m)-1 {
		return nil
	}
	return (*m)[p.y+1][p.x]
}
func (m *Maze) Left(p *Pipe) *Pipe {
	if p.x == 0 {
		return nil
	}
	return (*m)[p.y][p.x-1]
}
func (m *Maze) Right(p *Pipe) *Pipe {
	if p.x == len((*m)[0])-1 {
		return nil
	}
	return (*m)[p.y][p.x+1]
}

func Walk(m *Maze) {
	p := start
	count := 0
	for {
		p = Move(m, p)
		count++
		if p == start {
			break
		}
	}
	fmt.Printf("Move count to make the loop: %d => max distance: %d\n", count, count/2)
}

func Move(m *Maze, p *Pipe) *Pipe {
	if p.top {
		next := m.Top(p)
		if next != nil && next != p.in {
			p.out = next
			next.in = p
			return next
		}
	}
	if p.bottom {
		next := m.Bottom(p)
		if next != nil && next != p.in {
			p.out = next
			next.in = p
			return next
		}
	}
	if p.left {
		next := m.Left(p)
		if next != nil && next != p.in {
			p.out = next
			next.in = p
			return next
		}
	}
	if p.right {
		next := m.Right(p)
		if next != nil && next != p.in {
			p.out = next
			next.in = p
			return next
		}
	}
	panic(fmt.Sprintf("No issue from %#v", p))
}
