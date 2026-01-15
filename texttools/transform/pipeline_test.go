package transform

import (
	"testing"

	"go-reloaded/texttools/tokenizer"
)

func TestFullPipeline(t *testing.T) {
	input := `it (cap) was the best of times, it was the worst of times (up) , Simply add 42 (hex) and 10 (bin) , There is a amazing rock , Punctuation tests are ... kinda boring ,what do you think ? , ' I am the best ' , a apple and a banana`

	expected := `It was the best of times, it was the worst of TIMES, Simply add 66 and 2, There is an amazing rock, Punctuation tests are... kinda boring, what do you think?, 'I am the best', an apple and a banana`

	// Tokenize the input
	tokens := tokenizer.Tokenize(input)

	// Apply the full pipeline
	tokens = ApplyHexBin(tokens)
	tokens = ApplyCase(tokens)
	tokens = ApplyPunctuation(tokens)
	tokens = ApplyQuotes(tokens)
	tokens = ApplyArticles(tokens)

	// Reassemble into string
	got := Reassemble(tokens)

	if got != expected {
		t.Errorf("pipeline failed:\nexpected:\n%q\ngot:\n%q", expected, got)
	}
}

