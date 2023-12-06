package game6

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Game6() {
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
	space := regexp.MustCompile(`\s+`)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, space.ReplaceAllString(fileScanner.Text(), " "))
	}

	times := strings.Split(lines[0], " ")[1:]
	distances := strings.Split(lines[1], " ")[1:]

	var races []*Race
	for i := range times {
		c, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])

		races = append(races, &Race{chrono: float64(c), distance: float64(d)})
	}

	p := 1
	for _, r := range races {
		s := r.WinningTimes()
		if len(s) == 0 {
			p = 0
			break
		}
		fmt.Printf("%#v  %#v  %d\n", r, s, s[1]-s[0]+1)
		p = p * (s[1] - s[0] + 1)
	}
	fmt.Println(p)

	/*
		totalTime = tbutton + trace
		totalDistance = trace*tbutton

		time=7
		distance=9

		7 = tbutton + trace  <=> trace= 7-tbutton
		9 < trace*tbutton    <=> 9 < (7 - tbutton)*tbutton

		=> 9 + (tbutton - 7)*tbutton < 0
		=> tbutton*tbutton - 7*tbutton + 9 < 0

		=> t*t - time*t + distance < 0
		=> delta = time*time - 4*distance      => si delta < 0 pas de solution
		                                       => si delta = 0 solution unique =time/2
											   => si delta > 0 sol1=(time*sqrt(delta))/2   sol2=(time-sqrt(delta))/2
	*/

}

type Race struct {
	chrono   float64
	distance float64
}

func (r *Race) WinningTimes() []int {
	delta := r.chrono*r.chrono - 4*r.distance
	if delta <= 0 {
		return nil
	}
	s2 := (r.chrono + math.Sqrt(delta)) / 2
	s1 := (r.chrono - math.Sqrt(delta)) / 2

	M := int(math.Floor(s2))
	m := int(math.Ceil(s1))

	fmt.Printf("%v %v            %v %v\n", M, s2, m, s1)

	if (s2 - float64(M)) == 0 {
		M = M - 1
	}
	if (s1 - float64(m)) == 0 {
		m = m + 1
	}

	return []int{m, M}
}
