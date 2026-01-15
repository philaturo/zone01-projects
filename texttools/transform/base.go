package transform

import (
	"strconv"

	"go-reloaded/texttools/tokenizer"
)

func ApplyHexBin(tokens []tokenizer.Token) []tokenizer.Token {
	var result []tokenizer.Token
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok.Type == tokenizer.Marker && tok.Value == "(" &&
			i+2 < len(tokens) &&
			tokens[i+1].Type == tokenizer.Word &&
			(tokens[i+1].Value == "hex" || tokens[i+1].Value == "bin") &&
			tokens[i+2].Type == tokenizer.Marker && tokens[i+2].Value == ")" {

			if len(result) > 0 {
				prev := &result[len(result)-1]
				switch tokens[i+1].Value {
				case "hex":
					if n, err := strconv.ParseInt(prev.Value, 16, 64); err == nil {
						prev.Value = strconv.FormatInt(n, 10)
					}
				case "bin":
					if n, err := strconv.ParseInt(prev.Value, 2, 64); err == nil {
						prev.Value = strconv.FormatInt(n, 10)
					}
				}
			}
			i += 2
			continue
		}

		result = append(result, tok)
	}
	return result
}
