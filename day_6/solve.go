package day6

import (
	"fmt"
	"os"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_6/input.txt"
	if isTest {
		f = "day_6/input-test.txt"
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
	mt := newMatrix(data)

	start, err := startPath(mt)
	if err != nil {
		return nil, err
	}

	return len(guardPositions(mt, start)), nil
}

func partTwo(data string) (any, error) {
	mt := newMatrix(data)

	start, err := startPath(mt)
	if err != nil {
		return nil, err
	}

	positions := guardPositions(mt, start)

	ans := 0

	for pos := range positions {
		if pos == start.pos {
			continue
		}

		mt.values[pos.i][pos.j] = "#"

		if hasLoop(mt, start) {
			ans++
		}

		mt.values[pos.i][pos.j] = "."
	}

	return ans, nil
}

type Path struct {
	pos Pos
	dir string
}

type Pos struct {
	i, j int
}

type Matrix struct {
	values [][]string
	rsize  int
	csize  int
}

func newMatrix(data string) *Matrix {
	rows := strings.Split(data, "\n")
	var values [][]string
	for _, row := range rows {
		values = append(values, strings.Split(row, ""))
	}

	return &Matrix{
		values: values,
		rsize:  len(rows),
		csize:  len(rows[0]),
	}
}

func (m *Matrix) RowsLen() int {
	return m.rsize
}

func (m *Matrix) ColsLen() int {
	return m.csize
}

func (m *Matrix) At(i, j int) string {
	return m.values[i][j]
}

func (m *Matrix) InBounds(i int, j int) bool {
	return i >= 0 && i < m.RowsLen() && j >= 0 && j < m.ColsLen()
}

func (m *Matrix) IsBlock(i, j int) bool {
	return m.values[i][j] == "#"
}

func startPath(m *Matrix) (Path, error) {
	for i, row := range m.values {
		for j, cell := range row {
			if cell == "^" || cell == "v" || cell == "<" || cell == ">" {
				return Path{pos: Pos{i, j}, dir: cell}, nil
			}
		}
	}

	return Path{}, fmt.Errorf("start not found")
}

func isOut(path Path) bool {
	return path.pos.i == -1 && path.pos.j == -1
}

func guardPositions(m *Matrix, start Path) map[Pos]bool {
	visited := map[Pos]bool{
		start.pos: true,
	}

	finished := false
	path := start

	for !finished {
		path = next(m, path)

		if isOut(path) {
			finished = true
			continue
		}

		if _, ok := visited[path.pos]; ok {
			continue
		}

		visited[path.pos] = true
	}

	return visited
}

func next(m *Matrix, curr Path) Path {
	dirs := map[string][]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}

	nexts := map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}

	di := curr.pos.i + dirs[curr.dir][0]
	dj := curr.pos.j + dirs[curr.dir][1]
	newDir := curr.dir

	if !m.InBounds(di, dj) {
		return Path{Pos{-1, -1}, newDir}
	}

	if m.IsBlock(di, dj) {
		newDir = nexts[curr.dir]
		di = curr.pos.i
		dj = curr.pos.j
	}

	return Path{Pos{di, dj}, newDir}
}

func hasLoop(mt *Matrix, start Path) bool {
	visited := map[Path]bool{
		start: true,
	}

	path := start

	for {
		path = next(mt, path)

		if isOut(path) {
			return false
		}

		if visited[path] {
			return true
		}

		visited[path] = true
	}
}
