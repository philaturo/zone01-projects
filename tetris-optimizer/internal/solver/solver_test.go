// internal/solver/solver_test.go
package solver

import (
    "testing"
    "tetris-optimizer/internal/parser"
)

func TestSolve_Simple(t *testing.T) {
    // Create a simple I-shaped tetromino
    tetrominoes := []parser.Tetromino{
        {
            Shape:  [][]rune{{'#'}, {'#'}, {'#'}, {'#'}},
            Letter: 'A',
            Width:  1,
            Height: 4,
            Coords: [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
        },
    }

    solution := Solve(tetrominoes)
    if solution == nil {
        t.Error("Expected solution, got nil")
        return
    }

    // Count non-empty cells
    count := 0
    for i := 0; i < len(solution); i++ {
        for j := 0; j < len(solution[i]); j++ {
            if solution[i][j] != '.' {
                count++
            }
        }
    }
    if count != 4 {
        t.Errorf("Expected 4 cells filled, got %d", count)
    }
}

func TestSolve_TwoTetrominoes(t *testing.T) {
    tetrominoes := []parser.Tetromino{
        {
            Shape:  [][]rune{{'#'}, {'#'}, {'#'}, {'#'}},
            Letter: 'A',
            Width:  1,
            Height: 4,
            Coords: [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
        },
        {
            Shape:  [][]rune{{'#', '#'}, {'#', '#'}},
            Letter: 'B',
            Width:  2,
            Height: 2,
            Coords: [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
        },
    }

    solution := Solve(tetrominoes)
    if solution == nil {
        t.Error("Expected solution, got nil")
        return
    }

    // Count non-empty cells
    count := 0
    for i := 0; i < len(solution); i++ {
        for j := 0; j < len(solution[i]); j++ {
            if solution[i][j] != '.' {
                count++
            }
        }
    }
    if count != 8 {
        t.Errorf("Expected 8 cells filled, got %d", count)
    }
}