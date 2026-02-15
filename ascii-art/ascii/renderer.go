/*package ascii

import "strings"

func RenderText(text string, banner map[rune][]string) []string {
	// "" → print nothing
	if text == "" {
		return []string{}
	}

	lines := strings.Split(text, "\n")
	output := []string{}

	for _, line := range lines {
		// Empty logical line → blank line
		if line == "" {
			output = append(output, "")
			continue
		}

		// Render ASCII art
		for row := 0; row < charHeight; row++ {
			var sb strings.Builder
			for _, ch := range line {
				sb.WriteString(banner[ch][row])
			}
			output = append(output, sb.String())
		}
	}

	return output
}
*/
package ascii

import "strings"

func RenderText(text string, banner map[rune][]string) []string {
    if text == "" {
        return []string{}
    }

    lines := strings.Split(text, "\n")
    output := []string{}
    lastWidth := 0
    skipNextBlank := false

    for _, line := range lines {
        if line == "" {
            if lastWidth > 0 && !skipNextBlank {
                // Add only one blank line of spaces aligned with previous ASCII block
                output = append(output, strings.Repeat(" ", lastWidth))
                skipNextBlank = true // skip additional consecutive blanks
            } else {
                // skip extra blank lines
            }
            continue
        }

        skipNextBlank = false // reset for next non-empty line

        // Render ASCII characters for 8 lines
        rendered := make([]string, charHeight)
        for _, ch := range line {
            charArt, ok := banner[ch]
            if !ok {
                charArt = banner[' ']
            }
            for row := 0; row < charHeight; row++ {
                rendered[row] += charArt[row]
            }
        }

        output = append(output, rendered...)
        lastWidth = len(rendered[0])
    }

    return output
}
