package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum1, sum2 := 0, 0

	readFile, err := os.Open("data/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		l := fileScanner.Text()

		sum1 += lineToNumberCalibration1(l)
		sum2 += lineToNumberCalibration2(l)
	}

	fmt.Printf("Sum of calibration values (1): %d\n", sum1)
	fmt.Printf("Sum of calibration values (2): %d\n", sum2)
}

func lineToNumberCalibration1(l string) int {
	first := 0
	last := 0

	for _, c := range []rune(l) {
		if c >= '0' && c <= '9' {
			first = int(c) - '0'
			break
		}
	}

	for _, c := range reverse([]rune(l)) {
		if c >= '0' && c <= '9' {
			last = int(c) - '0'
			break
		}
	}

	return first*10 + last
}

func lineToNumberCalibration2(l string) int {
	first := 0
	strNum := []rune{}
	last := 0

	for _, c := range []rune(l) {
		if c >= '0' && c <= '9' {
			first = int(c) - '0'
			break
		} else {
			strNum = append(strNum, c)
			val := stringToNumber(string(strNum))
			if val >= 0 {
				first = val
				break
			}
		}
	}

	strNum = []rune{}

	for _, c := range reverse([]rune(l)) {
		if c >= '0' && c <= '9' {
			last = int(c) - '0'
			break
		} else {
			strNum = append([]rune{c}, strNum...)
			val := stringToNumber(string(strNum))
			if val >= 0 {
				last = val
				break
			}
		}
	}

	return first*10 + last
}

func reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func stringToNumber(s string) int {
	for start := 0; start < len(s); start++ {
		substr := s[start:]
		for end := len(substr); end > 0; end-- {
			switch substr[:end] {
			case "one":
				return 1
			case "two":
				return 2
			case "three":
				return 3
			case "four":
				return 4
			case "five":
				return 5
			case "six":
				return 6
			case "seven":
				return 7
			case "eight":
				return 8
			case "nine":
				return 9
			}
		}
	}

	return -1
}
