package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_11/input.txt"
	if isTest {
		f = "day_11/input-test.txt"
	}

	body, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	switch part {
	case 1:
		return partOne(string(body))
	case 2:
		return partTwo(string(body))
	}

	return nil, fmt.Errorf("part should be only 1 or 2")
}

func partOne(data string) (any, error) {
	initialStones := parseInput(data)
	stones := blinks(initialStones, 25)
	return len(stones), nil
}

func partTwo(data string) (any, error) {
	initialStones := parseInput(data)
	stoneCounts := map[int]int{}

	for _, stone := range initialStones {
		stoneCounts[stone]++
	}

	stoneCounts = blinkCounts(stoneCounts, 75)

	ans := 0
	for _, count := range stoneCounts {
		ans += count
	}

	return ans, nil
}

func parseInput(data string) []int {
	parts := strings.Fields(data)
	stones := make([]int, len(parts))
	for i, p := range parts {
		num, _ := strconv.Atoi(p)
		stones[i] = num
	}
	return stones
}

func blinks(stones []int, blinks int) []int {
	for b := 0; b < blinks; b++ {
		nextStones := []int{}
		for _, stone := range stones {
			if stone == 0 {
				nextStones = append(nextStones, 1)
			} else if isEvenDigit(stone) {
				left, right := splitStone(stone)
				nextStones = append(nextStones, left, right)
			} else {
				nextStones = append(nextStones, stone*2024)
			}
		}
		stones = nextStones
	}
	return stones
}

func isEvenDigit(num int) bool {
	digitCount := len(strconv.Itoa(num))
	return digitCount%2 == 0
}

func splitStone(num int) (int, int) {
	str := strconv.Itoa(num)
	mid := len(str) / 2
	left, _ := strconv.Atoi(str[:mid])
	right, _ := strconv.Atoi(str[mid:])
	return left, right
}

func blinkCounts(stoneCounts map[int]int, blinks int) map[int]int {
	for b := 0; b < blinks; b++ {
		newCounts := make(map[int]int)
		for stone, count := range stoneCounts {
			if stone == 0 {
				newCounts[1] += count
			} else if isEvenDigit(stone) {
				left, right := splitStone(stone)
				newCounts[left] += count
				newCounts[right] += count
			} else {
				newCounts[stone*2024] += count
			}
		}
		stoneCounts = newCounts
	}

	return stoneCounts
}
