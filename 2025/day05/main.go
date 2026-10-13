package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return lines
}

func isInRange(n, start, end int) bool {
	return n >= start && n <= end
}

func part1(ranges []string, ids []int) int {
	freshSet := make(map[int]struct{})
	for _, id := range ids {
		for _, r := range ranges {
			rangeSplit := strings.Split(r, "-")
			start := mustInt(rangeSplit[0])
			end := mustInt(rangeSplit[1])
			if isInRange(id, start, end) {
				_, exists := freshSet[id]
				if exists {
					continue
				} else {
					freshSet[id] = struct{}{}
				}
			}
		}
	}

	return len(freshSet)
}

func part2(ranges []string) int {
	slices.SortFunc(ranges, func(a, b string) int {
		aAsInt := mustInt(strings.Split(a, "-")[0])
		bAsInt := mustInt(strings.Split(b, "-")[0])

		return cmp.Compare(aAsInt, bAsInt)
	})

	fresh := 0
	for i, r := range ranges {
		split := strings.Split(r, "-")
		start := mustInt(split[0])
		end := mustInt(split[1])

		last := 0
		if i > 0 {
			prevSplit := strings.Split(ranges[i-1], "-")
			last = mustInt(prevSplit[1])
		}

		tmp := start
		offset := 0
		if last-tmp >= 0 {
			offset = (last - tmp) + 1
		}

		newStart := tmp + offset

		diff := 0
		if newStart <= end {
			diff = (end - newStart) + 1
		}
		fmt.Println(r)
		// fmt.Println("diff", diff)
		fresh += diff
	}

	/*

		3-8    => 6
		7-9    => 1
		11-20  => 10
		12-15  => 0
		22-33  => 12

		total  => 29

	*/

	return fresh
}

func main() {
	in := parseInput("input.txt")

	// var ranges []string
	// var ids []int
	// for _, l := range in {
	// 	if l == "" {
	// 		continue
	// 	}
	// 	if strings.Contains(l, "-") {
	// 		ranges = append(ranges, l)
	// 	} else {
	// 		asNum := mustInt(l)
	// 		ids = append(ids, asNum)
	// 	}

	// }

	// part1Sum := part1(ranges, ids)
	part2Sum := part2(in)
	// fmt.Println("part 1: ", part1Sum)
	fmt.Println("part 2: ", part2Sum)
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
