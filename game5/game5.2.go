package game5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Game5_2() {
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
	ReadFile2(lines)

	targetCategory := "location"
	var allMappedRanges []*RangeItem
	for _, sr := range seedRanges {
		mappedRanges := findMappedRanges([]*RangeItem{sr}, "seed", targetCategory)
		allMappedRanges = append(allMappedRanges, mappedRanges...)
	}
	sort.Sort((SortByStart)(allMappedRanges))
	fmt.Println(allMappedRanges[0].Start)
}

func findMappedRanges(ranges []*RangeItem, category, targetcategory string) []*RangeItem {
	if category == targetcategory {
		return ranges
	}
	mappedRanges := TransformAllRangesForMappers(ranges, Mappers[category])
	mappedCategory := chainCategory[category]
	return findMappedRanges(mappedRanges, mappedCategory, targetcategory)
}

func ReadFile2(lines []string) {
	ReadSeed2(lines[0])

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

func ReadSeed2(line string) {
	var numbers []int
	for _, str := range strings.Split(line[7:], " ") {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	for i := 0; i < len(numbers); i = i + 2 {
		ri := &RangeItem{Start: numbers[i], End: numbers[i] + numbers[i+1] - 1}
		seedRanges = append(seedRanges, ri)
	}
}

var seedRanges []*RangeItem

type RangeItem struct {
	Start, End int // bound are included in the range
	mapper     *ItemMapper
}

type SortByStart []*RangeItem

func (a SortByStart) Len() int           { return len(a) }
func (a SortByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

func TransformAllRangesForMappers(inputs []*RangeItem, mappers []*ItemMapper) (result []*RangeItem) {
	var ranges []*RangeItem
	for _, r := range inputs {
		ranges = append(ranges, TransformRangeForMappers(r, mappers)...)
	}

	// apply transformation
	for _, r := range ranges {
		result = append(result, ApplyInnerMapper(r))
	}

	return
}

func ApplyInnerMapper(r *RangeItem) *RangeItem {
	if r.mapper == nil {
		return r
	}

	// Check
	if !r.mapper.Match(r.Start) {
		panic("Mapper does not intercept Start of Range")
	}
	if !r.mapper.Match(r.End) {
		panic("Mapper does not intercept End of Range")
	}

	return &RangeItem{Start: r.mapper.MappedTo(r.Start), End: r.mapper.MappedTo(r.End)}

}

func TransformRangeForMappers(r *RangeItem, mappers []*ItemMapper) []*RangeItem {
	var mRanges []*RangeItem
	for _, m := range mappers {
		mRanges = append(mRanges, &RangeItem{Start: m.Source, End: m.Source + m.Length - 1, mapper: m})
	}
	return transformRange(r, mRanges)
}
func transformRange(r *RangeItem, mappersRanges []*RangeItem) []*RangeItem {

	var mapped []*RangeItem
	for _, m := range mappersRanges {
		if v := InterceptRange(r, m); v != nil {
			mapped = append(mapped, v)
		}
	}

	sort.Sort((SortByStart)(mapped))

	// check begining
	if len(mapped) == 0 {
		return []*RangeItem{{Start: r.Start, End: r.End}} // the range is unchanged
	}

	// Check Start and complete with initial range if needed
	if mapped[0].Start > r.Start {
		mapped = append([]*RangeItem{{Start: r.Start, End: mapped[0].Start - 1}}, mapped...)
	}

	// Check End and complete with terminal range if needed
	if mapped[len(mapped)-1].End < r.End {
		mapped = append(mapped, &RangeItem{Start: mapped[len(mapped)-1].End + 1, End: r.End})
	}
	result := []*RangeItem{mapped[0]}
	for i := 0; i < len(mapped)-1; i++ {
		if mapped[i+1].Start != result[len(result)-1].End+1 {
			result = append(result, &RangeItem{Start: result[len(result)-1].End + 1, End: mapped[i+1].Start - 1})
		}
		result = append(result, mapped[i+1])
	}

	return result

}

func InterceptRange(r, spliter *RangeItem) *RangeItem {

	if spliter.End < r.Start || spliter.Start > r.End {
		return nil
	}

	startMax := r.Start
	if spliter.Start > startMax {
		startMax = spliter.Start
	}
	endMin := r.End
	if spliter.End < endMin {
		endMin = spliter.End
	}
	return &RangeItem{Start: startMax, End: endMin, mapper: spliter.mapper}
}

// Pour chaque niveau, j'ai un ensemble Range (sauf niveau Seed ou j'en ai un, je les travaille 1 par 1)

// Pour chaque Range, je decoupe en sous-ranges pour associé un mapper à chaque sous-range. Je vérifie la continuité de l'ensemble des sous-ranges. Au besoin je complète avec des sous-range ayant le mapper identité

// J'applique le mapping à chaque sous-range. J'ai un ensemble de Range => je recommence
