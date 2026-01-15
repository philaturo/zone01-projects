package transform

import "strings"
import "go-reloaded/texttools/tokenizer"

func Reassemble(tokens []tokenizer.Token) string {
	var out strings.Builder
	for i, tok := range tokens {
		out.WriteString(tok.Value)
		if i < len(tokens)-1 {
			if tok.Type == tokenizer.Word && tokens[i+1].Type == tokenizer.Word {
				out.WriteString(" ")
			}
			if tok.Type == tokenizer.Punctuation && tokens[i+1].Type == tokenizer.Word {
				out.WriteString(" ")
			}
		}
	}
	return out.String()
}
