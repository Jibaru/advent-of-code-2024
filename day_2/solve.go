package day0

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_2/input.txt"
	if isTest {
		f = "day_2/input-test.txt"
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
	reports := strings.Split(data, "\n")

	ans := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")
		numLevels := levelsToNums(levels)

		if isSafe(numLevels) {
			ans++
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	reports := strings.Split(data, "\n")

	wg := &sync.WaitGroup{}
	wg.Add(len(reports))
	mut := &sync.Mutex{}
	ans := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")
		numLevels := levelsToNums(levels)

		go func() {
			defer wg.Done()

			if isSafe(numLevels) {
				mut.Lock()
				ans++
				mut.Unlock()
				return
			}

			for i := 0; i < len(numLevels); i++ {
				modifiedLevels := append([]int64{}, numLevels[:i]...)
				modifiedLevels = append(modifiedLevels, numLevels[i+1:]...)

				if isSafe(modifiedLevels) {
					mut.Lock()
					ans++
					mut.Unlock()
					return
				}
			}

		}()
	}

	wg.Wait()

	return ans, nil
}

func levelsToNums(levels []string) []int64 {
	numLevels := make([]int64, len(levels))
	for i, level := range levels {
		numLevels[i], _ = strconv.ParseInt(level, 10, 64)
	}
	return numLevels
}

func isSafe(levels []int64) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := false
	if levels[1] > levels[0] {
		increasing = true
	}

	for i := 1; i < len(levels); i++ {
		prev := levels[i-1]
		curr := levels[i]

		if curr > prev && !increasing {
			return false
		}

		if curr < prev && increasing {
			return false
		}

		diff := math.Abs(float64(prev) - float64(curr))
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}

	return true
}
