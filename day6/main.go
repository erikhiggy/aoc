package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := readInput()

	remaining := make([]int, 300)
	for _, n := range nums {
		remaining[n]++
	}

	sol(nums, remaining)
}

func sol(nums []int, rem []int) {
	// get the initial number of fish
	numFish := len(nums)
	for i := 0; i < 256; i++ {
		// sort so 0s are in the front
		if i == 80 {
			fmt.Printf("part1: %d\n", numFish)
		}

		numFish += rem[i]
		rem[i+7] += rem[i]
		rem[i+9] += rem[i]
	}
	fmt.Printf("part2: %d\n", numFish)
}

func readInput() []int {
	file, err := os.Open("input.txt")
	checkErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	var nums []int
	for _, n := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(n)
		checkErr(err)
		nums = append(nums, num)
	}

	return nums
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}