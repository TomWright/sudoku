package sudoku

import (
	"fmt"
	"reflect"
	"testing"
)

func benchmarkPuzzle(input []int) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			p, err := NewPuzzle(input)
			if err != nil {
				b.Errorf("could not create puzzle: %s", err)
				return
			}
			if err = p.Solve(); err != nil {
				b.Errorf("could not create puzzle: %s", err)
				return
			}
		}
	}
}

func BenchmarkPuzzle_Solve(b *testing.B) {
	b.Run("4x4", benchmarkPuzzle([]int{
		0, 0, 0, 3,
		0, 0, 0, 2,
		3, 0, 0, 0,
		4, 0, 0, 0,
	}))
	b.Run("9x9", benchmarkPuzzle([]int{
		6, 0, 0, 0, 0, 0, 1, 5, 0,
		9, 5, 4, 7, 1, 0, 0, 8, 0,
		0, 0, 0, 5, 0, 2, 6, 0, 0,
		8, 0, 0, 0, 9, 4, 0, 0, 6,
		0, 0, 3, 8, 0, 5, 4, 0, 0,
		4, 0, 0, 3, 7, 0, 0, 0, 8,
		0, 0, 6, 9, 0, 3, 0, 0, 0,
		0, 2, 0, 0, 4, 7, 8, 9, 3,
		0, 4, 9, 0, 0, 0, 0, 0, 5,
	}))
}

