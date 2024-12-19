package day18

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_18/input.txt"
	if isTest {
		f = "day_18/input-test.txt"
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
	bytes := newBytesFromString(data)

	var memory *Memory[string]
	var totalBytesToUse int
	if isTest {
		memory = NewMemory(7, 7)
		totalBytesToUse = 12
	} else {
		memory = NewMemory(71, 71)
		totalBytesToUse = 1024
	}

	start := Byte{0, 0}
	end := Byte{memory.RowsLen() - 1, memory.ColsLen() - 1}

	path, err := findSortestPath(memory, bytes, totalBytesToUse, start, end)
	if err != nil {
		return nil, err
	}

	return path.Steps(), nil
}

func partTwo(data string, isTest bool) (any, error) {
	bytes := newBytesFromString(data)

	var memory *Memory[string]
	var startAt int
	if isTest {
		memory = NewMemory(7, 7)
		startAt = 12
	} else {
		memory = NewMemory(71, 71)
		startAt = 1024
	}

	start := Byte{0, 0}
	end := Byte{memory.RowsLen() - 1, memory.ColsLen() - 1}

	var block *Byte

	for totalBytesToUse := startAt; block == nil; totalBytesToUse++ {
		_, err := findSortestPath(memory, bytes, totalBytesToUse, start, end)
		if err != nil {
			block = &bytes[totalBytesToUse-1]
		}
	}

	return fmt.Sprintf("%v,%v", block.j, block.i), nil
}

func newBytesFromString(data string) []Byte {
	bytes := []Byte{}

	for _, part := range strings.Split(data, "\n") {
		xy := strings.Split(part, ",")

		x, _ := strconv.ParseInt(xy[1], 10, 64)
		y, _ := strconv.ParseInt(xy[0], 10, 64)

		bytes = append(bytes, Byte{i: int(x), j: int(y)})
	}

	return bytes
}

type Path struct {
	steps    int
	sequence []Byte
}

func (p Path) Steps() int {
	return p.steps
}

func findSortestPath(memory *Memory[string], bytes []Byte, totalBytesToUse int, start, end Byte) (Path, error) {
	directions := []Byte{
		{i: 0, j: 1},  // right
		{i: 1, j: 0},  // down
		{i: 0, j: -1}, // left
		{i: -1, j: 0}, // up
	}

	// Simulate falling bytes
	for idx, b := range bytes {
		if idx >= totalBytesToUse {
			break
		}
		if memory.InBounds(b.i, b.j) {
			memory.values[b.i][b.j] = "#" // Mark as corrupted
		}
	}

	// BFS setup
	queue := []struct {
		position Byte
		steps    int
		path     []Byte
	}{
		{start, 0, []Byte{start}},
	}

	visited := make(map[Byte]bool)
	visited[start] = true

	// BFS to find the shortest path
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we reach the end, return the path
		if current.position == end {
			return Path{
				steps:    current.steps,
				sequence: current.path,
			}, nil
		}

		// Explore neighbors
		for _, dir := range directions {
			neighbor := current.position.Add(dir)

			if memory.InBounds(neighbor.i, neighbor.j) && // Must be within bounds
				memory.At(neighbor.i, neighbor.j) == "." && // Must be safe
				!visited[neighbor] { // Must not be visited

				visited[neighbor] = true
				queue = append(queue, struct {
					position Byte
					steps    int
					path     []Byte
				}{
					position: neighbor,
					steps:    current.steps + 1,
					path:     append(append([]Byte{}, current.path...), neighbor),
				})
			}
		}
	}

	return Path{}, errors.New("can not reach end")
}
