package main

import (
	"fmt"
	"strconv"
	"strings"
)

const sampleInput = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func parseInput(input string) []string {
	var res []string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		res = append(res, line)
	}
	return res
}

func main() {
	in := parseInput(sampleInput)
	sum := 0
	sumWithConcat := 0
	for _, v := range in {
		parts := strings.Split(v, ":")
		testVal := convertInt(parts[0])
		ints := toInts(parts[1])
		// Part 1
		if checkValid(testVal, ints, false) {
			sum += testVal
		}
		// Part 2
		if checkValid(testVal, ints, true) {
			sumWithConcat += testVal
		}
	}
	fmt.Println("Pt1:", sum)
	fmt.Println("Pt2:", sumWithConcat)
}

func checkValid(testVal int, vals []int, withConcat bool) bool {
	if len(vals) == 1 {
		return vals[0] == testVal
	}
	last := vals[len(vals)-1]

	if (testVal % last) == 0 {
		if checkValid(testVal/last, vals[:len(vals)-1], withConcat) {
			return true
		}
	}

	if (testVal - last) > 0 {
		if checkValid(testVal-last, vals[:len(vals)-1], withConcat) {
			return true
		}
	}

	if withConcat {
		if testVal%10 == last%10 {
			testValStr := strconv.Itoa(testVal)
			lastStr := strconv.Itoa(last)
			if strings.HasSuffix(testValStr, lastStr) {
				vStr := strings.TrimSuffix(testValStr, lastStr)
				v := 0
				if vStr != "" {
					v = convertInt(vStr)
				}
				if checkValid(v, vals[:len(vals)-1], withConcat) {
					return true
				}
			}
		}
	}

	return false
}

func toInts(s string) []int {
	var ret []int
	for _, p := range strings.Fields(s) {
		ret = append(ret, convertInt(p))
	}
	return ret
}

func convertInt(s string) int {
	if v, err := strconv.Atoi(s); err != nil {
		panic(fmt.Sprint(s, err))
	} else {
		return v
	}
}