func ExamplePuzzle_Solve() {
	input := []int{
		0, 0, 0, 3,
		0, 0, 0, 2,
		3, 0, 0, 0,
		4, 0, 0, 0,
	}
	p, _ := NewPuzzle(input)
	_ = p.Solve()
	res, _ := p.Result()

	printItems := func(data []int) {
		formatted, _ := FormatPuzzle(data)
		for _, line := range formatted {
			for k, cell := range line {
				fmt.Printf("%d", cell)
				if k != len(line)-1 {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
	printItems(input)
	fmt.Println("-------")
	printItems(res)

	// Output:
	// 0 0 0 3
	// 0 0 0 2
	// 3 0 0 0
	// 4 0 0 0
	// -------
	// 2 4 1 3
	// 1 3 4 2
	// 3 1 2 4
	// 4 2 3 1
}

func TestPuzzle_Solve(t *testing.T) {
	run := func(in []int, exp []int) func(*testing.T) {
		return func(t *testing.T) {
			p, err := NewPuzzle(in)
			if err != nil {
				t.Errorf("could not create new puzzle: %s", err)
				return
			}

			if err := p.Solve(); err != nil {
				t.Errorf("could not solve puzzle: %s", err)
				return
			}

			got, err := p.Result()
			if err != nil {
				t.Errorf("could not get result: %s", err)
				return
			}

			if !reflect.DeepEqual(exp, got) {
				t.Errorf("expected %v, got %v", exp, got)
				return
			}
		}
	}

	type def struct {
		Exp []int
		In  []int
	}

	t.Run("4x4", func(t *testing.T) {
		tests := []def{
			{
				Exp: []int{
					2, 4, 1, 3,
					1, 3, 4, 2,
					3, 1, 2, 4,
					4, 2, 3, 1,
				},
				In: []int{
					0, 0, 0, 3,
					0, 0, 0, 2,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		}
		for k, tc := range tests {
			t.Run(fmt.Sprint(k), run(tc.In, tc.Exp))
		}
	})

	t.Run("9x9", func(t *testing.T) {
		tests := []def{
			{
				Exp: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 3, 4, 5, 8, 2, 1,
					2, 5, 1, 8, 7, 6, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 4, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
				In: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 3, 4, 5, 8, 2, 1,
					2, 5, 1, 8, 7, 6, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 0, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
			},
			{
				Exp: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 3, 4, 5, 8, 2, 1,
					2, 5, 1, 8, 7, 6, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 4, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
				In: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 3, 4, 5, 8, 2, 1,
					2, 5, 1, 8, 7, 0, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 0, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
			},
			{
				Exp: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 3, 4, 5, 8, 2, 1,
					2, 5, 1, 8, 7, 6, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 4, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
				In: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 0, 4, 5, 8, 2, 1,
					2, 5, 1, 0, 0, 0, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 0, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
			},
			{
				Exp: []int{
					4, 8, 3, 9, 2, 1, 6, 5, 7,
					9, 6, 7, 3, 4, 5, 8, 2, 1,
					2, 5, 1, 8, 7, 6, 4, 9, 3,
					5, 4, 8, 1, 3, 2, 9, 7, 6,
					7, 2, 9, 5, 6, 4, 1, 3, 8,
					1, 3, 6, 7, 9, 8, 2, 4, 5,
					3, 7, 2, 6, 8, 9, 5, 1, 4,
					8, 1, 4, 2, 5, 3, 7, 6, 9,
					6, 9, 5, 4, 1, 7, 3, 8, 2,
				},
				In: []int{
					0, 0, 3, 0, 2, 0, 6, 0, 0,
					9, 0, 0, 3, 0, 5, 0, 0, 1,
					0, 0, 1, 8, 0, 6, 4, 0, 0,
					0, 0, 8, 1, 0, 2, 9, 0, 0,
					7, 0, 0, 0, 0, 0, 0, 0, 8,
					0, 0, 6, 7, 0, 8, 2, 0, 0,
					0, 0, 2, 6, 0, 9, 5, 0, 0,
					8, 0, 0, 2, 0, 3, 0, 0, 9,
					0, 0, 5, 0, 1, 0, 3, 0, 0,
				},
			},
			{
				Exp: []int{
					6, 3, 2, 4, 8, 9, 1, 5, 7,
					9, 5, 4, 7, 1, 6, 3, 8, 2,
					1, 7, 8, 5, 3, 2, 6, 4, 9,
					8, 1, 7, 2, 9, 4, 5, 3, 6,
					2, 9, 3, 8, 6, 5, 4, 7, 1,
					4, 6, 5, 3, 7, 1, 9, 2, 8,
					7, 8, 6, 9, 5, 3, 2, 1, 4,
					5, 2, 1, 6, 4, 7, 8, 9, 3,
					3, 4, 9, 1, 2, 8, 7, 6, 5,
				},
				In: []int{
					6, 0, 0, 0, 0, 0, 1, 5, 0,
					9, 5, 4, 7, 1, 0, 0, 8, 0,
					0, 0, 0, 5, 0, 2, 6, 0, 0,
					8, 0, 0, 0, 9, 4, 0, 0, 6,
					0, 0, 3, 8, 0, 5, 4, 0, 0,
					4, 0, 0, 3, 7, 0, 0, 0, 8,
					0, 0, 6, 9, 0, 3, 0, 0, 0,
					0, 2, 0, 0, 4, 7, 8, 9, 3,
					0, 4, 9, 0, 0, 0, 0, 0, 5,
				},
			},
		}
		for k, tc := range tests {
			t.Run(fmt.Sprint(k), run(tc.In, tc.Exp))
		}
	})
}

func testCalculatePuzzleSize(input []int, exp int) func(t *testing.T) {
	return func(t *testing.T) {
		got, err := CalculatePuzzleSize(input)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			return
		}
		if exp != got {
			t.Errorf("expected %d, got %d", exp, got)
			return
		}
	}
}

func TestCalculatePuzzleSize(t *testing.T) {
	t.Run("4", testCalculatePuzzleSize([]int{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}, 4))
	t.Run("9", testCalculatePuzzleSize([]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 9))
	t.Run("16", testCalculatePuzzleSize([]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 16))
}

func testCalculateSectionSize(input []int, exp int) func(t *testing.T) {
	return func(t *testing.T) {
		got, err := CalculateSectionSize(input, 0)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			return
		}
		if exp != got {
			t.Errorf("expected %d, got %d", exp, got)
			return
		}
	}
}

func TestSectionPuzzleSize(t *testing.T) {
	t.Run("2", testCalculateSectionSize([]int{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}, 2))
	t.Run("3", testCalculateSectionSize([]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 3))
	t.Run("4", testCalculateSectionSize([]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 4))
}

func TestFormatPuzzle(t *testing.T) {
	input := []int{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}
	exp := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	got, err := FormatPuzzle(input)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("exp: %v, got %v", exp, got)
		return
	}
}
