package main

import (
	"fmt"
	"strconv"
	"strings"
)

const sample = `2333133121414131402`

func parseInput(in string) []string {
	return strings.Split(strings.TrimSpace(in), "")
}

func makeDiskMap(input []string) []string {
	// The even indexes of the sample input represent the size of the file block
	// and the odd indexes represent the free space available.
	var result []string
	for i := 0; i < len(input); i++ {
		fieldAsInt, err := strconv.Atoi(input[i])
		// If the index is even, we know it's a file block
		if i%2 == 0 {
			indexAsString := "0"
			if i > 0 {
				indexAsString = strconv.Itoa(i / 2)
			}
			if err != nil {
				panic(err)
			}
			for j := 0; j < fieldAsInt; j++ {
				result = append(result, indexAsString)
			}
		} else {
			// If the index is odd, we know it's free space
			for j := 0; j < fieldAsInt; j++ {
				result = append(result, ".")
			}
		}
	}
	return result
}

func swapRightMostWithFreeSpace(input []string) {
	start, end := 0, len(input)-1
	for start < end {
		if input[start] != "." {
			start++
		} else if input[end] == "." {
			end--
		} else {
			swapIndexes(input, start, end)
			start++
			end--
		}
	}
}

func swapBulkWithFreeSpace(input []string) {
	end := len(input) - 1
	for end >= 0 {
		start := 0
		available, needed := 0, 0
		if input[end] == "." && end > 0 {
			end--
		} else {
			ptr := end
			currVal := input[ptr]
			for ptr >= 0 && input[ptr] == currVal {
				needed++
				ptr--
			}

			for available < needed && start < end {
				// only count contiguous "."s as available space
				if input[start] == "." {
					available++
				} else {
					available = 0
				}
				start++
			}
			// we now have enough free space to swap
			if available >= needed {
				for i := 0; i < needed; i++ {
					swapIndexes(input, end-i, start-i-1)
				}
			}
		}
		end -= needed
	}
}

func swapIndexes(input []string, i, j int) {
	input[i], input[j] = input[j], input[i]
}

func main() {
	parsed := parseInput(sample)

	// Part 1
	diskMap := makeDiskMap(parsed)
	swapRightMostWithFreeSpace(diskMap)
	sum := 0
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == "." {
			continue
		}
		asInt, err := strconv.Atoi(diskMap[i])
		if err != nil {
			panic(err)
		}
		sum += i * asInt
	}
	fmt.Println("Part 1:", sum)

	// Part 2
	diskMap2 := makeDiskMap(parsed)
	swapBulkWithFreeSpace(diskMap2)
	sum2 := 0
	for i := 0; i < len(diskMap2); i++ {
		if diskMap2[i] == "." {
			continue
		}
		asInt, err := strconv.Atoi(diskMap2[i])
		if err != nil {
			panic(err)
		}
		sum2 += i * asInt
	}
	fmt.Println("Part 2:", sum2)
}
