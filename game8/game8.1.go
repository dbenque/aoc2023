package game8

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() {
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
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	sequence = lines[0]

	readGraph(lines[2:])
	CheckAllNodes()
}

func Game8_1() {

	ReadInput()

	n := GetNode("AAA")
	var i, j int
	for {
		n, j = walkSequence(n, "ZZZ")
		if n.Name == "ZZZ" {
			break
		}
		i++
	}
	fmt.Println(i*len(sequence) + j)
}

func Game8_2() {
	ReadInput()
	fmt.Println(len(nodes))
	fmt.Println(len(sequence))
	n := GetNodesEndingBy('A')[0:]
	for _, v := range n {
		findLoop(v)
	}
}

var sequence string
var nodes = map[string]*Node{}

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func GetNode(name string) *Node {
	if n, ok := nodes[name]; ok {
		return n
	}
	n := &Node{Name: name}
	nodes[name] = n
	return n
}

func GetNodesEndingBy(a byte) (r []*Node) {
	for n, v := range nodes {
		if n[2] == a {
			r = append(r, v)
		}
	}
	return r
}

func readGraph(lines []string) {
	for _, l := range lines {
		name := l[0:3]
		n := GetNode(name)
		n.Left = GetNode(l[7:10])
		n.Right = GetNode(l[12:15])
	}
}

func walkSequence(n *Node, target string) (*Node, int) {
	for j, d := range []byte(sequence) {
		switch d {
		case 'L':
			n = n.Left
		case 'R':
			n = n.Right
		}
		if n.Name == target {
			return n, j + 1
		}
	}
	return n, -1
}

func walkSequence2(n []*Node, target byte) ([]*Node, int) {
	for j, d := range []byte(sequence) {
		switch d {
		case 'L':
			for i := range n {
				n[i] = n[i].Left
			}
		case 'R':
			for i := range n {
				n[i] = n[i].Right
			}
		}
		allTarget := true
		for _, v := range n {
			if v.Name[2] != target {
				allTarget = false
				break
			}
		}
		if allTarget {
			return n, j + 1
		}
	}
	return n, -1
}

var visit = map[*Node]map[int]int{}

func findLoop(n *Node) {
	visit = map[*Node]map[int]int{}
	var i, z int
	for {
		for j, d := range []byte(sequence) {
			switch d {
			case 'L':
				n = n.Left
			case 'R':
				n = n.Right
			}
			if n.Name[2] == 'Z' {
				z++
			}
			m, ok := visit[n]
			if !ok {
				m = map[int]int{}
				visit[n] = m
			}
			if v, ok := m[j]; ok {
				fmt.Printf("Loop %s from %d->%d at index %d, z=%d\n", n.Name, v, i, j, z)
				return
			}
			m[j] = i
		}
		i++
	}
}

func CheckAllNodes() {
	for k, v := range nodes {
		if k != v.Name {
			panic("NAME Violation")
		}

		if v.Left == nil {
			panic("LEFT is nil in " + k)
		}

		if v.Right == nil {
			panic("RIGHT is nil in " + k)
		}
	}
	fmt.Println("Check Node Ok")
}
