package game3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var engine []string

func Game3_1() {
	if len(os.Args) < 2 {
		panic("missing input file as second parameter")
	}
	filePath := os.Args[2]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		engine = append(engine, fileScanner.Text())
	}

	readFile.Close()
	sum := 0
	for _, v := range ParseEngine() {
		sum += v
	}
	fmt.Println(sum)
}

func ParseEngine() []int {
	var numbers []int

	for l, line := range engine {
		processNumber := func(l, b, e int) {
			if FoundNumberWithSymbol(l, b, e) {
				n, err := strconv.Atoi(line[b:e])
				if err != nil {
					panic(err.Error() + ", line:" + line[b:e])
				}
				numbers = append(numbers, n)
			}
		}

		b := -1
		for i, c := range []byte(line) {
			if c >= '0' && c <= '9' {
				if b == -1 {
					b = i
				}
				continue
			}
			if b != -1 {
				processNumber(l, b, i)
				b = -1
			}
		}
		if b != -1 {
			processNumber(l, b, len(line))
		}
	}
	return numbers
}

func FoundNumberWithSymbol(l, b, e int) bool {
	// Check previous line
	if l > 0 {
		if SearchSymbolOnLine(engine[l-1], b, e) {
			return true
		}
	}
	// Check next line
	if l < len(engine)-1 {
		if SearchSymbolOnLine(engine[l+1], b, e) {
			return true
		}
	}
	// Check current line
	if SearchSymbolOnLine(engine[l], b, e) {
		return true
	}

	return false
}

func IsExpectedSymbol(c byte) bool {
	return c != '.' && (c < '0' || c > '9')
}

func SearchSymbolOnLine(line string, b, e int) bool {
	if b > 0 {
		b = b - 1
	}
	if e < len(line) {
		e = e + 1
	}

	for _, c := range ([]byte)(line)[b:e] {
		if IsExpectedSymbol(c) {
			return true
		}
	}
	return false
}
