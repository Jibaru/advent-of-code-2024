package day15

import (
	"fmt"
	"strings"
)

type Matrix[T comparable] struct {
	values [][]T
	rsize  int
	csize  int
}

func (mt *Matrix[T]) String() string {
	var result string
	for i := 0; i < mt.rsize; i++ {
		for j := 0; j < mt.csize; j++ {
			result += fmt.Sprintf("%v", mt.values[i][j])
		}
		result += "\n"
	}
	return result
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

func (m *Matrix[T]) Copy() *Matrix[T] {
	newValues := make([][]T, m.rsize)
	for i := range m.values {
		newRow := make([]T, m.csize)
		copy(newRow, m.values[i])
		newValues[i] = newRow
	}

	return &Matrix[T]{
		values: newValues,
		rsize:  m.rsize,
		csize:  m.csize,
	}
}

type Pair struct {
	i, j int
}

func (p Pair) Add(o Pair) Pair {
	return Pair{p.i + o.i, p.j + o.j}
}
