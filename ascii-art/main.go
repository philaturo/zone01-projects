package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/ascii"
)

func main() {
	// Expect exactly one argument
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	input = strings.ReplaceAll(input, `\n`, "\n")

	data, err := ascii.LoadBannerFile("banners/standard.txt")
	if err != nil {
		fmt.Println()
		return
	}

	banner, err := ascii.LoadBanner(data)
	if err != nil {
		fmt.Println()
		return
	}

	output := ascii.RenderText(input, banner)

	for i, line := range output {
		// Skip only the leading empty line
		if i == 0 && line == "" {
			continue
		}
		fmt.Println(line)
	}
}
