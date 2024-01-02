package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePtr := flag.String("input", "", "input file to read")
	partPtr := flag.Int("part", 1, "part 1 or 2")
	flag.Parse()
	if *filePtr == "" {
		log.Fatal("Please input file to read using -input flag.")
	}

	body, err := os.ReadFile(*filePtr)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(NotQuiteLisp(body, *partPtr))
}

func NotQuiteLisp(input []byte, part int) int {
	floor := 0
	for i, c := range input {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		}

		if part == 2 && floor == -1 {
			return i + 1
		}
	}

	return floor
}
