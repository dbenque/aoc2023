package game5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Game5_1() {
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
	ReadFile(lines)
	target := "location"
	min := -1
	for _, id := range initialSeeds {
		// find relevant mapper
		idMapped := findMapped(id, "seed", target)
		fmt.Printf("%d ==( %s )==> %d\n", id, target, idMapped)
		if min == -1 || idMapped < min {
			min = idMapped
		}
	}
	fmt.Printf("Minimal %s = %d\n", target, min)
}

func findMapped(idSrc int, categorySrc, targetCategory string) int {

	if categorySrc == targetCategory {
		return idSrc
	}
	//fmt.Printf("Mappers %s=>%s  --------------\n", categorySrc, chainCategory[categorySrc])
	for _, m := range Mappers[categorySrc] {
		//fmt.Printf("mapper: %#v\n", *m)
		if m.Match(idSrc) {
			idDestination := m.MappedTo(idSrc)
			return findMapped(idDestination, chainCategory[categorySrc], targetCategory)
		}
	}
	//fmt.Printf("No mapping %s(%d) to %s\n", categorySrc, idSrc, chainCategory[categorySrc])
	// Map to same number
	return findMapped(idSrc, chainCategory[categorySrc], targetCategory)
}

func (m *ItemMapper) Match(id int) bool {
	return m.Source <= id && id <= m.Source+m.Length
}

func (m *ItemMapper) MappedTo(id int) int {
	return (id - m.Source) + m.Destination
}

var initialSeeds []int
var chainCategory = map[string]string{}

func ReadSeed(line string) {
	for _, str := range strings.Split(line[7:], " ") {
		id, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		initialSeeds = append(initialSeeds, id)
	}
}

func ReadCategoryMapper(lines []string) {
	mapText := strings.Split(lines[0], " ")[0]
	tokens := strings.Split(mapText, "-")
	sourceCategory := tokens[0]
	destinationCategory := tokens[2]

	chainCategory[sourceCategory] = destinationCategory

	if len(lines) == 1 {
		return
	}

	fmt.Printf("Reading %s to %s\n", sourceCategory, destinationCategory)

	for _, l := range lines[1:] {
		m := &ItemMapper{}
		fmt.Sscanf(l, "%d %d %d", &m.Destination, &m.Source, &m.Length)
		Mappers[sourceCategory] = append(Mappers[sourceCategory], m)
	}
}

func ReadFile(lines []string) {
	ReadSeed(lines[0])

	var group []string
	for _, l := range lines[2:] {
		if len(l) == 0 && group != nil {
			ReadCategoryMapper(group)
			group = nil
			continue
		}
		group = append(group, l)
	}
	if group != nil {
		ReadCategoryMapper(group)
	}
}

type ItemMapper struct {
	Source, Destination int
	Length              int
}

var Mappers = map[string][]*ItemMapper{}
