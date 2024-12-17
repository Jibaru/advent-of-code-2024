package day16

import (
	"container/heap"
	"fmt"
	"os"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_16/input.txt"
	if isTest {
		f = "day_16/input-test.txt"
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
	maze := NewCharMatrixFromString(data)

	start, end := findStartAndEnd(maze)

	path := findPathWithLowestScore(maze, start, end)

	return path.score, nil
}

func partTwo(data string) (any, error) {
	maze := NewCharMatrixFromString(data)

	start, end := findStartAndEnd(maze)

	minPath := findPathWithLowestScore(maze, start, end)
	paths := []Path{minPath}

	for _, pos := range minPath.positions {
		backup := maze.Copy()

		maze.values[pos.i][pos.j] = "#"

		path := findPathWithLowestScore(maze, start, end)

		maze = backup

		if path.score == minPath.score {
			paths = append(paths, path)
		}
	}

	uniquePositions := make(map[Pair]bool)
	for _, path := range paths {
		for _, pos := range path.positions {
			uniquePositions[pos] = true
		}
	}

	uniquePositions[start] = true
	uniquePositions[end] = true

	// Contar las posiciones únicas
	return len(uniquePositions), nil
}

type Path struct {
	positions []Pair
	direction Pair
	score     int
}

// PriorityQueue is a min-heap of Paths based on their score.
type PriorityQueue []Path

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Path))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

var directions = []Pair{
	{i: 0, j: 1},  // East
	{i: 0, j: -1}, // West
	{i: -1, j: 0}, // North
	{i: 1, j: 0},  // South
}

func findPathWithLowestScore(maze *Matrix[string], start, end Pair) Path {
	pq := &PriorityQueue{}
	heap.Init(pq)

	startPath := Path{
		positions: []Pair{start},
		direction: directions[1], // Facing East initially
		score:     0,
	}
	heap.Push(pq, startPath)

	visited := make(map[Pair]int)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Path)
		currentPos := current.positions[len(current.positions)-1]

		if currentPos == end {
			return current
		}

		if v, ok := visited[currentPos]; ok && v <= current.score {
			continue
		}
		visited[currentPos] = current.score

		for _, dir := range directions {
			nextPos := currentPos.Add(dir)
			if !maze.InBounds(nextPos.i, nextPos.j) || maze.Val(nextPos) == "#" {
				continue
			}

			rotationCost := 0
			if dir != current.direction {
				rotationCost = 1000
			}

			nextPath := Path{
				positions: append([]Pair{}, append(current.positions, nextPos)...),
				direction: dir,
				score:     current.score + 1 + rotationCost,
			}
			heap.Push(pq, nextPath)
		}
	}

	return Path{}
}

func findStartAndEnd(maze *Matrix[string]) (start, end Pair) {
	for i := 0; i < maze.RowsLen(); i++ {
		for j := 0; j < maze.ColsLen(); j++ {
			if maze.At(i, j) == "S" {
				start = Pair{i, j}
			} else if maze.At(i, j) == "E" {
				end = Pair{i, j}
			}
		}
	}

	return
}
