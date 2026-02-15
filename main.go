// main.go
package main

import (
    "fmt"
    "os"
    "tetris-optimizer/internal/parser"
    "tetris-optimizer/internal/solver"
)

func main() {
    // Check command line arguments
    if len(os.Args) != 2 {
        fmt.Println("ERROR")
        return
    }

    // Parse tetrominoes from file
    tetrominoes, err := parser.ParseFile(os.Args[1])
    if err != nil {
        fmt.Println("ERROR")
        return
    }

    // Solve the puzzle
    solution := solver.Solve(tetrominoes)
    if solution == nil {
        fmt.Println("ERROR")
        return
    }

    // Print the solution properly
    for _, row := range solution {
        for _, cell := range row {
            if cell == 0 || cell == '.' {
                fmt.Print(".")
            } else {
                fmt.Print(string(cell))
            }
        }
        fmt.Println()
    }
}