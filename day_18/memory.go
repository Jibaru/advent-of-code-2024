package day18

type Memory[T comparable] struct {
	values [][]T
	rsize  int
	csize  int
}

func NewMemory(rows, cols int) *Memory[string] {
	var values [][]string
	for i := 0; i < rows; i++ {
		row := []string{}
		for j := 0; j < cols; j++ {
			row = append(row, ".")
		}
		values = append(values, row)
	}

	return &Memory[string]{
		values: values,
		rsize:  rows,
		csize:  cols,
	}
}

func (m *Memory[T]) RowsLen() int {
	return m.rsize
}

func (m *Memory[T]) ColsLen() int {
	return m.csize
}

func (m *Memory[T]) At(i, j int) T {
	return m.values[i][j]
}

func (m *Memory[T]) Val(p Byte) T {
	return m.values[p.i][p.j]
}

func (m *Memory[T]) InBounds(i int, j int) bool {
	return i >= 0 && i < m.RowsLen() && j >= 0 && j < m.ColsLen()
}

type Byte struct {
	i, j int
}

func (p Byte) Add(o Byte) Byte {
	return Byte{p.i + o.i, p.j + o.j}
}
