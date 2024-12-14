package day14

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_14/input.txt"
	if isTest {
		f = "day_14/input-test.txt"
	}

	body, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	switch part {
	case 1:
		return partOne(string(body), isTest)
	case 2:
		return partTwo(string(body), isTest)
	}

	return nil, fmt.Errorf("part should be only 1 or 2")
}

func partOne(data string, isTest bool) (any, error) {
	motions := []Motion{}

	for _, line := range strings.Split(data, "\n") {
		motion, err := newMotion(line)
		if err != nil {
			return nil, err
		}

		motions = append(motions, *motion)
	}

	xSize := 101
	if isTest {
		xSize = 11
	}

	ySize := 103
	if isTest {
		ySize = 7
	}
	times := 100

	firstQuadrant := [][]int{{0, 0}, {xSize/2 - 1, (ySize / 2) - 1}}
	secondQuadrant := [][]int{{(xSize / 2) + 1, 0}, {xSize - 1, (ySize / 2) - 1}}
	thirdQuadrant := [][]int{{0, (ySize / 2) + 1}, {(xSize / 2) - 1, ySize - 1}}
	fourthQuadrant := [][]int{{(xSize / 2) + 1, (ySize / 2) + 1}, {xSize - 1, ySize - 1}}

	totalFirst := 0
	totalSecond := 0
	totalThird := 0
	totalFourth := 0

	for _, m := range motions {
		x := newPos(m.Position.X, m.Velocity.X, xSize, times)
		y := newPos(m.Position.Y, m.Velocity.Y, ySize, times)

		switch {
		case x >= firstQuadrant[0][0] && x <= firstQuadrant[1][0] && y >= firstQuadrant[0][1] && y <= firstQuadrant[1][1]:
			totalFirst++
		case x >= secondQuadrant[0][0] && x <= secondQuadrant[1][0] && y >= secondQuadrant[0][1] && y <= secondQuadrant[1][1]:
			totalSecond++
		case x >= thirdQuadrant[0][0] && x <= thirdQuadrant[1][0] && y >= thirdQuadrant[0][1] && y <= thirdQuadrant[1][1]:
			totalThird++
		case x >= fourthQuadrant[0][0] && x <= fourthQuadrant[1][0] && y >= fourthQuadrant[0][1] && y <= fourthQuadrant[1][1]:
			totalFourth++
		}

	}

	return totalFirst * totalSecond * totalThird * totalFourth, nil
}

func newPos(actual int, vX int, size int, times int) int {
	d := vX * times

	newActual := actual + d

	newActual = newActual % size

	if newActual < 0 {
		newActual = size + newActual
	}

	return newActual
}

func partTwo(data string, isTest bool) (any, error) {
	motions := []Motion{}

	for _, line := range strings.Split(data, "\n") {
		motion, err := newMotion(line)
		if err != nil {
			return nil, err
		}

		motions = append(motions, *motion)
	}

	xSize := 101
	if isTest {
		xSize = 11
	}

	ySize := 103
	if isTest {
		ySize = 7
	}
	finished := false
	ans := 0

	for i := 1; !finished; i++ {
		visited := map[Vector]bool{}

		for _, m := range motions {
			x := newPos(m.Position.X, m.Velocity.X, xSize, i)
			y := newPos(m.Position.Y, m.Velocity.Y, ySize, i)

			v := Vector{x, y}

			if _, ok := visited[v]; ok {
				break
			}

			visited[v] = true
		}

		if len(visited) == len(motions) {
			finished = true
			ans = i
		}
	}

	return ans, nil
}

type Vector struct {
	X int
	Y int
}

type Motion struct {
	Position Vector
	Velocity Vector
}

func newMotion(input string) (*Motion, error) {
	regex := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	matches := regex.FindStringSubmatch(input)
	if matches == nil {
		return nil, fmt.Errorf("input does not match the expected format: %s", input)
	}

	posX, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("invalid position x: %v", err)
	}
	posY, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("invalid position y: %v", err)
	}
	velX, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid velocity x: %v", err)
	}
	velY, err := strconv.Atoi(matches[4])
	if err != nil {
		return nil, fmt.Errorf("invalid velocity y: %v", err)
	}

	// Create and return the Motion struct
	motion := &Motion{
		Position: Vector{X: posX, Y: posY},
		Velocity: Vector{X: velX, Y: velY},
	}
	return motion, nil
}
