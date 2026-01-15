package transform

import "go-reloaded/texttools/tokenizer"

func ApplyPunctuation(tokens []tokenizer.Token) []tokenizer.Token {
	var result []tokenizer.Token
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok.Type == tokenizer.Punctuation {
			for i+1 < len(tokens) && tokens[i+1].Type == tokenizer.Punctuation {
				tok.Value += tokens[i+1].Value
				i++
			}
		}
		result = append(result, tok)
	}
	return result
}

