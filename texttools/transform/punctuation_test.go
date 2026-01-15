package transform

import (
	"testing"

	"go-reloaded/texttools/tokenizer"
)

func TestPunctuation(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"Hello , world !", "Hello, world!"},
		{"Wait ... what ?", "Wait... what?"},
		{"Yes ! No ?", "Yes! No?"},
	}

	for _, tt := range tests {
		tokens := tokenizer.Tokenize(tt.in)
		tokens = ApplyPunctuation(tokens)
		got := Reassemble(tokens)
		if got != tt.out {
			t.Errorf("input %q: expected %q, got %q", tt.in, tt.out, got)
		}
	}
}
