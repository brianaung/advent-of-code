package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2023/day04/input.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	partone(lines)
	parttwo(lines)
}

func partone(lines []string) {
	sum := 0
	for _, line := range lines {
		a := strings.Split(strings.Split(line, ":")[1], "|")
		win, mine := strings.Split(a[0], " "), strings.Split(a[1], " ")
		winset := make(map[string]bool, len(win)-1)
		for _, w := range win {
			if len(w) > 0 {
				winset[w] = true
			}
		}
		n := 0
		for _, m := range mine {
			if len(m) > 0 && winset[m] {
				n++
			}
		}
		sum += int(math.Pow(2, float64(n-1)))
	}
	fmt.Println(sum)
}

func parttwo(lines []string) {
	matches := make(map[int]int)
	for i, line := range lines {
		a := strings.Split(strings.Split(line, ":")[1], "|")
		win, mine := strings.Split(a[0], " "), strings.Split(a[1], " ")
		winset := make(map[string]bool, len(win)-1)
		for _, w := range win {
			if len(w) > 0 {
				winset[w] = true
			}
		}
		numMatch := 0
		for _, m := range mine {
			if len(m) > 0 && winset[m] {
				numMatch++
			}
		}
		matches[i] = numMatch
	}

	// to loop the matches map in order
	keys := make([]int, 0)
	for k := range matches {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// init instances map
	instances := make(map[int]int)
	for i := 0; i < len(keys); i++ {
		instances[i] = 1
	}
	for _, k := range keys {
		count := 0
		// update instances for the next "matches[k] (i.e. no. instances of curr card)" cards
		for count < matches[k] {
			instances[k+count+1] += instances[k]
			count++
		}
	}
	sum := 0
	for _, v := range instances {
		sum += v
	}
	fmt.Println(sum)
}
