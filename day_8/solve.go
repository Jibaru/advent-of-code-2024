package day8

import (
	"fmt"
	"os"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_8/input.txt"
	if isTest {
		f = "day_8/input-test.txt"
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

	freqs := map[string][]Point{}

	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if v := mt.At(i, j); v != "." {
				freqs[v] = append(freqs[v], Point{i: i, j: j})
			}
		}
	}

	visited := map[Point]bool{}
	ans := 0
	for _, points := range freqs {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				first, second := AntinodesOf(points[i], points[j])

				for _, anti := range []Point{first, second} {
					if _, ok := visited[anti]; ok {
						continue
					}

					if mt.InBounds(anti.i, anti.j) {
						ans++
						visited[anti] = true
					}
				}
			}
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	mt := newMatrix(data)

	freqs := map[string][]Point{}

	for i := 0; i < mt.RowsLen(); i++ {
		for j := 0; j < mt.ColsLen(); j++ {
			if v := mt.At(i, j); v != "." {
				freqs[v] = append(freqs[v], Point{i: i, j: j})
			}
		}
	}

	visited := map[Point]bool{}
	ans := 0
	for _, points := range freqs {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				antinodes := FullAntinodesOf(points[i], points[j], mt)
				antinodes = append(antinodes, points[i], points[j])

				for _, anti := range antinodes {
					if _, ok := visited[anti]; ok {
						continue
					}

					if mt.InBounds(anti.i, anti.j) {
						ans++
						visited[anti] = true
					}
				}
			}
		}
	}

	return ans, nil
}

type Point struct {
	i, j int
}

func (f Point) Distance(other Point) Point {
	return Point{
		i: f.i - other.i,
		j: f.j - other.j,
	}
}

func (f Point) Add(f2 Point) Point {
	return Point{f.i + f2.i, f.j + f2.j}
}

func (f Point) Substract(f2 Point) Point {
	return Point{f.i - f2.i, f.j - f2.j}
}

func AntinodesOf(a Point, b Point) (Point, Point) {
	dist := a.Distance(b)
	return a.Add(dist), b.Substract(dist)
}

func FullAntinodesOf(a Point, b Point, mt *Matrix) []Point {
	dist := a.Distance(b)
	first, second := AntinodesOf(a, b)
	antinodes := []Point{first, second}

	antinode := first
	for {
		newAntinode := antinode.Add(dist)
		if mt.InBounds(newAntinode.i, newAntinode.j) {
			antinodes = append(antinodes, newAntinode)
			antinode = newAntinode
		} else {
			break
		}
	}

	antinode = second
	for {
		newAntinode := antinode.Substract(dist)
		if mt.InBounds(newAntinode.i, newAntinode.j) {
			antinodes = append(antinodes, newAntinode)
			antinode = newAntinode
		} else {
			break
		}
	}

	return antinodes
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
