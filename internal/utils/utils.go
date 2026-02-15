// internal/utils/utils.go
package utils

import (
    "fmt"
    "strings"
)

// PrintBoard formats and prints the board
func PrintBoard(board [][]rune) {
    for _, row := range board {
        for _, cell := range row {
            if cell == 0 {
                fmt.Print(".")
            } else {
                fmt.Print(string(cell))
            }
        }
        fmt.Println()
    }
}

// BoardToString converts board to string for testing
func BoardToString(board [][]rune) string {
    var result strings.Builder
    for _, row := range board {
        for _, cell := range row {
            if cell == 0 {
                result.WriteString(".")
            } else {
                result.WriteString(string(cell))
            }
        }
        result.WriteString("\n")
    }
    return result.String()
}