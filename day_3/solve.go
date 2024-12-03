package day0

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_3/input.txt"
	if isTest {
		f = fmt.Sprintf("day_3/input-test-%v.txt", part)
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

type multiplication struct {
	multiplicand int
	multiplier   int
}

func (m multiplication) Solve() int {
	return m.multiplicand * m.multiplier
}

func newMultiplication(input string) (multiplication, error) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	match := re.FindStringSubmatch(input)

	if len(match) != 3 {
		return multiplication{}, fmt.Errorf("input is invalid")
	}

	num1, err := strconv.Atoi(match[1])
	if err != nil {
		return multiplication{}, fmt.Errorf("can not parse: %v", err)
	}

	num2, err := strconv.Atoi(match[2])
	if err != nil {
		return multiplication{}, fmt.Errorf("can not parse: %v", err)
	}

	return multiplication{multiplicand: num1, multiplier: num2}, nil
}

func partOne(data string) (any, error) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(data, -1)

	ans := 0

	for _, match := range matches {
		m, err := newMultiplication(match)
		if err != nil {
			return nil, err
		}

		ans += m.Solve()
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	matches := re.FindAllString(data, -1)

	enabled := true
	ans := 0

	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}

		if match == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		mp, err := newMultiplication(match)
		if err != nil {
			return nil, err
		}

		ans += mp.Solve()
	}

	return ans, nil
}
