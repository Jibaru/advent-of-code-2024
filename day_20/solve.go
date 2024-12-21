package day20

import (
	"fmt"
	"os"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_20/input.txt"
	if isTest {
		f = "day_20/input-test.txt"
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
	grid := NewCharMatrixFromString(data)
	dists := filledDistancesMatrix(grid)

	ans := 0
	skipDirs := []Pair{{2, 0}, {1, 1}, {0, 2}, {-1, 1}, {-2, 0}, {-1, -1}, {0, -2}, {1, -1}}

	for i := 0; i < grid.RowsLen(); i++ {
		for j := 0; j < grid.ColsLen(); j++ {
			if grid.At(i, j) == "#" {
				continue
			}

			for _, dir := range skipDirs {
				di, dj := i+dir.i, j+dir.j

				if !grid.InBounds(di, dj) {
					continue
				}

				if grid.At(di, dj) == "#" {
					continue
				}

				if dists.At(i, j)-dists.At(di, dj) >= 102 {
					ans++
				}
			}
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	grid := NewCharMatrixFromString(data)
	dists := filledDistancesMatrix(grid)

	ans := 0
	maxRadius := 20

	for i := 0; i < grid.RowsLen(); i++ {
		for j := 0; j < grid.ColsLen(); j++ {
			if grid.At(i, j) == "#" {
				continue
			}

			for radius := 2; radius <= maxRadius; radius++ {
				for di := 0; di <= radius; di++ {
					dj := radius - di

					neighbors := map[Pair]bool{}
					for _, dir := range []Pair{
						{i + di, j + dj},
						{i + di, j - dj},
						{i - di, j + dj},
						{i - di, j - dj},
					} {
						neighbor := Pair{dir.i, dir.j}

						if !grid.InBounds(neighbor.i, neighbor.j) {
							continue
						}

						if grid.Val(neighbor) == "#" {
							continue
						}

						neighbors[neighbor] = true
					}

					for neighbor := range neighbors {
						d := dists.At(i, j) - dists.At(neighbor.i, neighbor.j)

						if d >= (100 + radius) {
							ans++
						}
					}
				}
			}
		}
	}

	return ans, nil
}

func findStart(grid *Matrix[string]) (start Pair) {
	for i := 0; i < grid.RowsLen(); i++ {
		for j := 0; j < grid.ColsLen(); j++ {
			if grid.At(i, j) == "S" {
				start = Pair{i, j}
				return
			}
		}
	}

	return
}

func distancesMatrix(cols, rows int) *Matrix[int] {
	initialValue := -1

	values := make([][]int, rows)
	for i := range values {
		values[i] = make([]int, cols)
		for j := range values[i] {
			values[i][j] = initialValue
		}
	}

	return &Matrix[int]{
		values: values,
		rsize:  rows,
		csize:  cols,
	}
}

func filledDistancesMatrix(grid *Matrix[string]) *Matrix[int] {
	dists := distancesMatrix(grid.ColsLen(), grid.RowsLen())

	start := findStart(grid)
	i, j := start.i, start.j
	dists.Set(i, j, 0)

	dirs := []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for grid.At(i, j) != "E" {
		for _, dir := range dirs {
			di, dj := i+dir.i, j+dir.j
			if !grid.InBounds(di, dj) {
				continue
			}

			if grid.At(di, dj) == "#" {
				continue
			}

			if dists.At(di, dj) != -1 {
				continue
			}

			dists.Set(di, dj, dists.At(i, j)+1)

			i = di
			j = dj
		}
	}

	return dists
}
