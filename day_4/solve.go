package day4

import (
	"fmt"
	"os"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_4/input.txt"
	if isTest {
		f = "day_4/input-test.txt"
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

func (m *matrix) HasXmasWord(i, j, di, dj int) bool {
	ci := i
	cj := j
	word := []string{"X", "M", "A", "S"}
	wordIdx := 0

	for m.InBounds(ci, cj) {
		letter := word[wordIdx]
		if m.At(ci, cj) == letter {
			if letter == "S" {
				return true
			}

			wordIdx++
			ci += di
			cj += dj
		} else {
			break
		}
	}

	return false
}

func (m *matrix) HasXmasShape(i, j int) bool {
	if m.At(i, j) != "A" {
		return false
	}

	if !m.InBounds(i-1, j-1) ||
		!m.InBounds(i+1, j-1) ||
		!m.InBounds(i+1, j+1) ||
		!m.InBounds(i-1, j+1) {

		return false
	}

	return ((m.At(i-1, j-1) == "M" && m.At(i+1, j+1) == "S") ||
		(m.At(i-1, j-1) == "S" && m.At(i+1, j+1) == "M")) &&
		((m.At(i+1, j-1) == "M" && m.At(i-1, j+1) == "S") ||
			(m.At(i+1, j-1) == "S" && m.At(i-1, j+1) == "M"))
}

func partOne(data string) (any, error) {
	mt := newMatrix(data)

	dirs := [][]int{
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
	}

	ans := 0

	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			for _, dir := range dirs {
				if mt.HasXmasWord(i, j, dir[0], dir[1]) {
					ans++
				}
			}
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	mt := newMatrix(data)
	ans := 0

	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if mt.HasXmasShape(i, j) {
				ans++
			}
		}
	}

	return ans, nil
}
