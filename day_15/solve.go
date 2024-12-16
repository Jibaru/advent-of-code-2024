package day15

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrCanNotMove = errors.New("can not move")
)

var Dirs = map[string]Pair{
	">": {i: 0, j: 1},
	"v": {i: 1, j: 0},
	"<": {i: 0, j: -1},
	"^": {i: -1, j: 0},
}

func Solve(part int, isTest bool) (any, error) {
	f := "day_15/input.txt"
	if isTest {
		f = "day_15/input-test.txt"
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
	parts := strings.Split(data, "\n\n")

	mt := NewCharMatrixFromString(parts[0])
	blocks := map[string]bool{"O": true}

	current := StartPos(mt)

	for _, p := range strings.Split(parts[1], "") {
		dir := Dirs[p]

		newCurr, err := MoveTo(mt, current, dir, blocks)
		if err == nil {
			current = newCurr
		}
	}

	return SumPositions(BoxPositions(mt, "O")), nil
}

func partTwo(data string) (any, error) {
	parts := strings.Split(data, "\n\n")

	mt := NewCharMatrixFromString(parts[0])

	mt, err := TransformMatrix(mt)
	if err != nil {
		return nil, err
	}

	currentPos := StartPos(mt)
	blocks := map[string]bool{
		"[": true,
		"]": true,
	}

	for _, p := range strings.Split(parts[1], "") {
		dir := Dirs[p]
		backup := mt.Copy()

		var newPos Pair
		var err error

		if p == "<" || p == ">" {
			newPos, err = MoveTo(mt, currentPos, dir, blocks)
		} else {
			newPos, err = MoveUpOrDown(mt, currentPos, dir)
		}

		if err != nil {
			mt = backup
		} else {
			currentPos = newPos
		}
	}

	return SumPositions(BoxPositions(mt, "[")), nil
}

func StartPos(m *Matrix[string]) Pair {
	for i := 0; i < m.RowsLen(); i++ {
		for j := 0; j < m.ColsLen(); j++ {
			if m.At(i, j) == "@" {
				return Pair{i, j}
			}
		}
	}

	return Pair{-1, -1}
}

func SumPositions(boxPositions []Pair) int {
	ans := 0

	for _, pos := range boxPositions {
		ans += (100*pos.i + pos.j)
	}

	return ans
}

func MoveTo(m *Matrix[string], at, dir Pair, blocks map[string]bool) (Pair, error) {
	next := at.Add(dir)

	if !m.InBounds(next.i, next.j) {
		return Pair{}, ErrCanNotMove
	}

	if m.Val(next) == "#" {
		return Pair{}, ErrCanNotMove
	}

	if _, ok := blocks[m.Val(next)]; ok {
		_, err := MoveTo(m, next, dir, blocks)
		if err != nil {
			return Pair{}, err
		}
	}

	Swap(m, at, next)

	return next, nil
}

func IsPartOfABox(v string) bool {
	return v == "[" || v == "]"
}

func BoxPair(m *Matrix[string], p Pair) Pair {
	v := m.Val(p)

	other := Pair{p.i, p.j + 1}
	if v == "]" {
		other = Pair{p.i, p.j - 1}
	}

	return other
}

func AreSameBox(m *Matrix[string], a, b Pair) bool {
	if a.j < b.j && m.Val(a) == "[" && m.Val(b) == "]" {
		return true
	}

	if b.j < a.j && m.Val(b) == "[" && m.Val(a) == "]" {
		return true
	}

	return false
}

func MoveUpOrDown(m *Matrix[string], at, dir Pair) (Pair, error) {
	next := at.Add(dir)
	nextVal := m.Val(next)
	currVal := m.Val(at)

	if currVal == "@" {
		if nextVal == "." {
			Swap(m, at, next)
			return next, nil
		}

		if IsPartOfABox(nextVal) {
			_, err := MoveUpOrDown(m, next, dir)
			if err != nil {
				return Pair{}, ErrCanNotMove
			}

			Swap(m, at, next)
			return next, nil
		}

		// Wall
		return Pair{}, ErrCanNotMove
	}

	// "at" should be a block
	atPair := BoxPair(m, at)
	nextPair := atPair.Add(dir)
	nextPairVal := m.Val(nextPair)

	if nextVal == "#" || nextPairVal == "#" {
		return Pair{}, ErrCanNotMove
	}

	if nextPairVal == "." && IsPartOfABox(nextVal) {
		_, err := MoveUpOrDown(m, next, dir)
		if err != nil {
			return Pair{}, ErrCanNotMove
		}
	} else if nextVal == "." && IsPartOfABox(nextPairVal) {
		_, err := MoveUpOrDown(m, nextPair, dir)
		if err != nil {
			return Pair{}, ErrCanNotMove
		}
	} else if AreSameBox(m, next, nextPair) {
		_, err := MoveUpOrDown(m, next, dir)
		if err != nil {
			return Pair{}, ErrCanNotMove
		}
	} else if nextVal != "." && nextPairVal != "." && !AreSameBox(m, next, nextPair) {
		_, err := MoveUpOrDown(m, next, dir)
		if err != nil {
			return Pair{}, ErrCanNotMove
		}
		_, err = MoveUpOrDown(m, nextPair, dir)
		if err != nil {
			return Pair{}, ErrCanNotMove
		}
	}

	Swap(m, at, next)
	Swap(m, atPair, nextPair)
	return next, nil
}

func Swap(m *Matrix[string], from, to Pair) {
	aux := m.Val(to)
	m.values[to.i][to.j] = m.Val(from)
	m.values[from.i][from.j] = aux
}

func BoxPositions(m *Matrix[string], v string) []Pair {
	positions := []Pair{}

	for i := 0; i < m.RowsLen(); i++ {
		for j := 0; j < m.ColsLen(); j++ {
			if m.At(i, j) == v {
				positions = append(positions, Pair{i, j})
			}
		}
	}

	return positions
}

func TransformMatrix(m *Matrix[string]) (*Matrix[string], error) {
	transform := map[string][]string{
		"#": {"#", "#"},
		"O": {"[", "]"},
		".": {".", "."},
		"@": {"@", "."},
	}

	var newValues [][]string
	for i := 0; i < m.RowsLen(); i++ {
		var newRow []string
		for j := 0; j < m.ColsLen(); j++ {
			tile := m.At(i, j)
			if newTile, ok := transform[tile]; ok {
				newRow = append(newRow, newTile...)
			} else {
				return nil, fmt.Errorf("can not transform matrix")
			}
		}
		newValues = append(newValues, newRow)
	}

	return &Matrix[string]{
		values: newValues,
		rsize:  len(newValues),
		csize:  len(newValues[0]),
	}, nil
}
