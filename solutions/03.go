package main

import (
	"adventofgode-2022/helpers"
	"fmt"
	"unicode"
)

// Find character that appears in both strings
func checkMatch(one, two string) rune {
	encountered := make(map[rune]struct{}, len(one))
	for _, b := range one {
		encountered[b] = struct{}{}
	}

	for _, b := range two {
		if _, ok := encountered[b]; ok {
			return b
		}
	}

	panic("Did not find match.")
}

// Calculate score for matching character
func getPrio(match rune) int {
	if unicode.IsUpper(match) {
		return int(match-'A') + 27
	} else {
		return int(match-'a') + 1
	}
}

func Day3Part1() {
	rucksacks := helpers.ReadFileLines("../input/03")
	total := 0
	for _, r := range rucksacks {
		mid := len(r) / 2
		match := checkMatch(r[:mid], r[mid:])
		total += getPrio(match)
	}
	fmt.Println("Sum of priorities ", total)
}

// Finds character that appears in all input strings
func getMultiMatch(input []string) rune {
	encountered := make(map[rune]int, len(input))
	for _, i := range input {
		currEnc := make(map[rune]struct{}, len(i))
		// Map all characters in current string
		for _, c := range i {
			currEnc[c] = struct{}{}
		}
		// Add encountered characters to total string appearance count
		for k := range currEnc {
			encountered[k] += 1
			// If rune was encountered in all strings, we have the answer
			if encountered[k] == len(input) {
				return k
			}
		}
	}
	panic("No multi match found")
}

func Day3Part2() {
	rucksacks := helpers.ReadFileLines("../input/03")
	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		match := getMultiMatch(group)
		total += getPrio(match)
	}
	fmt.Println("Sum of priorities: ", total)
}
