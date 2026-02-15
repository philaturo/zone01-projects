// internal/parser/parser.go
package parser

import (
   // "bufio"
    "fmt"
    "os"
    "strings"
)

// Tetromino represents a tetromino piece
type Tetromino struct {
    Shape   [][]rune
    Letter  rune
    Width   int
    Height  int
    Coords  [][2]int // Trimmed coordinates (minimal shape)
}

// internal/parser/parser.go - Update the ParseFile function
func ParseFile(filename string) ([]Tetromino, error) {
    // Read entire file first to handle different line endings
    content, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("cannot open file")
    }

    // Normalize line endings
    text := strings.ReplaceAll(string(content), "\r\n", "\n")
    text = strings.TrimSpace(text)
    
    // Split into tetromino blocks (separated by blank lines)
    blocks := strings.Split(text, "\n\n")
    
    if len(blocks) == 0 {
        return nil, fmt.Errorf("no tetrominoes found")
    }

    var tetrominoes []Tetromino
    letter := 'A'

    for _, block := range blocks {
        // Split block into lines
        lines := strings.Split(block, "\n")
        
        // Remove empty lines
        var cleanLines []string
        for _, line := range lines {
            if strings.TrimSpace(line) != "" {
                cleanLines = append(cleanLines, line)
            }
        }
        
        // Check if we have exactly 4 lines
        if len(cleanLines) != 4 {
            return nil, fmt.Errorf("invalid tetromino format: wrong number of lines")
        }

        tetro, err := parseBlock(cleanLines, letter)
        if err != nil {
            return nil, err
        }
        tetrominoes = append(tetrominoes, tetro)
        letter++
    }

    return tetrominoes, nil
}

// parseBlock validates and converts a 4x4 block into a Tetromino
func parseBlock(block []string, letter rune) (Tetromino, error) {
    // Check each line has exactly 4 characters
    for _, line := range block {
        if len(line) != 4 {
            return Tetromino{}, fmt.Errorf("invalid line length")
        }
        for _, ch := range line {
            if ch != '#' && ch != '.' {
                return Tetromino{}, fmt.Errorf("invalid character")
            }
        }
    }

    // Count number of '#' characters
    hashCount := 0
    for _, line := range block {
        for _, ch := range line {
            if ch == '#' {
                hashCount++
            }
        }
    }
    if hashCount != 4 {
        return Tetromino{}, fmt.Errorf("tetromino must have exactly 4 blocks")
    }

    // Check connectivity using BFS
    if !isConnected(block) {
        return Tetromino{}, fmt.Errorf("tetromino blocks are not connected")
    }

    // Create shape matrix
    shape := make([][]rune, 4)
    for i := 0; i < 4; i++ {
        shape[i] = []rune(block[i])
    }

    // Trim the tetromino to its minimal shape
    trimmed, width, height := trimShape(shape)

    return Tetromino{
        Shape:   trimmed,
        Letter:  letter,
        Width:   width,
        Height:  height,
        Coords:  extractCoordinates(trimmed),
    }, nil
}

// isConnected checks if all '#' are connected via adjacency (up/down/left/right)
func isConnected(block []string) bool {
    // Find first '#'
    startI, startJ := -1, -1
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if block[i][j] == '#' {
                startI, startJ = i, j
                break
            }
        }
        if startI != -1 {
            break
        }
    }

    if startI == -1 {
        return false // No '#' found
    }

    // BFS to count connected '#'
    visited := make([][]bool, 4)
    for i := range visited {
        visited[i] = make([]bool, 4)
    }

    queue := [][2]int{{startI, startJ}}
    visited[startI][startJ] = true
    count := 1
    directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        for _, dir := range directions {
            ni, nj := current[0]+dir[0], current[1]+dir[1]
            if ni >= 0 && ni < 4 && nj >= 0 && nj < 4 && 
               !visited[ni][nj] && block[ni][nj] == '#' {
                visited[ni][nj] = true
                queue = append(queue, [2]int{ni, nj})
                count++
            }
        }
    }

    return count == 4
}

// trimShape removes empty rows and columns from the edges
func trimShape(shape [][]rune) ([][]rune, int, int) {
    // Find min and max rows and columns that contain '#'
    minRow, maxRow := 3, 0
    minCol, maxCol := 3, 0
    found := false

    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if shape[i][j] == '#' {
                found = true
                if i < minRow {
                    minRow = i
                }
                if i > maxRow {
                    maxRow = i
                }
                if j < minCol {
                    minCol = j
                }
                if j > maxCol {
                    maxCol = j
                }
            }
        }
    }

    if !found {
        return shape, 0, 0
    }

    // Create trimmed shape
    height := maxRow - minRow + 1
    width := maxCol - minCol + 1
    trimmed := make([][]rune, height)
    for i := range trimmed {
        trimmed[i] = make([]rune, width)
        for j := range trimmed[i] {
            trimmed[i][j] = shape[minRow+i][minCol+j]
        }
    }

    return trimmed, width, height
}

// extractCoordinates gets the coordinates of '#' within the trimmed shape
func extractCoordinates(shape [][]rune) [][2]int {
    var coords [][2]int
    for i := 0; i < len(shape); i++ {
        for j := 0; j < len(shape[i]); j++ {
            if shape[i][j] == '#' {
                coords = append(coords, [2]int{i, j})
            }
        }
    }
    return coords
}