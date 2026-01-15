/*package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"go-reloaded/texttools/tokenizer"
	"go-reloaded/texttools/transform"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input> <output>")
		os.Exit(1)
	}

inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	text := string(data)

	// 1. Tokenize
	tokens := tokenizer.Tokenize(text)

	// 2. Transform
	tokens = transform.ApplyHexBin(tokens)
	tokens = transform.ApplyCase(tokens)
	tokens = transform.ApplyPunctuation(tokens)
	tokens = transform.ApplyQuotes(tokens)
	tokens = transform.ApplyArticles(tokens)

	// 3. Reassemble
	var output strings.Builder
	for i, t := range tokens {
		output.WriteString(t.Value)
		if i < len(tokens)-1 {
			if tokens[i+1].Type == tokenizer.Word {
				output.WriteString(" ")
			}
		}
	}

	// 4. Write output
	err = ioutil.WriteFile(outputFile, []byte(output.String()), 0644)
	if err != nil {
		fmt.Println("Error writing output:", err)
		os.Exit(1)
	}
}
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"go-reloaded/texttools/tokenizer"
	"go-reloaded/texttools/transform"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input> <output>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	text := string(data)

	// --- Pipeline ---
	tokens := tokenizer.Tokenize(text)
	tokens = transform.ApplyHexBin(tokens)
	tokens = transform.ApplyCase(tokens)
	tokens = transform.ApplyPunctuation(tokens)
	tokens = transform.ApplyQuotes(tokens)
	tokens = transform.ApplyArticles(tokens)

	output := transform.Reassemble(tokens)

	if err := ioutil.WriteFile(outputFile, []byte(output), 0644); err != nil {
		fmt.Println("Error writing output:", err)
		os.Exit(1)
	}
}
