package main

import (
	"fmt"
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
		winning, mine := strings.Split(a[0], " "), strings.Split(a[1], " ")
		count, curr := 0, 0
		for _, m := range mine {
			for _, w := range winning {
				if len(m) > 0 && m == w {
					if count == 0 {
						curr += 1
					} else {
						curr *= 2
					}
					count++
					continue
				}
			}
		}
		sum += curr
	}
	fmt.Println(sum)
}

func parttwo(lines []string) {
	matches := make(map[int]int)
	for i, line := range lines {
		a := strings.Split(strings.Split(line, ":")[1], "|")
		winning, mine := strings.Split(a[0], " "), strings.Split(a[1], " ")
		count := 0
		for _, m := range mine {
			for _, w := range winning {
				if len(m) > 0 && m == w {
					count++
					continue
				}
			}
		}
		matches[i+1] = count
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
		instances[i+1] = 1
	}
	// for each card
	for _, k := range keys {
		count := 0
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
