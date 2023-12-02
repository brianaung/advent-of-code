package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("2023/day01/input.txt")
	lines := strings.Split(string(data), "\n")

	fmt.Printf("Part One: %v\n", Trebuchet(lines, 1))
	fmt.Printf("Part Two: %v\n", Trebuchet(lines, 2))
}

func Trebuchet(lines []string, part int) int {
	sum := 0
	for _, line := range lines {
		digits := []string{}
		for i, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, string(c))
			}

			//====================== part 2 ======================
			if part == 2 {
				for j, val := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
					if strings.HasPrefix(line[i:], val) {
						digits = append(digits, strconv.Itoa(j+1))
					}
				}
			}
			//===================================================

		}
		if len(digits) > 0 {
			d, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
			sum += d
		}
	}
	return sum
}
