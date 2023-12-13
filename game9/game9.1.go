package game9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	Numbers       []int
	ChildRow      *Row
	CompleteCount int
}

type Oasis []*Row

var mainOasis Oasis

func Game9_1() {
	ReadOasis()
	sum := 0
	for _, r := range mainOasis {
		ExpandRow(r)
		CompleteRow(r)
		v := r.Numbers[len(r.Numbers)-1]
		fmt.Println(v)
		sum += v
	}
	fmt.Println(sum)
}

func Game9_2() {
	ReadOasis()
	sum := 0
	for _, r := range mainOasis {
		ExpandRow(r)
		CompleteRow2(r)
		fmt.Printf("%#v\n", r.Numbers[0])

		sum += r.Numbers[0]

	}
	fmt.Printf("Sum= %d\n", sum)
}

func ReadOasis() {

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
		r := &Row{}
		for _, v := range strings.Split(fileScanner.Text(), " ") {
			i, _ := strconv.Atoi(v)
			r.Numbers = append(r.Numbers, i)
		}
		mainOasis = append(mainOasis, r)
	}
}

func ExpandRow(r *Row) {
	if r.IsZero() {
		return
	}
	r.ChildRow = &Row{}
	r.ChildRow.Numbers = make([]int, len(r.Numbers)-1)
	for i := 0; i < len(r.Numbers)-1; i++ {
		r.ChildRow.Numbers[i] = r.Numbers[i+1] - r.Numbers[i]
	}
	ExpandRow(r.ChildRow)
}

func CompleteRow(r *Row) {
	if r.IsZero() {
		r.Numbers = append(r.Numbers, 0)
		return
	}
	CompleteRow(r.ChildRow)
	r.Numbers = append(r.Numbers, r.Numbers[len(r.Numbers)-1]+r.ChildRow.Numbers[len(r.ChildRow.Numbers)-1])
}

func (r *Row) IsZero() bool {
	for i := range r.Numbers {
		if r.Numbers[i] != 0 {
			return false
		}
	}
	return true
}

func CompleteRow2(r *Row) {
	if r.IsZero() {
		r.Numbers = append([]int{0}, r.Numbers...)
		return
	}
	CompleteRow2(r.ChildRow)
	v := []int{r.Numbers[0] - r.ChildRow.Numbers[0]}

	fmt.Printf("%d %v\n", v, r.Numbers)

	r.Numbers = append(v, r.Numbers...)
	r.CompleteCount++
}
