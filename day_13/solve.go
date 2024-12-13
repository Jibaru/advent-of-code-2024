package day13

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_13/input.txt"
	if isTest {
		f = "day_13/input-test.txt"
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
	ans := 0

	for _, d := range strings.Split(data, "\n\n") {
		eq, err := newEquation(d)
		if err != nil {
			return nil, err
		}

		a := eq.Buttons[0]
		b := eq.Buttons[1]
		p := eq.Prize

		y := (p.Y*a.X - p.X*a.Y) / (b.Y*a.X - b.X*a.Y)
		if !IsNaturalNumber(y) {
			continue
		}

		x := (p.X - b.X*y) / a.X
		if !IsNaturalNumber(x) {
			continue
		}

		aTokens := x * 3
		bTokens := y * 1

		ans += (int(aTokens) + int(bTokens))
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	ans := 0

	for _, d := range strings.Split(data, "\n\n") {
		eq, err := newEquation(d)
		if err != nil {
			return nil, err
		}

		eq.Prize.X += 10000000000000
		eq.Prize.Y += 10000000000000

		a := eq.Buttons[0]
		b := eq.Buttons[1]
		p := eq.Prize

		y := (p.Y*a.X - p.X*a.Y) / (b.Y*a.X - b.X*a.Y)
		if !IsNaturalNumber(y) {
			continue
		}

		x := (p.X - b.X*y) / a.X
		if !IsNaturalNumber(x) {
			continue
		}

		aTokens := x * 3
		bTokens := y * 1

		ans += (int(aTokens) + int(bTokens))
	}

	return ans, nil
}

type Button struct {
	Name string
	X    float64
	Y    float64
}

type Prize struct {
	X float64
	Y float64
}

type Equation struct {
	Buttons []Button
	Prize   Prize
}

func newEquation(input string) (Equation, error) {
	var equation Equation

	buttonRegex := regexp.MustCompile(`Button (\w+): X\+(\d+), Y\+(\d+)`)
	matches := buttonRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		x, _ := strconv.Atoi(match[2])
		y, _ := strconv.Atoi(match[3])
		equation.Buttons = append(equation.Buttons, Button{
			Name: match[1],
			X:    float64(x),
			Y:    float64(y),
		})
	}

	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	prizeMatch := prizeRegex.FindStringSubmatch(input)
	if len(prizeMatch) > 0 {
		x, _ := strconv.Atoi(prizeMatch[1])
		y, _ := strconv.Atoi(prizeMatch[2])
		equation.Prize = Prize{
			X: float64(x),
			Y: float64(y),
		}
	}

	return equation, nil
}

func IsNaturalNumber(f float64) bool {
	return f >= 0 && f == math.Trunc(f)
}
