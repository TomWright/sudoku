package sudoku

import (
	"errors"
	"fmt"
	"math"
)

var (
	// ErrNoMoreMove is returned when there are no available moves on the current iteration.
	ErrNoMoreMoves = errors.New("no more moves")
	// ErrMissingIteration is returned when a required iteration is missing.
	ErrMissingIteration = errors.New("missing iteration")
)

// NewPuzzle returns a new puzzle.
func NewPuzzle(items []int) (*Puzzle, error) {
	puzzleSize := int(math.Sqrt(float64(len(items))))
	sectionSize := int(math.Sqrt(float64(puzzleSize)))
	firstIteration := newIteration(items, puzzleSize, sectionSize)
	return &Puzzle{
		puzzleSize: puzzleSize,
		iterations: []*iteration{
			firstIteration,
		},
		lastIteration:    nil,
		currentIteration: firstIteration,
	}, nil
}

// Puzzle is a sudoku puzzle.
type Puzzle struct {
	iterations       []*iteration
	lastIteration    *iteration
	currentIteration *iteration
	puzzleSize       int
}

// revert reverts back to the most recent applicable iteration.
func (p *Puzzle) revert() error {
	if p.lastIteration == nil {
		return ErrMissingIteration
	}

	// remove the last iteration.
	p.iterations = p.iterations[:len(p.iterations)-1]
	iterationsLen := len(p.iterations)
	if iterationsLen < 1 {
		return ErrMissingIteration
	}
	// set the current iteration and increment the min value.
	p.currentIteration = p.iterations[iterationsLen-1]
	p.currentIteration.minValue++

	// set the last iteration.
	if iterationsLen >= 2 {
		p.lastIteration = p.iterations[iterationsLen-2]
	}

	// if the min value is too high, revert again.
	if p.currentIteration.minValue > p.puzzleSize {
		return p.revert()
	}
	return nil
}

// next moves on to the next iteration.
func (p *Puzzle) next() error {
	nextIteration := p.currentIteration.iterate()
	p.lastIteration = p.currentIteration
	p.currentIteration = nextIteration
	p.iterations = append(p.iterations, p.currentIteration)

	return nil
}

// Solve solves the puzzle.
func (p *Puzzle) Solve() error {
	if err := p.next(); err != nil {
		return fmt.Errorf("initial iteration failed: %w", err)
	}
	p.currentIteration.index = 0
	for {
		err := p.currentIteration.solve()
		switch err {
		case nil:
			// we were able to find a matching value
			if p.currentIteration.finished() {
				// all cells have a value
				return nil
			}
			// continue to the next index.
			if err := p.next(); err != nil {
				return fmt.Errorf("could not move to next iteration: %w", err)
			}
			continue

		case ErrNoMoreMoves:
			// the iteration ran out of moves.
			// revert back to previous iteration with an incremented min value
			if err := p.revert(); err != nil {
				return fmt.Errorf("could not revert: %w", err)
			}
			continue

		default:
			return fmt.Errorf("could not solve iteration: %w", err)
		}
	}
}

// Result returns the last iteration.
func (p *Puzzle) Result() ([]int, error) {
	if p.currentIteration == nil {
		return nil, ErrMissingIteration
	}
	return p.currentIteration.items(), nil
}
