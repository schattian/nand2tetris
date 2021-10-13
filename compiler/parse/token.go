package parse

import "github.com/schattian/nand2tetris/compiler/token"

type Token struct {
	Token   token.Token
	Literal string
}

func NewToken(tok token.Token, lit string) *Token {
	return &Token{Literal: lit, Token: tok}
}

func (t *Token) Is(tok token.Token) bool {
	if t == nil {
		return false
	}
	return t.Token == tok
}
