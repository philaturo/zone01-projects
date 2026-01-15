package transform

import (
	"strings"

	"go-reloaded/texttools/tokenizer"
)

func ApplyArticles(tokens []tokenizer.Token) []tokenizer.Token {
	for i := 0; i < len(tokens)-1; i++ {
		if tokens[i].Type == tokenizer.Word && strings.ToLower(tokens[i].Value) == "a" &&
			tokens[i+1].Type == tokenizer.Word {
			first := strings.ToLower(tokens[i+1].Value[:1])
			if strings.ContainsAny(first, "aeiouh") {
				tokens[i].Value = "an"
			}
		}
	}
	return tokens
}
