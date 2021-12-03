package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bNumbers := readInput()
	// fmt.Println(part1(bNumbers))
	fmt.Println(part2(bNumbers))
}

func part1(bNumbers []string) int {
	gammaRate := ""

	for i := 0; i < len(bNumbers[0]); i++ {
		zeroCounter := 0
		oneCounter := 0
		for _, bin := range bNumbers {
			if bin[i] == '0' {
				zeroCounter++
			} else {
				oneCounter++
			}
		}
		bitToAppend := "1"
		if zeroCounter > oneCounter {
			bitToAppend = "0"
		}
		gammaRate += bitToAppend
	}

	// get the epsilon rate - the compliment of the gamma rate
	epsilonRate := ""
	for _, bit := range gammaRate {
		if bit == '0' {
			epsilonRate += "1"
		} else {
			epsilonRate += "0"
		}
	}

	gammaDec := stringToBin(gammaRate)
	// fmt.Printf("gamma decimal: %d\n", gammaDec)
	epDec := stringToBin(epsilonRate)
	// fmt.Printf("epsilon decimal: %d\n", epDec)
	return epDec*gammaDec
}

func part2(bNumbers []string) int {
	zeroCounter := 0
	oneCounter := 0
	mostCommonBitArray := make([]int, len(bNumbers[0]))
	for i := 0; i < len(bNumbers[0]); i++ {
		zeroCounter = 0
		oneCounter = 0
		for _, bin := range bNumbers {
			if bin[i] == '0' {
				zeroCounter++
			} else {
				oneCounter++
			}
		}
		if zeroCounter > oneCounter {
			mostCommonBitArray[i] = 0
		} else {
			mostCommonBitArray[i] = 1
		}
	}

	fmt.Println(mostCommonBitArray)
	return 0
}

func stringToBin(s string) int {
	decimalRep := 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '1' {
			decimalRep += 1<<(len(s)-1-i)
		}
	}
	return decimalRep
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