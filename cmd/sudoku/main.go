package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/tomwright/sudoku"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// MonitorCompletionRateInterval defines how often we should check the status of the puzzle.
const MonitorCompletionRateInterval = time.Millisecond * 200

func main() {
	in := flag.String("in", "", "File path to an input file containing the sudoku puzzle to solve")
	out := flag.String("out", "", "File path where the solved sudoku puzzle will be written")
	flag.Parse()

	if in == nil || *in == "" {
		_, _ = fmt.Fprintf(os.Stderr, "missing required -in argument\n")
		os.Exit(2)
	}
	if out == nil || *out == "" {
		_, _ = fmt.Fprintf(os.Stderr, "missing required -out argument\n")
		os.Exit(2)
	}

	input := getInput(*in)

	puzzle, err := sudoku.NewPuzzle(input)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create puzzle instance: %s\n", err)
		os.Exit(4)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	var bar *pb.ProgressBar
	{
		completionRate, err := puzzle.CompletionRate()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to get initial puzzle completion rate: %s\n", err)
			os.Exit(4)
		}

		// initialise the progress bar
		bar = pb.New(completionRate.TotalCells - completionRate.FixedCells)
	}

	_, _ = fmt.Fprintf(os.Stdout, "Solving puzzle...\n")
	// start the progress bar
	bar.Start()

	// solve the puzzle in a routine
	go solvePuzzle(wg, puzzle)
	// periodically fetch the puzzle completion rate and print it
	go monitorCompletionRate(wg, puzzle, bar)

	// wait for the routines to finish
	wg.Wait()

	// finish the progress bar
	bar.Finish()

	completionRate, err := puzzle.CompletionRate()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to get puzzle completion rate: %s\n", err)
		os.Exit(4)
	}

	switch {
	case completionRate.Completed:
		_, _ = fmt.Fprintf(os.Stderr, "Solved puzzle in %s\n", completionRate.CompletedAt.Sub(completionRate.StartedAt))
		break
	case completionRate.Failed:
		_, _ = fmt.Fprintf(os.Stderr, "Failed to solve puzzle: %v\n", completionRate.Error)
		os.Exit(4)
		return
	default:
		panic("unexpected completion rate status")
	}

	writeOutput(*out, puzzle)
}

func solvePuzzle(wg *sync.WaitGroup, p *sudoku.Puzzle) {
	defer wg.Done()
	_ = p.Solve()
}

func monitorCompletionRate(wg *sync.WaitGroup, p *sudoku.Puzzle, bar *pb.ProgressBar) {
	defer wg.Done()
	for {
		completionRate, err := p.CompletionRate()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "could not get completion rate: %s\n", err)
			os.Exit(3)
		}

		bar.SetCurrent(int64(completionRate.FilledCells - completionRate.FixedCells))

		switch {
		case completionRate.Completed:
			return
		case completionRate.Failed:
			return
		default:
			// still in progress.
			time.Sleep(MonitorCompletionRateInterval)
		}
	}
}

func getInput(path string) []int {
	// open input file
	inFile, err := os.Open(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot read input file: %s\n", err)
		os.Exit(3)
	}
	defer inFile.Close()

	input := make([]int, 0)

	// read unsolved puzzle from the file
	inScanner := bufio.NewScanner(inFile)
	for inScanner.Scan() {
		split := strings.Split(strings.TrimSpace(inScanner.Text()), " ")
		for _, s := range split {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			parsed, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "bad input: must only contain integers: %s\n", err)
				os.Exit(3)
			}
			input = append(input, int(parsed))
		}
	}
	if err := inScanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed when reading input file: %s\n", err)
		os.Exit(3)
	}
	if len(input) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "no input in file\n")
		os.Exit(3)
	}

	return input
}

func writeOutput(path string, p *sudoku.Puzzle) {
	// open output file
	outFile, err := os.Create(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot create output file: %s\n", err)
		os.Exit(5)
	}
	defer outFile.Close()

	results, err := p.Result()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot get puzzle results: %s\n", err)
		os.Exit(5)
	}

	formattedResults, err := sudoku.FormatPuzzle(results)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot format puzzle results: %s\n", err)
		os.Exit(5)
	}

	resultsString := ""
	for _, line := range formattedResults {
		for k, cell := range line {
			resultsString += fmt.Sprintf("%d", cell)
			if k != len(line)-1 {
				resultsString += " "
			}
		}
		resultsString += "\n"
	}

	_, err = outFile.WriteString(resultsString)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not write results to file: %s\n", err)
		os.Exit(5)
	}

	_, _ = fmt.Fprintf(os.Stdout, "Solved puzzle written to file: %s\n", path)
}
