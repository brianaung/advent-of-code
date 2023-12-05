package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/brianaung/advent-of-code/meth"
)

type mapinfo struct {
	dst  int
	src  int
	size int
}

type rng struct {
	st int
	ed int
}

func main() {
	data, _ := os.ReadFile(os.Args[1])
	G := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	// get seeds
	RS := strings.Split(strings.TrimSpace(strings.Split(G[0], ":")[1]), " ")
	// for part1
	seeds := make([]int, 0)
	for _, rs := range RS {
		s, _ := strconv.Atoi(rs)
		seeds = append(seeds, s)
	}
	// for part2
	pairs := make([]rng, 0)
	for i := 0; i < len(RS); i += 2 {
		src, _ := strconv.Atoi(RS[i])
		size, _ := strconv.Atoi(RS[i+1])
		pairs = append(pairs, rng{src, src + size})
	}
	// for each group/block excluding seeds
	for _, ms := range G[1:] {
		// get current block mapping infos
		M := make([]mapinfo, 0)
		for _, m := range strings.Split(ms, "\n")[1:] {
			N := strings.Split(m, " ")
			dst, _ := strconv.Atoi(N[0])
			src, _ := strconv.Atoi(N[1])
			size, _ := strconv.Atoi(N[2])
			M = append(M, mapinfo{dst, src, size})
		}
		// part1
		NS := make([]int, 0)
		for _, s := range seeds {
			v := s
			for _, m := range M {
				if v >= m.src && v < m.src+m.size {
					v = m.dst + v - m.src
					break
				}
			}
			NS = append(NS, v)
		}
		seeds = NS
		// part2
		inter := make([]rng, 0)
		for len(pairs) > 0 {
			st, ed := pairs[0].st, pairs[0].ed
			pairs = pairs[1:]
			matches := 0
			for _, m := range M {
				// checking intersection
				s := meth.Max(m.src, st)
				e := meth.Min(m.src+m.size, ed)
				if s < e {
					matches++
					inter = append(inter, rng{m.dst + s - m.src, m.dst + e - m.src})
					if s > st {
						pairs = append(pairs, rng{st, s})
					}
					if e < ed {
						pairs = append(pairs, rng{e, ed})
					}
					break
				}
			}
			if matches == 0 {
				inter = append(inter, rng{st, ed})
			}
		}
		pairs = inter
	}
	p2 := math.MaxInt
	for _, p := range pairs {
		if p.st < p2 {
			p2 = p.st
		}
	}

	fmt.Println("part1:", meth.Min(seeds...))
	fmt.Println("part2:", p2)
}

// for checking intersection:
// for each of the mapping, check if it overlaps with the seed range
//      seed range :    s1[       ]e1
//      src range  :        s2[        ]e2
//      intersect = {s: max(s1, s2), e: min(e1, e2)}, if s < e (e[      ]s makes no sense)
// still need to consider rest of seed range, to see if it intersect with next matches
//      seed range :    s1[       ]e1
//      src range  :        s2[        ]e2
//                        [xxx] push this to seed range
//   (OR)
//      seed range :             s1[       ]e1
//      src range  :        s2[        ]e2
//                                     [xxx] push this to seed range
//   (OR)
//      seed range :      s1[                ]e1
//      src range  :            s2[     ]e2
//                          [xxx]       [xxx] push these to seed range
