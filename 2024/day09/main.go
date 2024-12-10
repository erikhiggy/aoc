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
		//fmt.Println("Start:", start, "End:", end)
		if input[start] != "." {
			start++
		} else if input[end] == "." {
			end--
		} else {
			//fmt.Println("Swapping", start, end)
			swapIndexes(input, start, end)
			//fmt.Println(input)
			start++
			end--
		}
	}
}

func swapBulkWithFreeSpace(input []string) {
	start, end := 0, len(input)-1
	for start < end {
		availableSpace := 0
		neededSpace := 0

		if input[start] != "." {
			start++
		}
		if input[end] == "." {
			end--
		}
		if input[start] == "." && input[end] != "." {
			// count available space
			ptr := start
			for input[ptr] == "." {
				availableSpace++
				ptr++
			}

			// count needed space
			ptr = end
			currVal := input[ptr]
			for input[ptr] == currVal {
				neededSpace++
				ptr--
			}

			if availableSpace >= neededSpace {
				fmt.Println("Swapping", start, end)
				for i := 0; i < neededSpace; i++ {
					swapIndexes(input, start+i, end-i)
				}
				fmt.Println(input)
				start += neededSpace
				end -= neededSpace
			} else {
				start += availableSpace
				end -= availableSpace
			}
		}
	}
}

func swapIndexes(input []string, i, j int) {
	input[i], input[j] = input[j], input[i]
}

func main() {
	parsed := parseInput(sample)

	// Part 1
	diskMap := makeDiskMap(parsed)
	//fmt.Println(diskMap)
	swapRightMostWithFreeSpace(diskMap)
	//fmt.Println(diskMap)
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
	fmt.Println(diskMap2)
	swapBulkWithFreeSpace(diskMap2)
	fmt.Println(diskMap2)
}
