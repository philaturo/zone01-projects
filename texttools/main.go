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
	tokens = transform.ApplyArticles(tokens)
	tokens = transform.ApplyQuotes(tokens)
	

	output := transform.Reassemble(tokens)

	if err := ioutil.WriteFile(outputFile, []byte(output), 0644); err != nil {
		fmt.Println("Error writing output:", err)
		os.Exit(1)
	}
}
