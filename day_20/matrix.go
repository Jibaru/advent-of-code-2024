package day20

import (
	"strings"
)

type Matrix[T comparable] struct {
	values [][]T
	rsize  int
	csize  int
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

func (m *Matrix[T]) Val(p Pair) T {
	return m.values[p.i][p.j]
}

func (m *Matrix[T]) InBounds(i int, j int) bool {
	return i >= 0 && i < m.RowsLen() && j >= 0 && j < m.ColsLen()
}

func (m *Matrix[T]) Set(i, j int, value T) {
	m.values[i][j] = value
}

type Pair struct {
	i, j int
}
