package sudoku

// groups is a set of groups.
type groups []group

// group is a group of cells.
type group []*cell

// usedValues returns all of the used values in the given group.
func (g group) usedValues() []int {
	values := make([]int, 0, len(g))
	for _, c := range g {
		if c.value > 0 {
			values = append(values, c.value)
		}
	}
	return values
}
