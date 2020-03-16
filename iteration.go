package sudoku

// iteration is a single iteration of a puzzle.
type iteration struct {
	// index is the cell index that this iteration will effect
	index int
	// minValue is the min value this iteration used when effecting the cell
	minValue int

	puzzleSize  int
	sectionSize int
	cells       group
	rows        [][]int
	columns     [][]int
	sections    [][]int
}

// newIteration returns a new iteration with the given items.
func newIteration(items []int, puzzleSize int, sectionSize int) *iteration {
	rows := make([][]int, puzzleSize)
	columns := make([][]int, puzzleSize)
	sections := make([][]int, puzzleSize)

	it := &iteration{
		puzzleSize:  puzzleSize,
		sectionSize: sectionSize,
		cells:       make(group, len(items)),
		rows:        rows,
		columns:     columns,
		sections:    sections,
	}
	for i, item := range items {
		c := cell{
			fixed: item > 0,
			value: item,
			index: i,
		}
		it.cells[i] = &c

		rowIndex := getRowFromIndex(c.index, it.puzzleSize)
		if it.rows[rowIndex] == nil {
			it.rows[rowIndex] = make([]int, 0, it.puzzleSize)
		}
		it.rows[rowIndex] = append(it.rows[rowIndex], c.index)

		columnIndex := getColumnFromIndex(c.index, it.puzzleSize)
		if it.columns[columnIndex] == nil {
			it.columns[columnIndex] = make([]int, 0, it.puzzleSize)
		}
		it.columns[columnIndex] = append(it.columns[columnIndex], c.index)

		sectionIndex := getSectionFromIndex(c.index, it.puzzleSize, it.sectionSize)
		if it.sections[sectionIndex] == nil {
			it.sections[sectionIndex] = make([]int, 0, it.puzzleSize)
		}
		it.sections[sectionIndex] = append(it.sections[sectionIndex], c.index)
	}
	return it
}

// items returns all of the items in the iteration.
func (i *iteration) items() []int {
	res := make([]int, len(i.cells))
	for i, c := range i.cells {
		res[i] = c.value
	}
	return res
}

// iterate returns a iterate of the iteration.
func (i *iteration) iterate() *iteration {
	it := &iteration{
		puzzleSize:  i.puzzleSize,
		sectionSize: i.sectionSize,
		cells:       make(group, len(i.cells)),
		rows:        i.rows,
		columns:     i.columns,
		sections:    i.sections,
		minValue:    1,
		index:       i.index + 1,
	}
	for i, item := range i.cells {
		it.cells[i] = item.copy()
	}
	return it
}

// finished returns true if there are no zero values left in the iteration.
func (i *iteration) finished() bool {
	for _, c := range i.cells {
		if c.value == 0 {
			return false
		}
	}
	return true
}

// solve attempts to solve the current iteration.
func (i *iteration) solve() error {
	cell := i.cells[i.index]
	if cell.fixed {
		return nil
	}

	nextValue, err := i.findNextValue(cell, i.minValue)
	if err != nil {
		return err
	}
	i.minValue = nextValue
	cell.value = nextValue
	return nil
}

func (i *iteration) findNextValue(cell *cell, minValue int) (int, error) {
	usedMap := map[int]bool{}

	for _, cellIndex := range i.rows[getRowFromIndex(cell.index, i.puzzleSize)] {
		if usedMap[i.cells[cellIndex].value] || i.cells[cellIndex].value < minValue {
			continue
		}
		usedMap[i.cells[cellIndex].value] = true
	}
	for _, cellIndex := range i.columns[getColumnFromIndex(cell.index, i.puzzleSize)] {
		if usedMap[i.cells[cellIndex].value] || i.cells[cellIndex].value < minValue {
			continue
		}
		usedMap[i.cells[cellIndex].value] = true
	}
	for _, cellIndex := range i.sections[getSectionFromIndex(cell.index, i.puzzleSize, i.sectionSize)] {
		if usedMap[i.cells[cellIndex].value] || i.cells[cellIndex].value < minValue {
			continue
		}
		usedMap[i.cells[cellIndex].value] = true
	}

	for value := minValue; value <= i.puzzleSize; value++ {
		if usedMap[value] {
			continue
		}
		return value, nil
	}

	return 0, ErrNoMoreMoves
}
