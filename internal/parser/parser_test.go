// internal/parser/parser_test.go
package parser

import (
    "os"
    "path/filepath"
    "testing"
)

func createTestFile(t *testing.T, content string) string {
    t.Helper()
    tmpfile := filepath.Join(t.TempDir(), "test.txt")
    err := os.WriteFile(tmpfile, []byte(content), 0644)
    if err != nil {
        t.Fatal(err)
    }
    return tmpfile
}

func TestParseFile_Valid(t *testing.T) {
    content := `#...
#...
#...
#...

....
....
..##
..##`

    tmpfile := createTestFile(t, content)
    
    tetrominoes, err := ParseFile(tmpfile)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if len(tetrominoes) != 2 {
        t.Errorf("Expected 2 tetrominoes, got %d", len(tetrominoes))
    }
}

func TestParseFile_InvalidCharacter(t *testing.T) {
    content := `#...
#..X
#...
#...

....
....
..##
..##`

    tmpfile := createTestFile(t, content)
    
    _, err := ParseFile(tmpfile)
    if err == nil {
        t.Error("Expected error for invalid character, got nil")
    }
}

func TestParseFile_WrongHashCount(t *testing.T) {
    content := `#...
#...
#...
#..

....
....
..##
..##`

    tmpfile := createTestFile(t, content)
    
    _, err := ParseFile(tmpfile)
    if err == nil {
        t.Error("Expected error for wrong hash count, got nil")
    }
}

func TestParseFile_NotConnected(t *testing.T) {
    content := `#...
....
..#.
....

....
....
..##
..##`

    tmpfile := createTestFile(t, content)
    
    _, err := ParseFile(tmpfile)
    if err == nil {
        t.Error("Expected error for disconnected pieces, got nil")
    }
}

func TestParseFile_EmptyFile(t *testing.T) {
    content := ``
    tmpfile := createTestFile(t, content)
    
    _, err := ParseFile(tmpfile)
    if err == nil {
        t.Error("Expected error for empty file, got nil")
    }
}