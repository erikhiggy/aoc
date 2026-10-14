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

type Range struct {
	start int
	end   int
}

func mergeRanges(ranges []string) int {
	// convert to int slice
	var ingredientRanges []Range
	for _, r := range ranges {
		split := strings.Split(r, "-")
		start := mustInt(split[0])
		end := mustInt(split[1])
		ingredientRanges = append(ingredientRanges, Range{start, end})
	}
	// sort based on first entry
	slices.SortFunc(ingredientRanges, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})

	var mergedRanges []Range

	for {
		// comparisons done since only have one more range
		if len(ingredientRanges) == 1 {
			mergedRanges = append(mergedRanges, ingredientRanges...)
			break
		}

		// pop the first two off to compare
		first := ingredientRanges[0]
		second := ingredientRanges[1]
		ingredientRanges = ingredientRanges[2:]

		// ranges do not intersect
		if second.start > first.end {
			mergedRanges = append(mergedRanges, first)
			ingredientRanges = append([]Range{second}, ingredientRanges...)
		} else {
			// ranges do intersect
			ingredientRanges = append([]Range{{first.start, max(first.end, second.end)}}, ingredientRanges...)
		}
	}

	fresh := 0
	for _, m := range mergedRanges {
		fresh += (m.end - m.start) + 1
	}
	return fresh
}

func main() {
	in := parseInput("input.txt")

	var ranges []string
	var ids []int
	for _, l := range in {
		if l == "" {
			continue
		}
		if strings.Contains(l, "-") {
			ranges = append(ranges, l)
		} else {
			asNum := mustInt(l)
			ids = append(ids, asNum)
		}

	}

	part1Sum := part1(ranges, ids)
	part2Sum := mergeRanges(ranges)
	fmt.Println("part 1: ", part1Sum)
	fmt.Println("part 2: ", part2Sum)
}

func mustInt(str string) int {
	asNum, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("mustInt: failed to parse %q: %v", str, err))
	}
	return asNum
}
