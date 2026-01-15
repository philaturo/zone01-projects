package transform

import (
	"testing"

	"go-reloaded/texttools/tokenizer"
)

func TestQuotes(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"' hello '", "'hello'"},
		{"I am ' very cool ' today", "I am 'very cool' today"},
	}

	for _, tt := range tests {
		tokens := tokenizer.Tokenize(tt.in)
		tokens = ApplyQuotes(tokens)
		got := Reassemble(tokens)

		if got != tt.out {
			t.Errorf("expected %q, got %q", tt.out, got)
		}
	}
}
