// internal/solver/solver.go
package solver

import (
    "tetris-optimizer/internal/parser"
)

// Board represents the game board
type Board struct {
    Grid    [][]rune
    Size    int
}

// Solve finds the smallest square that fits all tetrominoes
func Solve(tetrominoes []parser.Tetromino) [][]rune {
    // Calculate minimum possible size (ceil(sqrt(4 * numPieces)))
    minSize := 1
    for minSize*minSize < 4*len(tetrominoes) {
        minSize++
    }

    // Try increasing board sizes until solution found
    for size := minSize; size <= minSize+10; size++ {
        board := &Board{
            Grid: make([][]rune, size),
            Size: size,
        }
        for i := range board.Grid {
            board.Grid[i] = make([]rune, size)
            for j := range board.Grid[i] {
                board.Grid[i][j] = '.'
            }
        }

        if backtrack(board, tetrominoes, 0) {
            return board.Grid
        }
    }
    return nil
}

// backtrack recursively tries to place tetrominoes
func backtrack(board *Board, tetrominoes []parser.Tetromino, index int) bool {
    if index == len(tetrominoes) {
        return true // All pieces placed
    }

    tetro := tetrominoes[index]

    // Try every possible position on the board
    for i := 0; i <= board.Size-tetro.Height; i++ {
        for j := 0; j <= board.Size-tetro.Width; j++ {
            if canPlace(board, tetro, i, j) {
                place(board, tetro, i, j, tetro.Letter)
                if backtrack(board, tetrominoes, index+1) {
                    return true
                }
                remove(board, tetro, i, j)
            }
        }
    }
    return false
}

// canPlace checks if tetromino fits at position (row, col)
func canPlace(board *Board, tetro parser.Tetromino, row, col int) bool {
    for _, coord := range tetro.Coords {
        r, c := row+coord[0], col+coord[1]
        if r < 0 || r >= board.Size || c < 0 || c >= board.Size {
            return false
        }
        if board.Grid[r][c] != '.' {
            return false
        }
    }
    return true
}

// place puts tetromino on board at position (row, col)
func place(board *Board, tetro parser.Tetromino, row, col int, letter rune) {
    for _, coord := range tetro.Coords {
        r, c := row+coord[0], col+coord[1]
        board.Grid[r][c] = letter
    }
}

// remove takes tetromino off board from position (row, col)
func remove(board *Board, tetro parser.Tetromino, row, col int) {
    for _, coord := range tetro.Coords {
        r, c := row+coord[0], col+coord[1]
        board.Grid[r][c] = '.'
    }
}