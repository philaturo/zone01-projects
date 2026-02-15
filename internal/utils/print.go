// internal/utils/print.go
package utils

import "fmt"

// PrintBoard formats and prints the board with proper characters
func PrintBoard(board [][]rune) {
    for _, row := range board {
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

// BoardToString converts board to string for comparison
func BoardToString(board [][]rune) string {
    result := ""
    for _, row := range board {
        for _, cell := range row {
            if cell == 0 || cell == '.' {
                result += "."
            } else {
                result += string(cell)
            }
        }
        result += "\n"
    }
    return result
}