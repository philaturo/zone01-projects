# Tetris Optimizer

A powerful Go program that arranges tetromino pieces into the smallest possible square using backtracking algorithms.

## ✨ Features

- **File Parsing**: Reads tetromino configurations from text files
- **Validation**: Ensures tetrominoes are valid (4 connected blocks, proper format)
- **Optimization**: Finds the smallest square that can contain all pieces
- **Visual Output**: Displays the solution with letters (A, B, C, etc.) for each tetromino
- **Error Handling**: Gracefully handles invalid inputs with "ERROR" message

## Getting Started

### Prerequisites
- Go 1.21 or higher

### Installation
```bash
git clone https://github.com/philaturo/zone01-projects.git
cd zone01-projects/tetris-optimizer
```

## Usage

```bash
go run . [filename]
```

### Examples

```bash
$ cat sample.txt
#...
#...
#...
#...

....
....
..##
..##

$ go run . sample.txt
AA..
AA..
....
....

```

## Project structure

```bash
tetris-optimizer/
├── main.go                 # Entry point
├── internal/
│   ├── parser/            # File parsing & validation
│   │   └── parser.go
│   ├── solver/            # Backtracking algorithm
│   │   └── solver.go
│   └── utils/             # Helper functions
│       └── print.go
├── tests/
│   └── testdata/          # Test files
│       ├── good00.txt
│       ├── bad00.txt
│       └── hard.txt
├── go.mod
└── README.md
```

## Algorithm

The solver uses a backtracking approach:

    Parse & Validate: Read tetrominoes from file, ensure each has exactly 4 connected '#' blocks

    Calculate Minimum Size: Start with smallest possible square that can fit all pieces

    Recursive Placement: Try placing pieces one by one at every position

    Backtrack: If a piece doesn't fit, try the next position or increase board size

    Optimize: Continue until all pieces are placed successfully

## Test Coverage

The program includes comprehensive tests 

```bash
# Run all tests
go test -v ./...

# Run specific test
go test -v ./internal/parser -run TestParseFile_Valid

# Run with coverage
go test -cover ./...

```

## Error Handling

The program validates:

    File existence and readability

    Correct number of tetromino blocks

    Each block has exactly 4 lines

    Each line has exactly 4 characters

    Only '#' and '.' characters allowed

    Exactly 4 '#' per tetromino

    All '#' are connected (adjacent horizontally/vertically)

Any violation results in: ERROR

## Performance

    Efficient backtracking with pruning

    Starts with minimum possible board size

    Handles complex configurations

    Optimized for the hard example (8 tetrominoes)

