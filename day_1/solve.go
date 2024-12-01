package day1

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_1/input.txt"
	if isTest {
		f = "day_1/input-test.txt"
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
	firstLocations := []int{}
	secondLocations := []int{}

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, "   ")
		location1, _ := strconv.ParseInt(parts[0], 10, 32)
		location2, _ := strconv.ParseInt(parts[1], 10, 32)

		firstLocations = append(firstLocations, int(location1))
		secondLocations = append(secondLocations, int(location2))
	}

	slices.Sort(firstLocations)
	slices.Sort(secondLocations)

	ans := 0

	for i, firstLocation := range firstLocations {
		ans += int(math.Abs(float64(firstLocation) - float64(secondLocations[i])))
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	lines := strings.Split(data, "\n")

	freqs := map[int]int{}
	nums := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		numInFirstList, _ := strconv.ParseInt(parts[0], 10, 32)
		numInSecondList, _ := strconv.ParseInt(parts[1], 10, 32)

		nums = append(nums, int(numInFirstList))
		freqs[int(numInSecondList)] += 1
	}

	ans := 0
	for _, num := range nums {
		ans += (num * freqs[num])
	}

	return ans, nil
}
