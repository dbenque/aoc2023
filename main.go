package main

import (
	"aoc/game10"
	"aoc/game2"
	"aoc/game3"
	"aoc/game4"
	"aoc/game5"
	"aoc/game6"
	"aoc/game7"
	"aoc/game8"
	"aoc/game9"
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
	case "7.1":
		game7.Game7_1()
	case "7.2":
		game7.Game7_2()
	case "8.1":
		game8.Game8_1()
	case "8.2":
		game8.Game8_2()
	case "9.1":
		game9.Game9_1()
	case "9.2":
		game9.Game9_2()
	case "10":
		game10.Game10_1()

	default:
		panic("game not found: " + os.Args[1])
	}

}
