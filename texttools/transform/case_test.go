package transform

import (
	"testing"

	"go-reloaded/texttools/tokenizer"
)

func TestApplyCase(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"hello (up)", "HELLO"},
		{"HELLO (low)", "hello"},
		{"hello world (cap, 2)", "Hello World"},
		{"hello (cap, -2)", "Hello"}, // negative count defaults to 1
	}

	for _, tt := range tests {
		tokens := tokenizer.Tokenize(tt.in)
		tokens = ApplyCase(tokens)
		got := Reassemble(tokens)

		if got != tt.out {
			t.Errorf("input %q: expected %q, got %q", tt.in, tt.out, got)
		}
	}
}
