package transform

import (
	"strconv"
	"strings"

	"go-reloaded/texttools/tokenizer"
)

func capitalizeWord(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func ApplyCase(tokens []tokenizer.Token) []tokenizer.Token {
	var result []tokenizer.Token
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if tok.Type == tokenizer.Marker && tok.Value == "(" &&
			i+1 < len(tokens) && tokens[i+1].Type == tokenizer.Word {

			cmd := strings.ToLower(tokens[i+1].Value)
			count := 1

			// detect (cmd, N)
			if i+3 < len(tokens) && tokens[i+2].Value == "," && tokens[i+3].Type == tokenizer.Word {
				if n, err := strconv.Atoi(tokens[i+3].Value); err == nil && n > 0 {
					count = n
				}
				i += 2 // skip ',' and number
			}

			// Apply transformation to previous N words
			applied := 0
			for j := len(result) - 1; j >= 0 && applied < count; j-- {
				if result[j].Type == tokenizer.Word {
					switch cmd {
					case "up":
						result[j].Value = strings.ToUpper(result[j].Value)
					case "low":
						result[j].Value = strings.ToLower(result[j].Value)
					case "cap":
						result[j].Value = capitalizeWord(result[j].Value)
					}
					applied++
				}
			}

			// Skip command and closing ')'
			i += 1
			if i+1 < len(tokens) && tokens[i+1].Value == ")" {
				i++
			}
			continue
		}

		result = append(result, tok)
	}
	return result
}
