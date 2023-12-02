package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunPart1() {
	// read file
	data, err := os.ReadFile("day1/input.txt")
	check(err)

	lines := string(data)
	sum := 0
	for _, line := range strings.Split(lines, "\n") {
		var firstVal rune
		var lastVal rune
		// find the first value
		for i := 0; i < len(line); i++ {
			isFirstDigit := unicode.IsDigit(rune(line[i]))
			if isFirstDigit {
				firstVal = rune(line[i])
				break
			}
		}
		// find the last value
		for i := len(line) - 1; i >= 0; i-- {
			isLastDigit := unicode.IsDigit(rune(line[i]))
			if isLastDigit {
				lastVal = rune(line[i])
				break
			}
		}
		val, err := strconv.Atoi(string(firstVal) + string(lastVal))
		check(err)
		sum = sum + val
	}
	fmt.Println(sum)
}

func RunPart2() {
	// read file
	data, err := os.ReadFile("day1/input.txt")
	check(err)

	textNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	textToRune := map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}
	lines := string(data)
	sum := 0
	for _, line := range strings.Split(lines, "\n") {
		var firstVal rune
		var lastVal rune
		// find the first value
		for i := 0; i < len(line); i++ {
			isFirstDigit := unicode.IsDigit(rune(line[i]))
			if isFirstDigit {
				firstVal = rune(line[i])
				break
			}
			// handle text numbers
			for _, val := range textNumbers {
				if strings.HasPrefix(line[i:], val) {
					firstVal = textToRune[val]
				}
			}
			if firstVal != 0 {
				break
			}
		}
		// find the last value
		for i := len(line) - 1; i >= 0; i-- {
			isLastDigit := unicode.IsDigit(rune(line[i]))
			if isLastDigit {
				lastVal = rune(line[i])
				break
			}
			// handle text numbers
			for _, val := range textNumbers {
				if strings.HasPrefix(line[i:], val) {
					lastVal = textToRune[val]
				}
			}
			if lastVal != 0 {
				break
			}
		}
		val, err := strconv.Atoi(string(firstVal) + string(lastVal))
		check(err)
		sum = sum + val
	}
	fmt.Println(sum)
}
