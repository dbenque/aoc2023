package main

import (
	"aoc/game2"
	"aoc/game3"
	"aoc/game4"
	"aoc/game5"
	"aoc/game6"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		panic("missing game number (2.1,3.2,...)")
	}

	switch os.Args[1] {
	case "2.1":
		game2.Game2_1()
	case "2.2":
		game2.Game2_2()
	case "3.1":
		game3.Game3_1()
	case "3.2":
		game3.Game3_2()
	case "4.1":
		game4.Game4_1()
	case "4.2":
		game4.Game4_2()
	case "5.1":
		game5.Game5_1()
	case "5.2":
		game5.Game5_2()
	case "6":
		game6.Game6() // input.txt for 1 and input.2.txt for 2

	default:
		panic("game not found: " + os.Args[1])
	}

}
