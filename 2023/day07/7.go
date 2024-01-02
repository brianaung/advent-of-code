package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])

	handmap := make(map[string]int)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		fields := strings.Fields(line)
		hand := string(fields[0])
		bid, _ := strconv.Atoi(string(fields[1]))
		handmap[hand] = bid
	}

	hands := make([]string, 0)
	for k := range handmap {
		hands = append(hands, k)
	}
	// sort hands based on their strength
	sort.Slice(hands, func(i, j int) bool {
		h1 := strings.Split(hands[i], "")
		h2 := strings.Split(hands[j], "")
		if strength(h1) < strength(h2) { // first ordering rule
			return true
		} else if strength(h1) == strength(h2) { // second ordering rule
			return replacecard(h1) < replacecard(h2)
		}
		return false
	})

	total := 0
	for i, h := range hands {
		total += handmap[h] * (i + 1)
	}
	fmt.Println(total)
}

func replacecard(hand []string) string {
	ret := ""
	for _, h := range hand {
		switch h {
		case "A":
			ret += "E"
		case "K":
			ret += "D"
		case "Q":
			ret += "C"
		case "J":
			// ret += "B" // part 1
			ret += "1" // part 2 (J is the weakest)
		case "T":
			ret += "A"
		default:
			ret += h
		}
	}
	return ret
}

func strength(hand []string) int {
	counter := count(hand)
	switch {
	case reflect.DeepEqual(counter, []int{5}): // five of a kind
		return 7
	case reflect.DeepEqual(counter, []int{1, 4}): // four of a kind
		return 6
	case reflect.DeepEqual(counter, []int{2, 3}): // full house
		return 5
	case reflect.DeepEqual(counter, []int{1, 1, 3}): // three of a kind
		return 4
	case reflect.DeepEqual(counter, []int{1, 2, 2}): // two pair
		return 3
	case reflect.DeepEqual(counter, []int{1, 1, 1, 2}): // one pair
		return 2
	case reflect.DeepEqual(counter, []int{1, 1, 1, 1, 1}): // one pair
		return 1
	default:
		return 0
	}
}

// part 2
func count(slice []string) (counter []int) {
	counter = make([]int, 0)
	cmap := make(map[string]int)
	jcount := 0
	for i, s1 := range slice {
		if s1 == "J" {
			jcount++
			continue
		}
		count := 0
		for _, s2 := range slice[i:] {
			if s1 == s2 {
				count++
			}
		}
		if _, ok := cmap[s1]; !ok {
			cmap[s1] = count
		}
	}
	for _, v := range cmap {
		counter = append(counter, v)
	}

	if len(counter) > 0 {
		hi := counter[0]
		j := 0
		for i, c := range counter {
			if c > hi {
				hi = c
				j = i
			}
		}
		counter[j] += jcount
	} else {
		counter = append(counter, jcount)
	}
	sort.Ints(counter)

	return
}

// part 1
// func count[T comparable](slice []T) (counter []int) {
// 	counter = make([]int, 0)
// 	cmap := make(map[T]int)
// 	for i, s1 := range slice {
// 		count := 0
// 		for _, s2 := range slice[i:] {
// 			if s1 == s2 {
// 				count++
// 			}
// 		}
// 		if _, ok := cmap[s1]; !ok {
// 			cmap[s1] = count
// 		}
// 	}
// 	for _, v := range cmap {
// 		counter = append(counter, v)
// 	}
// 	sort.Ints(counter)
// 	return
// }
