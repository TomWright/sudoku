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
	rows        groups
	columns     groups
	sections    groups
}

// newIteration returns a new iteration with the given items.
func newIteration(items []int, puzzleSize int, sectionSize int) *iteration {
	it := &iteration{
		puzzleSize:  puzzleSize,
		sectionSize: sectionSize,
		cells:       make(group, len(items)),
		rows:        make(groups, puzzleSize),
		columns:     make(groups, puzzleSize),
		sections:    make(groups, puzzleSize),
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
			it.rows[rowIndex] = make(group, 0, it.puzzleSize)
		}
		it.rows[rowIndex] = append(it.rows[rowIndex], &c)

		columnIndex := getColumnFromIndex(c.index, it.puzzleSize)
		if it.columns[columnIndex] == nil {
			it.columns[columnIndex] = make(group, 0, it.puzzleSize)
		}
		it.columns[columnIndex] = append(it.columns[columnIndex], &c)

		sectionIndex := getSectionFromIndex(c.index, it.puzzleSize, it.sectionSize)
		if it.sections[sectionIndex] == nil {
			it.sections[sectionIndex] = make(group, 0, it.puzzleSize)
		}
		it.sections[sectionIndex] = append(it.sections[sectionIndex], &c)
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

// copy returns a copy of the iteration.
func (i *iteration) copy() *iteration {
	it := &iteration{
		puzzleSize:  i.puzzleSize,
		sectionSize: i.sectionSize,
		cells:       make(group, len(i.cells)),
		rows:        make(groups, i.puzzleSize),
		columns:     make(groups, i.puzzleSize),
		sections:    make(groups, i.puzzleSize),
		minValue:    1,
		index:       i.index,
	}
	for i, item := range i.cells {
		c := cell{
			fixed: item.fixed,
			value: item.value,
			index: i,
		}
		it.cells[i] = &c
	}

	it.rows = make([]group, len(i.rows))
	for i, row := range i.rows {
		it.rows[i] = make(group, len(row))
		for ci, c := range row {
			it.rows[i][ci] = it.cells[c.index]
		}
	}
	it.columns = make([]group, len(i.columns))
	for i, column := range i.columns {
		it.columns[i] = make(group, len(column))
		for ci, c := range column {
			it.columns[i][ci] = it.cells[c.index]
		}
	}
	it.sections = make([]group, len(i.sections))
	for i, section := range i.sections {
		it.sections[i] = make(group, len(section))
		for ci, c := range section {
			it.sections[i][ci] = it.cells[c.index]
		}
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

	row := i.rows[getRowFromIndex(cell.index, i.puzzleSize)]
	column := i.columns[getColumnFromIndex(cell.index, i.puzzleSize)]
	section := i.sections[getSectionFromIndex(cell.index, i.puzzleSize, i.sectionSize)]

	nextValue, err := i.findNextValue(row, column, section, i.minValue)
	if err != nil {
		return err
	}
	i.minValue = nextValue
	cell.value = nextValue
	return nil
}

func (i *iteration) findNextValue(row group, column group, section group, minValue int) (int, error) {
	usedMap := map[int]bool{}
	for _, usedValue := range row.usedValues() {
		if usedMap[usedValue] {
			continue
		}
		usedMap[usedValue] = true
	}
	for _, usedValue := range column.usedValues() {
		if usedMap[usedValue] {
			continue
		}
		usedMap[usedValue] = true
	}
	for _, usedValue := range section.usedValues() {
		if usedMap[usedValue] {
			continue
		}
		usedMap[usedValue] = true
	}

	for value := minValue; value <= i.puzzleSize; value++ {
		if usedMap[value] {
			continue
		}
		return value, nil
	}

	return 0, ErrNoMoreMoves
}
