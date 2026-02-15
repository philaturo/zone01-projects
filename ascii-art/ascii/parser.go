package ascii

import (
	"fmt"
	"strings"
)

const (
	firstASCII = 32
	charHeight = 8
)

// LoadBanner parses banner data into a map of rune → ASCII art
func LoadBanner(data []byte) (map[rune][]string, error) {
	lines := strings.Split(string(data), "\n")

	expectedLines := (126-firstASCII+1)*(charHeight+1)
	if len(lines) < expectedLines {
		return nil, fmt.Errorf("invalid banner file format")
	}

	banner := make(map[rune][]string)
	index := 0

	for r := firstASCII; r <= 126; r++ {
		var charLines []string
		for i := 0; i < charHeight; i++ {
			charLines = append(charLines, lines[index])
			index++
		}
		banner[rune(r)] = charLines
		index++ // skip empty separator line
	}

	return banner, nil
}
