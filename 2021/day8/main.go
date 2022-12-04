package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	displays := readInput()
	fmt.Printf("part1: %d\n", part1(displays))
	fmt.Printf("part2: %d\n", part2(displays))
}

func part1(displays []Display) int {
	sum := 0
	for _, display := range displays {
		for _, output := range display.output {
			switch len(output) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	return sum
}

func part2(displays []Display) int {
	return Decode(displays)

}

// Decode returns the segment value of the display
func Decode(displays []Display) int {
	total := 0
	for _, display := range displays {
		var TOP byte
		var TOP_RIGHT byte
		var BOT_RIGHT byte
		var BOT_LEFT byte
		var MIDDLE byte
		var TOP_LEFT byte
		var BOT byte
		one := ""
		seven := ""
		eight := ""
		four := ""
		six := ""
		five := ""
		nine := ""
		zero := ""
		// first let's compare 1 and 7 to find the top segment
		for _, pattern := range display.patterns {
			if len(pattern) == 7 {
				eight = pattern
			}
			if len(pattern) == 4 {
				four = pattern
			}
			if len(pattern) == 3 {
				// we found a seven
				seven = pattern
			}
			if len(pattern) == 2 {
				// we found a one
				one = pattern
			}
		}

		for _, letter := range seven {
			if letter == rune(one[0]) || letter == rune(one[1]) {
				// found a match
				continue
			}
			// found our top
			TOP = byte(letter)
		}

		// find 6 -> six chars long and will only have one of the two letters in 1
		for _, pattern := range display.patterns {
			if len(pattern) == 6 {
				if strings.ContainsRune(pattern, rune(one[0])) && strings.ContainsRune(pattern, rune(one[1])) { // this is NOT 6 bc we found entire 1 inside it
					continue
				}
				// we have gone through the other length 6 patterns, this one has to be 6
				six = pattern
				break
			}
		}

		// look through 6 to find the BOT_RIGHT
		for _, letter := range six {
			if letter == rune(one[0]) {
				BOT_RIGHT = byte(letter)
				TOP_RIGHT = one[1]
			} else if letter == rune(one[1]) {
				BOT_RIGHT = byte(letter)
				TOP_RIGHT = one[0]
			}
		}

		// now we can look at 6 to find 5 since 6 will include all of 5!
		foundFive := false
		for _, pattern := range display.patterns {
			if len(pattern) == 5 {
				five = ""
				for _, p := range pattern {
					if strings.ContainsRune(six, p) {
						five += string(p)
					}
					if len(five) == 5 {
						foundFive = true
						break
					}
				}
				if foundFive {
					break
				}
			}
		}

		// compare 5 and 6 to find BOT_LEFT
		for _, letter := range six {
			if !strings.ContainsRune(five, letter) {
				BOT_LEFT = byte(letter)
			}
		}

		// we can find 9 since we know its all of 8 except BOT_LEFT
		for _, pattern := range display.patterns {
			if len(pattern) == 6 {
				if pattern != six && !strings.ContainsRune(pattern, rune(BOT_LEFT)) {
					nine = pattern
				}
			}
		}

		// we can also find 0 bc it is the last remaining digit with len == 6
		for _, pattern := range display.patterns {
			if len(pattern) == 6 {
				if pattern != six && pattern != nine {
					zero = pattern
				}
			}
		}

		// now we can use 0 and 8 to find the middle letter
		for _, letter := range eight {
			if !strings.ContainsRune(zero, letter) {
				MIDDLE = byte(letter)
			}
		}

		// now that we have the middle, we can use 4 to find TOP_LEFT
		for _, letter := range four {
			if letter != rune(TOP_RIGHT) && letter != rune(BOT_RIGHT) && letter != rune(MIDDLE) {
				TOP_LEFT = byte(letter)
			}
		}

		// therfore BOT is the last remaining char
		for _, letter := range eight {
			if letter != rune(TOP) && letter != rune(TOP_RIGHT) && letter != rune(BOT_RIGHT) && letter != rune(BOT_LEFT) && letter != rune(TOP_LEFT) && letter != rune(MIDDLE) {
				BOT = byte(letter)
			}
		}

		// map errythang
		posMap := map[byte]string{
			TOP:       "TOP",
			TOP_RIGHT: "TOP_RIGHT",
			MIDDLE:    "MIDDLE",
			BOT_RIGHT: "BOT_RIGHT",
			BOT:       "BOT",
			BOT_LEFT:  "BOT_LEFT",
			TOP_LEFT:  "TOP_LEFT",
		}

		intMap := map[string]int{
			"TOP":       1,
			"TOP_RIGHT": 2,
			"MIDDLE":    3,
			"BOT_RIGHT": 4,
			"BOT":       5,
			"BOT_LEFT":  6,
			"TOP_LEFT":  7,
		}

		digitMap := map[int]string{
			25: "0",
			17: "2",
			15: "3",
			20: "5",
			26: "6",
			22: "9",
		}

		outStr := ""
		for _, out := range display.output {
			// first check for the definites -> 1, 4, 7, 8
			if len(out) == 2 {
				outStr += "1"
				continue
			}
			if len(out) == 4 {
				outStr += "4"
				continue
			}
			if len(out) == 3 {
				outStr += "7"
				continue
			}
			if len(out) == 7 {
				outStr += "8"
				continue
			}
			// now we do the mapping
			sum := 0
			for _, letter := range out {
				pos := posMap[byte(letter)]
				numVal := intMap[pos]
				sum += numVal
			}
			outStr += digitMap[sum]
		}

		outInt, err := strconv.Atoi(outStr)
		checkErr(err)
		total += outInt
	}

	return total
}

type Display struct {
	patterns []string
	output   []string
}

func readInput() []Display {
	file, err := os.Open("input.txt")
	checkErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var displays []Display
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		display := Display{
			patterns: strings.Fields(line[0]),
			output:   strings.Fields(line[1]),
		}
		displays = append(displays, display)
	}

	return displays
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
