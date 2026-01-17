package transform

import (
	"strings"

	"go-reloaded/texttools/tokenizer"
)

func ApplyArticles(tokens []tokenizer.Token) []tokenizer.Token {
	for i := 0; i < len(tokens)-1; i++ {
		tok := tokens[i]

		if tok.Type != tokenizer.Word {
			continue
		}

		word := strings.ToLower(tok.Value)
		if word != "a" && word != "an" {
			continue
		}

		j := i + 1
		for j < len(tokens) && tokens[j].Type != tokenizer.Word {
			j++
		}
		if j >= len(tokens) {
			continue
		}

		next := strings.ToLower(tokens[j].Value)
		if len(next) == 0 {
			continue
		}

		start := next[0]
		isVowel := strings.ContainsRune("aeiou", rune(start))

		if isVowel && word == "a" {
			tokens[i].Value = "an"
		}
		if !isVowel && word == "an" {
			tokens[i].Value = "a"
		}
	}
	return tokens
}
