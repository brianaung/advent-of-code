package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile(os.Args[1])
	data := strings.Split(string(input), "\n\n")
	I, N := strings.Split(strings.TrimSpace(data[0]), ""), strings.Split(strings.TrimSpace(data[1]), "\n")

	imap := make(map[string][]string)
	for _, n := range N {
		replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
		ns := strings.Fields(replacer.Replace(n))
		imap[ns[0]] = []string{ns[1], ns[2]}
	}

	partone(I, imap)
	parttwo(I, imap)
}

func partone(I []string, imap map[string][]string) {
	p1, curr := 0, "AAA"
	for curr != "ZZZ" && len(I) > 0 {
		switch I[0] {
		case "L":
			curr = imap[curr][0]
		case "R":
			curr = imap[curr][1]
		}
		I = append(I[1:], I[0])
		p1++
	}
	fmt.Println(p1)
}

func parttwo(I []string, imap map[string][]string) {
	// curr is now a list of strings
	curr := make([]string, 0)
	for k := range imap {
		if strings.HasSuffix(k, "A") {
			curr = append(curr, k)
		}
	}

	ss := make([]int, 0)
	for _, c := range curr {
		s := 0
		tmp := I
		for !strings.HasSuffix(c, "Z") && len(tmp) > 0 {
			switch tmp[0] {
			case "L":
				c = imap[c][0]
			case "R":
				c = imap[c][1]
			}
			tmp = append(tmp[1:], tmp[0])
			s++
		}
		ss = append(ss, s)
	}
	fmt.Println(ss)
}
