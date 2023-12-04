package day3

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

func hasValidNeighbor(lineI int, i int, start int, lines []string, lineL int) (rune, int, int) {
	// search neighbors
	var hasNeighbor rune
	var row int
	var col int
	// above
	if lineI > 0 {
		for x, neighbor := range lines[lineI-1][start:i] {
			if rune(neighbor) != '.' && !unicode.IsDigit(neighbor) {
				hasNeighbor = neighbor
				row = lineI - 1
				col = start + x
			}
		}
	}
	// below
	if lineI < len(lines)-1 {
		//if hasNeighbor == 0 && lineI < len(lines)-1 {
		for x, neighbor := range lines[lineI+1][start:i] {
			if rune(neighbor) != '.' && !unicode.IsDigit(neighbor) && hasNeighbor != '*' {
				hasNeighbor = neighbor
				row = lineI + 1
				col = start + x
			}
		}
	}
	// left
	if start > 0 &&
		rune(lines[lineI][start-1]) != '.' &&
		!unicode.IsDigit(rune(lines[lineI][start-1])) && hasNeighbor != '*' {

		hasNeighbor = rune(lines[lineI][start-1])
		row = lineI
		col = start - 1
	}
	// right
	if i < lineL &&
		rune(lines[lineI][i]) != '.' &&
		!unicode.IsDigit(rune(lines[lineI][i])) && hasNeighbor != '*' {

		hasNeighbor = rune(lines[lineI][i])
		row = lineI
		col = i
	}
	// diagonal top left
	if lineI > 0 &&
		start > 0 &&
		rune(lines[lineI-1][start-1]) != '.' &&
		!unicode.IsDigit(rune(lines[lineI-1][start-1])) && hasNeighbor != '*' {

		hasNeighbor = rune(lines[lineI-1][start-1])
		row = lineI - 1
		col = start - 1
	}
	// diagonal bottom left
	if lineI < len(lines)-1 &&
		start > 0 &&
		rune(lines[lineI+1][start-1]) != '.' &&
		!unicode.IsDigit(rune(lines[lineI+1][start-1])) && hasNeighbor != '*' {

		hasNeighbor = rune(lines[lineI+1][start-1])
		row = lineI + 1
		col = start - 1
	}
	// diagonal top right
	if lineI > 0 &&
		i < lineL &&
		rune(lines[lineI-1][i]) != '.' &&
		!unicode.IsDigit(rune(lines[lineI-1][i])) && hasNeighbor != '*' {

		hasNeighbor = rune(lines[lineI-1][i])
		row = lineI - 1
		col = i
	}
	// diagonal bottom right
	if hasNeighbor == 0 &&
		lineI < len(lines)-1 &&
		i < lineL &&
		rune(lines[lineI+1][i]) != '.' &&
		!unicode.IsDigit(rune(lines[lineI+1][i])) && hasNeighbor != '*' {

		hasNeighbor = rune(lines[lineI+1][i])
		row = lineI + 1
		col = i
	}

	//return hasNeighbor, row, col
	return hasNeighbor, row, col
}

func RunPart1() {
	// read file
	data, err := os.ReadFile("day3/input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	lineL := len(lines[0])

	ans := 0
	for lineI, line := range lines {
		i := 0
		start := -1
		for i < lineL {
			if unicode.IsDigit(rune(line[i])) {
				if start == -1 {
					start = i
				}
			} else {
				if start != -1 {
					neighbor, _, _ := hasValidNeighbor(lineI, i, start, lines, lineL)
					if neighbor != 0 {
						num, err := strconv.Atoi(line[start:i])
						check(err)
						ans += num
					}
				}
				start = -1
			}
			i += 1
		}
		if start != -1 {
			neighbor, _, _ := hasValidNeighbor(lineI, i, start, lines, lineL)
			if neighbor != 0 {
				num, err := strconv.Atoi(line[start:])
				check(err)
				ans += num
			}
		}
	}
	fmt.Println(ans)
}

func RunPart2() {
	// read file
	data, err := os.ReadFile("day3/input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")
	lineL := len(lines[0])

	gears := map[string]int{}
	gearCount := map[string]int{}

	ans := 0
	for lineI, line := range lines {
		i := 0
		start := -1
		for i < lineL {
			if unicode.IsDigit(rune(line[i])) {
				if start == -1 {
					start = i
				}
			} else {
				if start != -1 {
					neighbor, row, col := hasValidNeighbor(lineI, i, start, lines, lineL)
					if neighbor == '*' {
						if gears[strconv.Itoa(row)+"-"+strconv.Itoa(col)] == 0 {
							gears[strconv.Itoa(row)+"-"+strconv.Itoa(col)] = 1
						}
						num, err2 := strconv.Atoi(line[start:i])
						check(err2)
						gears[strconv.Itoa(row)+"-"+strconv.Itoa(col)] *= num
						gearCount[strconv.Itoa(row)+"-"+strconv.Itoa(col)] += 1
					}
				}
				start = -1
			}
			i += 1
		}
		if start != -1 {
			neighbor, row, col := hasValidNeighbor(lineI, i, start, lines, lineL)
			if neighbor == '*' {
				if gears[strconv.Itoa(row)+"-"+strconv.Itoa(col)] == 0 {
					gears[strconv.Itoa(row)+"-"+strconv.Itoa(col)] = 1
				}
				num, err2 := strconv.Atoi(line[start:i])
				check(err2)
				gears[strconv.Itoa(row)+"-"+strconv.Itoa(col)] *= num
				gearCount[strconv.Itoa(row)+"-"+strconv.Itoa(col)] += 1
			}
		}
	}
	c := 0
	for k, val := range gears {
		if gearCount[k] == 2 {
			ans += val
		}
		c += 1
	}
	fmt.Println(ans)
}
