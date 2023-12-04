package game2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Game2_2() {
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
	sum := 0

	for fileScanner.Scan() {
		if id := isPossible2(fileScanner.Text()); id > 0 {
			sum += id
		}
	}

	readFile.Close()
	fmt.Println(sum)
}

func isPossible2(line string) int {
	g := strings.Split(line, ":")
	// id,err:=strconv.Atoi(g[0][5:])
	// if err!=nil {
	// 	panic(err.Error()+" "+line)
	// }
	maxes := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, tirage := range strings.Split(g[1], ";") {

		for _, cube := range strings.Split(strings.ReplaceAll(tirage, " ", ""), ",") {
			for i, b := range []byte(cube) {
				if b > '9' || b < '0' {
					colour := cube[i:]
					count, err := strconv.Atoi(cube[0:i])
					if err != nil {
						fmt.Println(tirage)
						fmt.Println(cube)
						fmt.Println(i)
						panic(err)
					}
					if count > maxes[colour] {
						maxes[colour] = count
					}
					break
				}
			}
		}
	}
	p := 1
	for _, c := range maxes {
		p = p * c
	}

	return p
}
