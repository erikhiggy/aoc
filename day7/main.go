package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	positions := readInput()
	x := math.MaxInt64
	for i := 0; i < len(positions); i++ {
		diff := computeDiff(positions, i)
		if diff < x {
			x = diff
		}
	}

	fmt.Printf("part2: %d\n", x)

}

func computeDiff(pos []int, key int) int {
	sum := 0
	for _, p := range pos {
		sum += arithmetic(1, int(math.Abs(float64(p)-float64(key))), int(math.Abs(float64(p)-float64(key))))
	}
	return sum
}

func arithmetic(a1, an, n int) int {
	sum := (n*(a1 + an))/2
	return sum
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