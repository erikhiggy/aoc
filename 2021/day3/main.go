package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bNumbers := readInput()
	fmt.Println(part1(bNumbers))
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
	// keep the bNumbers with the matching most common bit
	// in the corresponding position
	oxyRating := getOxygenGenRating(bNumbers)
	CO2Rating := getCO2Rating(bNumbers)

	oxyInt := stringToBin(oxyRating[0])
	CO2Int := stringToBin(CO2Rating[0])
	return oxyInt*CO2Int
}

func getOxygenGenRating(arr []string) []string {
	bitLen := len(arr[0])
	vsf := arr
	for i := 0; i < bitLen; i++ {
		importantBit := getImportantBit(i, vsf, true)
		if len(vsf) == 1 { break }
		vsf = Filter(vsf, func(v string) bool {
			return rune(v[i]) == importantBit
		})
	}
	return vsf
}

func getCO2Rating(arr []string) []string {
	bitLen := len(arr[0])
	vsf := arr
	for i := 0; i < bitLen; i++ {
		importantBit := getImportantBit(i, vsf, false)
		if len(vsf) == 1 { break }
		vsf = Filter(vsf, func(v string) bool {
			return rune(v[i]) == importantBit
		})
	}
	return vsf
}

func getImportantBit(i int, arr []string, mostFrequent bool) rune {
	zeroCounter := 0
	oneCounter := 0
	for _, bin := range arr {
		if bin[i] == '0' {
			zeroCounter++
		} else {
			oneCounter++
		}
	}

	mostImportantBit := '1'
	if mostFrequent {
		if zeroCounter > oneCounter {
			mostImportantBit = '0'
		}
	} else {
		if zeroCounter <= oneCounter {
			mostImportantBit = '0'
		}
	}

	return mostImportantBit
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

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}