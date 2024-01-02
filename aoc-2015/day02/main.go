package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/brianaung/advent-of-code/meth"
)

type dimension struct {
	l int
	w int
	h int
}

func main() {
	fmt.Println("Part1:", part1())
	fmt.Println("Part2:", part2())
}

func part1() int {
	total := 0
	for _, str := range parseInputFile() {
		d, err := parseDimensions(str)
		if err != nil {
			continue
		}
		l := d.l
		w := d.w
		h := d.h
		total += 2*l*w + 2*w*h + 2*h*l +
			meth.MinInt(l*w, l*h, w*h)
	}
	return total
}

func part2() int {
	total := 0
	for _, str := range parseInputFile() {
		d, err := parseDimensions(str)
		if err != nil {
			continue
		}
		sides := []int{d.l, d.w, d.h}
		sort.Ints(sides)
		total += 2*sides[0] + 2*sides[1] + (d.l * d.w * d.h)
	}
	return total
}

func parseDimensions(s string) (*dimension, error) {
	r, _ := regexp.Compile(`(\d+)x(\d+)x(\d+)`)
	matches := r.FindStringSubmatch(s)

	if len(matches) != 4 {
		return nil, errors.New("error matching substring")
	}

	l, _ := strconv.Atoi(matches[1])
	w, _ := strconv.Atoi(matches[2])
	h, _ := strconv.Atoi(matches[3])

	return &dimension{l, w, h}, nil
}

func parseInputFile() []string {
	body, err := os.ReadFile("2015/day02/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	return strings.Split(string(body), "\n")
}
