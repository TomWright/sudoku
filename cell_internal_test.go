package sudoku

import (
	"fmt"
	"testing"
)

func TestGetRowFromIndex(t *testing.T) {
	tests := []struct {
		Index int
		Size  int
		Out   int
	}{
		// row 0
		{Index: 0, Size: 9, Out: 0}, {Index: 1, Size: 9, Out: 0}, {Index: 2, Size: 9, Out: 0},
		{Index: 3, Size: 9, Out: 0}, {Index: 4, Size: 9, Out: 0}, {Index: 5, Size: 9, Out: 0},
		{Index: 6, Size: 9, Out: 0}, {Index: 7, Size: 9, Out: 0}, {Index: 8, Size: 9, Out: 0},
		// row 1
		{Index: 9, Size: 9, Out: 1}, {Index: 10, Size: 9, Out: 1}, {Index: 11, Size: 9, Out: 1},
		{Index: 12, Size: 9, Out: 1}, {Index: 13, Size: 9, Out: 1}, {Index: 14, Size: 9, Out: 1},
		{Index: 15, Size: 9, Out: 1}, {Index: 16, Size: 9, Out: 1}, {Index: 17, Size: 9, Out: 1},
		// row 2
		{Index: 18, Size: 9, Out: 2}, {Index: 19, Size: 9, Out: 2}, {Index: 20, Size: 9, Out: 2},
		{Index: 21, Size: 9, Out: 2}, {Index: 22, Size: 9, Out: 2}, {Index: 23, Size: 9, Out: 2},
		{Index: 24, Size: 9, Out: 2}, {Index: 25, Size: 9, Out: 2}, {Index: 26, Size: 9, Out: 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.Index), func(t *testing.T) {
			exp := tc.Out
			got := getRowFromIndex(tc.Index, tc.Size)
			if exp != got {
				t.Errorf("expected %d, got %d", exp, got)
			}
		})
	}
}

func TestGetColumnFromIndex(t *testing.T) {
	tests := []struct {
		Index int
		Size  int
		Out   int
	}{
		// col 0
		{Index: 0, Size: 9, Out: 0}, {Index: 9, Size: 9, Out: 0}, {Index: 18, Size: 9, Out: 0},
		// col 1
		{Index: 1, Size: 9, Out: 1}, {Index: 10, Size: 9, Out: 1}, {Index: 19, Size: 9, Out: 1},
		// col 2
		{Index: 2, Size: 9, Out: 2}, {Index: 11, Size: 9, Out: 2}, {Index: 20, Size: 9, Out: 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.Index), func(t *testing.T) {
			exp := tc.Out
			got := getColumnFromIndex(tc.Index, tc.Size)
			if exp != got {
				t.Errorf("expected %d, got %d", exp, got)
			}
		})
	}
}

func TestGetSectionRowFromIndex(t *testing.T) {
	tests := []struct {
		Index       int
		PuzzleSize  int
		SectionSize int
		Out         int
	}{
		// section 0
		{Index: 0, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 1, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 2, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 9, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 10, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 11, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 18, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 19, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 20, PuzzleSize: 9, SectionSize: 3, Out: 0},
		// section 1
		{Index: 3, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 4, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 5, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 12, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 13, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 14, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 21, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 22, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 23, PuzzleSize: 9, SectionSize: 3, Out: 0},
		// section 2
		{Index: 6, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 7, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 8, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 15, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 16, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 17, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 24, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 25, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 26, PuzzleSize: 9, SectionSize: 3, Out: 0},
		// section 3
		{Index: 27, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 28, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 29, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 36, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 37, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 38, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 45, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 46, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 47, PuzzleSize: 9, SectionSize: 3, Out: 1},
		// section 7
		{Index: 57, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 58, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 59, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 66, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 67, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 68, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 75, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 76, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 77, PuzzleSize: 9, SectionSize: 3, Out: 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.Index), func(t *testing.T) {
			exp := tc.Out
			got := getSectionRowFromIndex(tc.Index, tc.PuzzleSize, tc.SectionSize)
			if exp != got {
				t.Errorf("expected %d, got %d", exp, got)
			}
		})
	}
}

