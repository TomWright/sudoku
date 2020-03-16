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
		cells:       make(group, 0),
		rows:        make(groups, 0),
		columns:     make(groups, 0),
		sections:    make(groups, 0),
	}
	for i, item := range items {
		c := cell{
			fixed: item > 0,
			value: item,
			index: i,
		}
		it.cells = append(it.cells, &c)

		rowIndex := getRowFromIndex(c.index, it.puzzleSize)
		if len(it.rows) < rowIndex+1 {
			it.rows = append(it.rows, make(group, 0))
		}
		it.rows[rowIndex] = append(it.rows[rowIndex], &c)

		columnIndex := getColumnFromIndex(c.index, it.puzzleSize)
		if len(it.columns) < columnIndex+1 {
			it.columns = append(it.columns, make(group, 0))
		}
		it.columns[columnIndex] = append(it.columns[columnIndex], &c)

		sectionIndex := getSectionFromIndex(c.index, it.puzzleSize, it.sectionSize)
		if len(it.sections) < sectionIndex+1 {
			it.sections = append(it.sections, make(group, 0))
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
		cells:       make(group, 0),
		rows:        make(groups, 0),
		columns:     make(groups, 0),
		sections:    make(groups, 0),
		minValue:    1,
		index:       i.index,
	}
	for i, item := range i.cells {
		c := cell{
			fixed: item.fixed,
			value: item.value,
			index: i,
		}
		it.cells = append(it.cells, &c)

		rowIndex := getRowFromIndex(c.index, it.puzzleSize)
		if len(it.rows) < rowIndex+1 {
			it.rows = append(it.rows, make(group, 0))
		}
		it.rows[rowIndex] = append(it.rows[rowIndex], &c)

		columnIndex := getColumnFromIndex(c.index, it.puzzleSize)
		if len(it.columns) < columnIndex+1 {
			it.columns = append(it.columns, make(group, 0))
		}
		it.columns[columnIndex] = append(it.columns[columnIndex], &c)

		sectionIndex := getSectionFromIndex(c.index, it.puzzleSize, it.sectionSize)
		if len(it.sections) < sectionIndex+1 {
			it.sections = append(it.sections, make(group, 0))
		}
		it.sections[sectionIndex] = append(it.sections[sectionIndex], &c)
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
	usedValues := make([]int, 0)
	usedValues = append(usedValues, row.usedValues()...)
	usedValues = append(usedValues, column.usedValues()...)
	usedValues = append(usedValues, section.usedValues()...)

	isUsed := func(value int) bool {
		for _, usedValue := range usedValues {
			if value == usedValue {
				return true
			}
		}
		return false
	}

	for value := minValue; value <= i.puzzleSize; value++ {
		if isUsed(value) {
			continue
		}
		return value, nil
	}

	return 0, ErrNoMoreMoves
}
