package transform

import (
	"strings"

	"go-reloaded/texttools/tokenizer"
)

func ApplyQuotes(tokens []tokenizer.Token) []tokenizer.Token {
	var result []tokenizer.Token
	i := 0
	for i < len(tokens) {
		if tokens[i].Type == tokenizer.Quote {
			start := i
			i++
			for i < len(tokens) && tokens[i].Type != tokenizer.Quote {
				i++
			}
			if i < len(tokens) {
				var val strings.Builder
				for j := start + 1; j < i; j++ {
					val.WriteString(tokens[j].Value)
					if j+1 < i && tokens[j+1].Type == tokenizer.Word {
						val.WriteString(" ")
					}
				}
				result = append(result, tokenizer.Token{Value: "'" + val.String() + "'", Type: tokenizer.Word})
			}
		} else {
			result = append(result, tokens[i])
		}
		i++
	}
	return result
}
