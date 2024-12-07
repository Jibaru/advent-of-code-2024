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

type Pair struct {
	i, j int
}

func partOne(data string) (any, error) {
	mt := newMatrix(data)

	i, j, val := mt.StartPos()

	visited := map[Pair]bool{
		{i, j}: true,
	}

	ans := 1
	finished := false

	for !finished {
		i, j, val = mt.Next(i, j, val)

		if i == -1 && j == -1 {
			finished = true
			continue
		}

		if _, ok := visited[Pair{i, j}]; ok {
			continue
		}

		visited[Pair{i, j}] = true
		ans++
	}

	return ans, nil
}

func partTwo(_ string) (any, error) {
	return "not solved yet", nil
}

type matrix struct {
	values [][]string
	rsize  int
	csize  int
}

func newMatrix(data string) *matrix {
	rows := strings.Split(data, "\n")
	var values [][]string
	for _, row := range rows {
		values = append(values, strings.Split(row, ""))
	}

	return &matrix{
		values: values,
		rsize:  len(rows),
		csize:  len(rows[0]),
	}
}

func (m *matrix) RowsLen() int {
	return m.rsize
}

func (m *matrix) ColsLen() int {
	return m.csize
}

func (m *matrix) At(i, j int) string {
	return m.values[i][j]
}

func (m *matrix) InBounds(i int, j int) bool {
	return i >= 0 && i < m.RowsLen() && j >= 0 && j < m.ColsLen()
}

func (m *matrix) IsBlock(i, j int) bool {
	return m.values[i][j] == "#"
}

func (m *matrix) StartPos() (int, int, string) {
	for i, row := range m.values {
		for j, cell := range row {
			if cell == "^" || cell == "v" || cell == "<" || cell == ">" {
				return i, j, cell
			}
		}
	}

	return -1, -1, ""
}

func (m *matrix) Next(i, j int, v string) (int, int, string) {
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

	di := i + dirs[v][0]
	dj := j + dirs[v][1]
	nv := v

	if !m.InBounds(di, dj) {
		return -1, -1, ""
	}

	if m.IsBlock(di, dj) {
		nv = nexts[v]
		di = i + dirs[nv][0]
		dj = j + dirs[nv][1]
	}

	return di, dj, nv
}