func TestGetSectionColumnFromIndex(t *testing.T) {
	tests := []struct {
		Index       int
		PuzzleSize  int
		SectionSize int
		Out         int
	}{
		// section 0
		{Index: 0, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 1, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 2, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 9, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 10, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 11, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 18, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 19, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 20, PuzzleSize: 9, SectionSize: 3, Out: 0},
		// section 1
		{Index: 3, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 4, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 5, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 12, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 13, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 14, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 21, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 22, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 23, PuzzleSize: 9, SectionSize: 3, Out: 1},
		// section 2
		{Index: 6, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 7, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 8, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 15, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 16, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 17, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 24, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 25, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 26, PuzzleSize: 9, SectionSize: 3, Out: 2},
		// section 3
		{Index: 27, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 28, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 29, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 36, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 37, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 38, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 45, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 46, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 47, PuzzleSize: 9, SectionSize: 3, Out: 0},
		// section 7
		{Index: 57, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 58, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 59, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 66, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 67, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 68, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 75, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 76, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 77, PuzzleSize: 9, SectionSize: 3, Out: 1},
		// section 8
		{Index: 60, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 61, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 62, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 69, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 70, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 71, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 78, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 79, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 80, PuzzleSize: 9, SectionSize: 3, Out: 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.Index), func(t *testing.T) {
			exp := tc.Out
			got := getSectionColumnFromIndex(tc.Index, tc.PuzzleSize, tc.SectionSize)
			if exp != got {
				t.Errorf("expected %d, got %d", exp, got)
			}
		})
	}
}

func TestGetSectionFromIndex(t *testing.T) {
	tests := []struct {
		Index       int
		PuzzleSize  int
		SectionSize int
		Out         int
	}{
		// section 0
		{Index: 0, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 1, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 2, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 9, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 10, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 11, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 18, PuzzleSize: 9, SectionSize: 3, Out: 0}, {Index: 19, PuzzleSize: 9, SectionSize: 3, Out: 0},
		{Index: 20, PuzzleSize: 9, SectionSize: 3, Out: 0},
		// section 1
		{Index: 3, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 4, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 5, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 12, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 13, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 14, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 21, PuzzleSize: 9, SectionSize: 3, Out: 1}, {Index: 22, PuzzleSize: 9, SectionSize: 3, Out: 1},
		{Index: 23, PuzzleSize: 9, SectionSize: 3, Out: 1},
		// section 2
		{Index: 6, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 7, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 8, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 15, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 16, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 17, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 24, PuzzleSize: 9, SectionSize: 3, Out: 2}, {Index: 25, PuzzleSize: 9, SectionSize: 3, Out: 2},
		{Index: 26, PuzzleSize: 9, SectionSize: 3, Out: 2},
		// section 3
		{Index: 27, PuzzleSize: 9, SectionSize: 3, Out: 3}, {Index: 28, PuzzleSize: 9, SectionSize: 3, Out: 3},
		{Index: 29, PuzzleSize: 9, SectionSize: 3, Out: 3}, {Index: 36, PuzzleSize: 9, SectionSize: 3, Out: 3},
		{Index: 37, PuzzleSize: 9, SectionSize: 3, Out: 3}, {Index: 38, PuzzleSize: 9, SectionSize: 3, Out: 3},
		{Index: 45, PuzzleSize: 9, SectionSize: 3, Out: 3}, {Index: 46, PuzzleSize: 9, SectionSize: 3, Out: 3},
		{Index: 47, PuzzleSize: 9, SectionSize: 3, Out: 3},
		// section 7
		{Index: 57, PuzzleSize: 9, SectionSize: 3, Out: 7}, {Index: 58, PuzzleSize: 9, SectionSize: 3, Out: 7},
		{Index: 59, PuzzleSize: 9, SectionSize: 3, Out: 7}, {Index: 66, PuzzleSize: 9, SectionSize: 3, Out: 7},
		{Index: 67, PuzzleSize: 9, SectionSize: 3, Out: 7}, {Index: 68, PuzzleSize: 9, SectionSize: 3, Out: 7},
		{Index: 75, PuzzleSize: 9, SectionSize: 3, Out: 7}, {Index: 76, PuzzleSize: 9, SectionSize: 3, Out: 7},
		{Index: 77, PuzzleSize: 9, SectionSize: 3, Out: 7},
		// section 8
		{Index: 60, PuzzleSize: 9, SectionSize: 3, Out: 8}, {Index: 61, PuzzleSize: 9, SectionSize: 3, Out: 8},
		{Index: 62, PuzzleSize: 9, SectionSize: 3, Out: 8}, {Index: 69, PuzzleSize: 9, SectionSize: 3, Out: 8},
		{Index: 70, PuzzleSize: 9, SectionSize: 3, Out: 8}, {Index: 71, PuzzleSize: 9, SectionSize: 3, Out: 8},
		{Index: 78, PuzzleSize: 9, SectionSize: 3, Out: 8}, {Index: 79, PuzzleSize: 9, SectionSize: 3, Out: 8},
		{Index: 80, PuzzleSize: 9, SectionSize: 3, Out: 8},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.Index), func(t *testing.T) {
			exp := tc.Out
			got := getSectionFromIndex(tc.Index, tc.PuzzleSize, tc.SectionSize)
			if exp != got {
				t.Errorf("expected %d, got %d", exp, got)
			}
		})
	}
}
