package day12

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_12/input.txt"
	if isTest {
		f = "day_12/input-test.txt"
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
	mt := NewCharMatrixFromString(data)

	visited := map[Cell]bool{}
	ans := 0

	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if _, ok := visited[Cell{i, j}]; ok {
				continue
			}

			region := findRegion(mt, i, j)

			for k := range region {
				visited[k] = true
			}

			a := area(region)
			p := perimeter(mt, region, mt.At(i, j))

			ans += (p * a)
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	mt := NewCharMatrixFromString(data)

	visited := map[Cell]bool{}
	ans := 0

	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if _, ok := visited[Cell{i, j}]; ok {
				continue
			}

			region := findRegion(mt, i, j)

			for k := range region {
				visited[k] = true
			}

			a := area(region)
			c := corners(mt, region, mt.At(i, j))

			ans += (c * a)
		}
	}

	return ans, nil
}

func findRegion(m *Matrix[string], i, j int) map[Cell]bool {
	region := map[Cell]bool{}

	queue := NewQueue[Cell]()
	queue.Push(Cell{i, j})
	plantType := m.At(i, j)

	dirs := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for !queue.IsEmpty() {
		cell, _ := queue.Pop()

		if _, ok := region[cell]; ok {
			continue
		}

		region[cell] = true

		for _, dir := range dirs {
			di := cell.i + dir[0]
			dj := cell.j + dir[1]
			if m.InBounds(di, dj) && m.At(di, dj) == plantType {
				queue.Push(Cell{di, dj})
			}
		}
	}

	return region
}

func area(region map[Cell]bool) int {
	return len(region)
}

func perimeter(m *Matrix[string], region map[Cell]bool, plantType string) int {
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	p := 0

	for cell := range region {
		for _, dir := range dirs {
			di := cell.i + dir[0]
			dj := cell.j + dir[1]

			if !m.InBounds(di, dj) {
				p++
				continue
			}

			if m.At(di, dj) != plantType {
				p++
				continue
			}
		}
	}

	return p
}

func corners(m *Matrix[string], region map[Cell]bool, plantType string) int {
	u := []int{-1, 0}
	d := []int{1, 0}
	l := []int{0, -1}
	r := []int{0, 1}
	ul := []int{-1, -1}
	ur := []int{-1, 1}
	dl := []int{1, -1}
	dr := []int{1, 1}

	outCorners := 0
	inCorners := 0

	for cell := range region {
		up := Cell{cell.i + u[0], cell.j + u[1]}
		down := Cell{cell.i + d[0], cell.j + d[1]}
		left := Cell{cell.i + l[0], cell.j + l[1]}
		right := Cell{cell.i + r[0], cell.j + r[1]}
		upLeft := Cell{cell.i + ul[0], cell.j + ul[1]}
		upRight := Cell{cell.i + ur[0], cell.j + ur[1]}
		downLeft := Cell{cell.i + dl[0], cell.j + dl[1]}
		downRight := Cell{cell.i + dr[0], cell.j + dr[1]}

		if (!m.InBounds(left.i, left.j) || m.At(left.i, left.j) != plantType) &&
			(!m.InBounds(up.i, up.j) || m.At(up.i, up.j) != plantType) {
			outCorners++
		}

		if (!m.InBounds(right.i, right.j) || m.At(right.i, right.j) != plantType) &&
			(!m.InBounds(up.i, up.j) || m.At(up.i, up.j) != plantType) {
			outCorners++
		}

		if (!m.InBounds(left.i, left.j) || m.At(left.i, left.j) != plantType) &&
			(!m.InBounds(down.i, down.j) || m.At(down.i, down.j) != plantType) {
			outCorners++
		}

		if (!m.InBounds(right.i, right.j) || m.At(right.i, right.j) != plantType) &&
			(!m.InBounds(down.i, down.j) || m.At(down.i, down.j) != plantType) {
			outCorners++
		}

		if m.InBounds(left.i, left.j) && m.At(left.i, left.j) == plantType &&
			m.InBounds(up.i, up.j) && m.At(up.i, up.j) == plantType &&
			m.InBounds(upLeft.i, upLeft.j) && m.At(upLeft.i, upLeft.j) != plantType {
			inCorners++
		}

		if m.InBounds(right.i, right.j) && m.At(right.i, right.j) == plantType &&
			m.InBounds(up.i, up.j) && m.At(up.i, up.j) == plantType &&
			m.InBounds(upRight.i, upRight.j) && m.At(upRight.i, upRight.j) != plantType {
			inCorners++
		}

		if m.InBounds(left.i, left.j) && m.At(left.i, left.j) == plantType &&
			m.InBounds(down.i, down.j) && m.At(down.i, down.j) == plantType &&
			m.InBounds(downLeft.i, downLeft.j) && m.At(downLeft.i, downLeft.j) != plantType {
			inCorners++
		}

		if m.InBounds(right.i, right.j) && m.At(right.i, right.j) == plantType &&
			m.InBounds(down.i, down.j) && m.At(down.i, down.j) == plantType &&
			m.InBounds(downRight.i, downRight.j) && m.At(downRight.i, downRight.j) != plantType {
			inCorners++
		}
	}

	return outCorners + inCorners
}

type Matrix[T comparable] struct {
	values [][]T
	rsize  int
	csize  int
}

type Cell struct {
	i, j int
}

func NewCharMatrixFromString(data string) *Matrix[string] {
	rows := strings.Split(data, "\n")
	var values [][]string
	for _, row := range rows {
		values = append(values, strings.Split(row, ""))
	}

	return &Matrix[string]{
		values: values,
		rsize:  len(rows),
		csize:  len(rows[0]),
	}
}

func (m *Matrix[T]) RowsLen() int {
	return m.rsize
}

func (m *Matrix[T]) ColsLen() int {
	return m.csize
}

func (m *Matrix[T]) At(i, j int) T {
	return m.values[i][j]
}

func (m *Matrix[T]) InBounds(i int, j int) bool {
	return i >= 0 && i < m.RowsLen() && j >= 0 && j < m.ColsLen()
}

type Queue[T comparable] struct {
	items    []T
	existing map[T]bool
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		items:    make([]T, 0),
		existing: make(map[T]bool, 0),
	}
}

func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
	q.existing[item] = true
}

func (q *Queue[T]) Pop() (T, error) {
	if len(q.items) == 0 {
		var v T
		return v, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]

	delete(q.existing, item)

	return item, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Has(item T) bool {
	_, exists := q.existing[item]
	return exists
}
