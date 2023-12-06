package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	lines := strings.Split(string(data), "\n")
	partone(lines)
	parttwo(lines)
}

func partone(lines []string) {
	times, dis := strings.Fields(lines[0])[1:], strings.Fields(lines[1])[1:]
	p1 := 1
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(string(times[i]))
		d, _ := strconv.Atoi(string(dis[i]))
		ways := solve(t, d)
		p1 *= ways
	}
	fmt.Println(p1)
}

func parttwo(lines []string) {
	t, _ := strconv.Atoi(strings.Replace(strings.Split(lines[0], ":")[1], " ", "", -1))
	d, _ := strconv.Atoi(strings.Replace(strings.Split(lines[1], ":")[1], " ", "", -1))
	fmt.Println(solve(t, d))
}

func solve(t int, d int) int {
	i, ret := 0, 0
	for i < t {
		if (i * (t - i)) > d {
			ret++
		}
		i++
	}
	return ret
}
