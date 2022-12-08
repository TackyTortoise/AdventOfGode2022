package main

import (
	"adventofgode-2022/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func parsePair(in string) (int, int) {
	split := strings.Split(in, "-")
	one, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	two, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return one, two
}

func checkFullOverlap(l1, h1, l2, h2 int) bool {
	return (l1 >= l2 && h1 <= h2) || (l2 >= l1 && h2 <= h1)
}

func checkOverlap(l1, h1, l2, h2 int) bool {
	if h1 < h2 {
		return h1 >= l2
	}
	return h2 >= l1
}

func Day4Part1() {
	lines := helpers.ReadFileLines("../input/04")
	overlapCount := 0
	for _, l := range lines {
		split := strings.Split(l, ",")
		start1, end1 := parsePair(split[0])
		start2, end2 := parsePair(split[1])
		if checkFullOverlap(start1, end1, start2, end2) {
			overlapCount++
		}
	}
	fmt.Println("Overlapping elfs: ", overlapCount)
}

func Day4Part2() {
	lines := helpers.ReadFileLines("../input/04")
	overlapCount := 0
	for _, l := range lines {
		split := strings.Split(l, ",")
		start1, end1 := parsePair(split[0])
		start2, end2 := parsePair(split[1])
		if checkOverlap(start1, end1, start2, end2) {
			overlapCount++
		}
	}
	fmt.Println("Overlapping elfs: ", overlapCount)
}
