package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type pos struct {
	x int
	y int
}

func main() {
	part := flag.Int("part", 1, "part 1 or 2")
	flag.Parse()

	body, err := os.ReadFile("2015/day03/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	if *part == 1 {
		fmt.Println(part1(string(body)))
	} else if *part == 2 {
		fmt.Println(part2(string(body)))
	}
}

func part1(input string) int {
	visited := make(map[pos]int)
	curr := pos{0, 0}
	visited[curr] += 1
	for _, c := range input {
		deliverPresents(c, visited, &curr)
	}

	return len(visited)
}

func part2(input string) int {
	visited := make(map[pos]int)
	santaCurr, roboCurr := pos{0, 0}, pos{0, 0}
	visited[santaCurr] += 1
	for i, c := range input {
		if i%2 == 0 {
			deliverPresents(c, visited, &santaCurr)
		} else {
			deliverPresents(c, visited, &roboCurr)
		}
	}
	return len(visited)
}

func deliverPresents(c rune, visited map[pos]int, curr *pos) {
	switch c {
	case '^': // north
		curr.y += 1
		visited[pos{curr.x, curr.y}] += 1
	case '>': // east
		curr.x += 1
		visited[pos{curr.x, curr.y}] += 1
	case 'v': // south
		curr.y -= 1
		visited[pos{curr.x, curr.y}] += 1
	case '<': // west
		curr.x -= 1
		visited[pos{curr.x, curr.y}] += 1
	}
}
