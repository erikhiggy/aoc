package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	caloriesList := readInput()
	fmt.Println(part1(caloriesList))
	fmt.Println(part2(caloriesList))
}

func part1(caloriesList []string) int {
	sum := 0
	max := 0
	for _, c := range caloriesList {
		if c != "" {
			n, err := strconv.Atoi(c)
			if err != nil {
				fmt.Errorf("could not convert calorie string to int")
			}
			sum += n
		}

		if c == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}

	return max
}

func part2(caloriesList []string) int {
	sum := 0
	var caloriesSumList []int
	for _, c := range caloriesList {
		if c != "" {
			n, err := strconv.Atoi(c)
			if err != nil {
				fmt.Errorf("could not convert calorie string to int")
			}
			sum += n
		}

		if c == "" {
			caloriesSumList = append(caloriesSumList, sum)
			sum = 0
		}
	}

	sort.Ints(caloriesSumList)
	return caloriesSumList[len(caloriesSumList)-1] + caloriesSumList[len(caloriesSumList)-2] + caloriesSumList[len(caloriesSumList)-3]
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}