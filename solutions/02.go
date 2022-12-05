package main

import (
	"adventofgode-2022/helpers"
	"fmt"
	"log"
)

const (
	drawScore = 3
	winScore  = 6
)

func checkWin(enemy, mine byte) int {
	res := int(int(mine) - int(enemy))
	if res < 0 {
		return 3 + res
	}
	return res
}

func Day2Part1() {
	lines := helpers.ReadFileLines("../input/02")
	score := 0
	for _, l := range lines {
		enemy := l[0] - 'A'
		mine := l[2] - 'X'
		res := checkWin(enemy, mine)
		if res == 0 {
			score += drawScore
		} else if res == 1 {
			score += winScore
		} else if res != 2 {
			log.Fatal("Something went wrong")
		}
		score += int(mine) + 1
	}
	fmt.Println("Score: ", score)
}

func Day2Part2() {
	lines := helpers.ReadFileLines("../input/02")
	score := 0
	for _, l := range lines {
		enemy := l[0] - 'A'
		// 0 -> loss (2)
		// 1 -> draw (0)
		// 2 -> win (1)
		desired := int(l[2]-'X') - 1
		if desired < 0 {
			desired = 3 + desired
		}

		// Naive approach but math is hard
		toPlay := -1
		for i := 0; i < 3; i++ {
			if res := checkWin(enemy, byte(i)); res == desired {
				toPlay = i
			}
		}
		if toPlay == -1 {
			log.Fatal("uh oh")
		}

		if desired == 0 {
			score += drawScore
		} else if desired == 1 {
			score += winScore
		} else if desired != 2 {
			log.Fatal("Something went wrong")
		}
		score += toPlay + 1
	}
	fmt.Println("Score: ", score)
}
