package transform

import (
	"testing"

	"go-reloaded/texttools/tokenizer"
)

func TestArticles(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"a apple", "an apple"},
		{"a house", "an house"},
		{"a banana", "a banana"},
	}

	for _, tt := range tests {
		tokens := tokenizer.Tokenize(tt.in)
		tokens = ApplyArticles(tokens)
		got := Reassemble(tokens)

		if got != tt.out {
			t.Errorf("expected %q, got %q", tt.out, got)
		}
	}
}
