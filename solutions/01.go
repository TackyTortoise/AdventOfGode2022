package main

import (
	"adventofgode-2022/helpers"
	"fmt"
	"log"
	"strconv"
)

func Day1Part1() {
	lines := helpers.ReadFileLines("../input/01")
	currTotal := 0
	max := 0
	for _, l := range lines {
		if len(l) == 0 {
			if currTotal > max {
				max = currTotal
			}
			currTotal = 0
			continue
		}
		num, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		currTotal += num
	}
	fmt.Printf("Max calories: %v\n", max)
}

func checkMaxes(v int, maxes []int) {
	for i, m := range maxes {
		if v > m {
			maxes[i] = v
			checkMaxes(m, maxes)
			break
		}
	}
}

func Day1Part2() {
	lines := helpers.ReadFileLines("../input/01")
	currTotal := 0
	maxes := make([]int, 3)

	for _, l := range lines {
		if len(l) == 0 {
			checkMaxes(currTotal, maxes)
			currTotal = 0
			continue
		}
		num, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		currTotal += num
	}

	sum := 0
	for _, m := range maxes {
		sum += m
	}
	fmt.Printf("Total calories: %v\n", sum)
}
