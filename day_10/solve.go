package day10

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_10/input.txt"
	if isTest {
		f = "day_10/input-test.txt"
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
	mt := NewNumsMatrixFromString(data)

	ans := 0
	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if mt.At(i, j) == 0 {
				ans += FindScore(mt, i, j)
			}
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	mt := NewNumsMatrixFromString(data)

	ans := 0
	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if mt.At(i, j) == 0 {
				ans += FindRating(mt, i, j)
			}
		}
	}

	return ans, nil
}

// FindScore uses BFS
func FindScore(m *Matrix[int], i, j int) int {
	queue := NewQueue[Item]()
	queue.Push(Item{i, j, m.At(i, j)})

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	score := 0

	for !queue.IsEmpty() {
		item, _ := queue.Pop()

		if item.value == 9 {
			score++
			continue
		}

		next := item.value + 1

		for _, dir := range dirs {
			di := item.i + dir[0]
			dj := item.j + dir[1]

			if m.InBounds(di, dj) && m.At(di, dj) == next {
				newItem := Item{di, dj, m.At(di, dj)}
				if !queue.Has(newItem) {
					queue.Push(newItem)
				}
			}
		}
	}

	return score
}

// FindRating uses DFS
func FindRating(m *Matrix[int], i, j int) int {
	stack := NewStack[Item]()
	stack.Push(Item{i, j, m.At(i, j)})

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	rating := 0

	for !stack.IsEmpty() {
		item, _ := stack.Pop()

		if item.value == 9 {
			rating++
			continue
		}

		next := item.value + 1

		for _, dir := range dirs {
			di := item.i + dir[0]
			dj := item.j + dir[1]

			if m.InBounds(di, dj) && m.At(di, dj) == next {
				newItem := Item{di, dj, m.At(di, dj)}
				stack.Push(newItem)
			}
		}
	}

	return rating
}

type Matrix[T comparable] struct {
	values [][]T
	rsize  int
	csize  int
}

func NewNumsMatrixFromString(data string) *Matrix[int] {
	rows := strings.Split(data, "\n")
	var values [][]int
	for _, row := range rows {
		cols := []int{}
		for _, s := range strings.Split(row, "") {
			value, _ := strconv.ParseInt(s, 10, 64)
			cols = append(cols, int(value))
		}

		values = append(values, cols)
	}

	return &Matrix[int]{
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

type Item struct {
	i, j, value int
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

type Stack[T comparable] struct {
	items []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var v T
		return v, errors.New("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
