# sudoku

[![Go Report Card](https://goreportcard.com/badge/github.com/TomWright/sudoku)](https://goreportcard.com/report/github.com/TomWright/sudoku)
[![Documentation](https://godoc.org/github.com/TomWright/sudoku?status.svg)](https://godoc.org/github.com/TomWright/sudoku)
![Test](https://github.com/TomWright/sudoku/workflows/Test/badge.svg)
![Build](https://github.com/TomWright/sudoku/workflows/Build/badge.svg)

Automatically solve sudoku puzzles of any size.

## Usage

Download an executable from the [latest release](https://github.com/TomWright/sudoku/releases/latest), or build locally.

You may have to `chmod +x` the download file.

Write our unsolved puzzle to a file:
```
echo "6 0 0 0 0 0 1 5 0
9 5 4 7 1 0 0 8 0
0 0 0 5 0 2 6 0 0
8 0 0 0 9 4 0 0 6
0 0 3 8 0 5 4 0 0
4 0 0 3 7 0 0 0 8
0 0 6 9 0 3 0 0 0
0 2 0 0 4 7 8 9 3
0 4 9 0 0 0 0 0 5" > unsolved_puzzle.txt
```

Run the `sudoku` command:
```
sudoku -in unsolved_puzzle.txt -out solved_puzzle.txt
```

View the solved puzzle:
```
cat solved_puzzle.txt
```

Which gives you:
```
6 3 2 4 8 9 1 5 7
9 5 4 7 1 6 3 8 2
1 7 8 5 3 2 6 4 9
8 1 7 2 9 4 5 3 6
2 9 3 8 6 5 4 7 1
4 6 5 3 7 1 9 2 8
7 8 6 9 5 3 2 1 4
5 2 1 6 4 7 8 9 3
3 4 9 1 2 8 7 6 5
```

## Puzzle requirements

Puzzle sizes must be correct otherwise you may get unexpected results.

A puzzle should have a size of:
- 2x2
- 3x3
- 4x4
- 5x5
- etc

See the examples on godoc for more information.
