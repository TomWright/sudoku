package sudoku

// cell is a single cell in the puzzle.
type cell struct {
	fixed bool
	value int
	index int
}

// copy returns a copy of the cell.
func (c *cell) copy() *cell {
	return &cell{
		fixed: c.fixed,
		value: c.value,
		index: c.index,
	}
}

// getRowFromIndex returns the row index for the given cell index.
func getRowFromIndex(index int, puzzleSize int) int {
	return index / puzzleSize
}

// getColumnFromIndex returns the column index for the given cell index.
func getColumnFromIndex(index int, puzzleSize int) int {
	return index % puzzleSize
}

// getSectionRowFromIndex returns the sections row index for the given cell index.
func getSectionRowFromIndex(index int, puzzleSize int, sectionSize int) int {
	rowI := getRowFromIndex(index, puzzleSize)
	return getRowFromIndex(rowI, sectionSize)
}

// getSectionColumnFromIndex returns the sections column index for the given cell index.
func getSectionColumnFromIndex(index int, puzzleSize int, sectionSize int) int {
	colI := getColumnFromIndex(index, puzzleSize)
	return getRowFromIndex(colI, sectionSize)
}

// getSectionFromIndex returns the index for the section from the given cell index.
func getSectionFromIndex(index int, puzzleSize int, sectionSize int) int {
	row := getSectionRowFromIndex(index, puzzleSize, sectionSize)
	column := getSectionColumnFromIndex(index, puzzleSize, sectionSize)
	return (row * sectionSize) + column
}
