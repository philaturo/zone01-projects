package tokenizer

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	Word TokenType = iota
	Marker
	Punctuation
	Quote
)

type Token struct {
	Value string
	Type  TokenType
}

// Tokenize splits text into words, punctuation, quotes, markers
func Tokenize(text string) []Token {
	var tokens []Token
	var current strings.Builder

	addToken := func() {
		if current.Len() > 0 {
			tokens = append(tokens, Token{Value: current.String(), Type: Word})
			current.Reset()
		}
	}

	for _, r := range text {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			current.WriteRune(r)
		case r == '(' || r == ')':
			addToken()
			tokens = append(tokens, Token{Value: string(r), Type: Marker})
		case strings.ContainsRune(".,!?;:", r):
			addToken()
			tokens = append(tokens, Token{Value: string(r), Type: Punctuation})
		case r == '\'':
			if current.Len() > 0 {
				current.WriteRune(r)
			}else {
				addToken()
			tokens = append(tokens, Token{Value: "'", Type: Quote})
			}
		case unicode.IsSpace(r):
			addToken()
		default:
			current.WriteRune(r)
		}
	}
	addToken()
	return tokens
}
