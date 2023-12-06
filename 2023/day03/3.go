package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("2023/day03/input.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	fmt.Println(partone(lines))
	fmt.Println(parttwo(lines))
}

func partone(grids []string) int {
	sum := 0
	for r, grid := range grids {
		n := 0
		ispart := false
		for c, ch := range grid {
			if unicode.IsDigit(ch) {
				// get part number (so far)
				a, _ := strconv.Atoi(string(ch))
				n = n*10 + a
				// check adjacents
				for _, ar := range []int{-1, 0, 1} {
					for _, ac := range []int{-1, 0, 1} {
						if r+ar >= 0 && r+ar < len(grids) && c+ac >= 0 && c+ac < len(grid) {
							if adj := grids[r+ar][c+ac]; adj != '.' && !unicode.IsDigit(rune(adj)) {
								ispart = true
							}
						}
					}
				}
			}
			// if curr ch is not a digit (or) we are at the last pos -> we got the full number
			if !unicode.IsDigit(ch) || c == len(grid)-1 {
				if ispart {
					sum += n
				}
				n = 0
				ispart = false
			}
		}
	}
	return sum
}

func parttwo(grids []string) int {
	gearNumMap := make(map[[2]int][]int)
	for r, grid := range grids {
		gear, n := [2]int{-1, -1}, 0
		for c, ch := range grid {
			if unicode.IsDigit(ch) {
				// get part number (so far)
				a, _ := strconv.Atoi(string(ch))
				n = n*10 + a
				// check adjacents and remember gear informations
				for _, ar := range []int{-1, 0, 1} {
					for _, ac := range []int{-1, 0, 1} {
						if r+ar >= 0 && r+ar < len(grids) && c+ac >= 0 && c+ac < len(grid) {
							if adj := grids[r+ar][c+ac]; adj == '*' {
								gear = [2]int{r + ar, c + ac}
							}
						}
					}
				}
			}
			// update map with gears and part numbers
			if (!unicode.IsDigit(ch) || c == len(grid)-1) && n > 0 {
				if gear[0] >= 0 && gear[1] >= 0 {
					gearNumMap[gear] = append(gearNumMap[gear], n)
				}
				n = 0
				gear = [2]int{-1, -1}
			}
		}
	}
	sum := 0
	for _, v := range gearNumMap {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}
