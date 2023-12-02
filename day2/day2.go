package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunPart1() {
	bagCount := map[string]int{"red": 12, "green": 13, "blue": 14}

	// read file
	data, err := os.ReadFile("day2/input.txt")
	check(err)

	lines := string(data)
	sum := 0
	for _, line := range strings.Split(lines, "\n") {
		isValidGame := true

		// parse game id
		parseGameId := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.Split(parseGameId[0], " ")[1])
		check(err)

		// parse hands
		hands := strings.Split(parseGameId[1], "; ")
		for _, hand := range hands {
			counts := strings.Split(hand, ", ")
			for _, count := range counts {
				parseCount := strings.Split(count, " ")
				num, err := strconv.Atoi(parseCount[0])
				check(err)
				color := parseCount[1]
				if bagCount[color] < num {
					isValidGame = false
					break
				}
			}
			if !isValidGame {
				break
			}
		}
		if isValidGame {
			sum = sum + gameId
		}
	}
	fmt.Println(sum)
}

func RunPart2() {

	// read file
	data, err := os.ReadFile("day2/input.txt")
	check(err)

	lines := string(data)
	sum := 0
	for _, line := range strings.Split(lines, "\n") {
		bagCount := map[string]int{"red": 0, "green": 0, "blue": 0}

		// parse game id
		parseGameId := strings.Split(line, ": ")
		check(err)

		// parse hands
		hands := strings.Split(parseGameId[1], "; ")
		for _, hand := range hands {
			counts := strings.Split(hand, ", ")
			for _, count := range counts {
				parseCount := strings.Split(count, " ")
				num, err := strconv.Atoi(parseCount[0])
				check(err)
				color := parseCount[1]
				bagCount[color] = max(bagCount[color], num)
			}
		}
		sum = sum + bagCount["red"]*bagCount["green"]*bagCount["blue"]
	}
	fmt.Println(sum)
}
