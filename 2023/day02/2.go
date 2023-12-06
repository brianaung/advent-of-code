package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("2023/day02/input.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	fmt.Println(partone(lines))
	fmt.Println(parttwo(lines))
}

func partone(games []string) int {
	sum := 0
	for _, game := range games {
		gamedata := strings.Split(game, ": ")
		id := gamedata[0]
		valid := true
		for _, set := range strings.Split(gamedata[1], "; ") {
			for _, data := range strings.Split(set, ", ") {
				pair := strings.Split(data, " ")
				color := pair[1]
				count, _ := strconv.Atoi(pair[0])
				if count > map[string]int{"red": 12, "green": 13, "blue": 14}[color] {
					valid = false
				}
			}
		}
		if valid {
			id, _ := strconv.Atoi(strings.Split(id, " ")[1])
			sum += id
		}
	}
	return sum
}

func parttwo(games []string) int {
	sum := 0
	for _, game := range games {
		gamedata := strings.Split(game, ": ")
		pairmap := make(map[string]int)
		for _, set := range strings.Split(gamedata[1], "; ") {
			for _, data := range strings.Split(set, ", ") {
				pair := strings.Split(data, " ")
				count, _ := strconv.Atoi(pair[0])
				if v, ok := pairmap[pair[1]]; !ok || count > v {
					pairmap[pair[1]] = count
				}
			}
		}
		power := 1
		for _, v := range pairmap {
			power *= v
		}
		sum += power
	}
	return sum
}
