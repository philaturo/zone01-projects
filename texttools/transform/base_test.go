package transform

import (
	"testing"

	"go-reloaded/texttools/tokenizer"
)

func TestApplyHexBin(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"42 (hex)", "66"},
		{"10 (bin)", "2"},
		{"ZZ (hex)", "ZZ"}, // invalid hex ignored
	}

	for _, tt := range tests {
		tokens := tokenizer.Tokenize(tt.in)
		tokens = ApplyHexBin(tokens)
		got := Reassemble(tokens)

		if got != tt.out {
			t.Errorf("input %q: expected %q, got %q", tt.in, tt.out, got)
		}
	}
}
