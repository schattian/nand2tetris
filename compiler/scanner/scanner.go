package scanner

import (
	"unicode"

	"github.com/schattian/nand2tetris/compiler/token"
)

const eof rune = -1

type Scanner struct {
	src    []byte
	char   rune // current char
	offset int
}

func New(src []byte) *Scanner {
	s := &Scanner{src: src}
	s.init()
	return s
}

func (s *Scanner) init() {
	s.next()
}

func (s *Scanner) next() {
	if s.offset > len(s.src)-1 {
		s.char = eof
	} else {
		s.char = rune(s.src[s.offset])
	}
	s.offset += 1
}

func (s *Scanner) prev() {
	s.offset -= 1
	s.char = rune(s.src[s.offset-1])
}

func (s *Scanner) skipComments() {
	if s.char != '/' {
		return
	}

	s.next()
	if s.char != '/' && s.char != '*' {
		s.prev()
		return
	}

	if s.char == '*' {
		s.next()
		s.skipWildcardComment()
	}

	for s.char != '\n' && s.char != eof {
		s.next()
	}
}

func (s *Scanner) skipWildcardComment() {
	for s.char != '*' {
		s.next()
	}
	s.next()
	if s.char != '/' {
		s.skipWildcardComment()
	}
}

func (s *Scanner) Scan() (tok token.Token, lit string) {
	s.skipComments()
	tok, isLL1 := ll1Tokens[s.char]
	if isLL1 {
		lit = string(s.char)
		s.next()
	} else if unicode.IsDigit(s.char) {
		tok, lit = s.scanDigits()
	} else if unicode.IsSpace(s.char) {
		s.next()
		return s.Scan()
	} else if isStringLiteralStart(s.char) {
		tok, lit = s.scanStringLiteral()
	} else if isIdentStart(s.char) {
		tok, lit = s.scanIdentifier()
	} else {
		panic(string(s.char))
	}
	return
}

func (s *Scanner) scanStringLiteral() (tok token.Token, lit string) {
	s.next()
	tok = token.STRING_CONST
	for s.char != '"' {
		lit += string(s.char)
		s.next()
	}
	s.next()
	return
}

func (s *Scanner) scanIdentifier() (tok token.Token, lit string) {
	for isIdentBody(s.char) {
		lit += string(s.char)
		s.next()
	}
	if kwTok, isKw := llnTokens[lit]; isKw {
		tok = kwTok
	} else {
		tok = token.IDENT
	}
	return
}

func (s *Scanner) scanDigits() (tok token.Token, lit string) {
	tok = token.INTEGER_CONST
	for unicode.IsDigit(s.char) {
		lit += string(s.char)
		s.next()
	}
	return
}

func isStringLiteralStart(r rune) bool {
	return r == '"'
}

func isIdentStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isIdentBody(r rune) bool {
	return isIdentStart(r) || unicode.IsDigit(r)
}

var ll1Tokens = map[rune]token.Token{
	-1:  token.EOF,
	'(': token.LPAREN,
	'{': token.LBRACE,
	'[': token.LBRACK,
	',': token.COMMA,
	'.': token.DOT,
	')': token.RPAREN,
	'}': token.RBRACE,
	']': token.RBRACK,
	';': token.SEMICOLON,
	':': token.COLON,
	'+': token.ADD,
	'-': token.SUB,
	'*': token.MUL,
	'/': token.DIV,
	'&': token.AND,
	'|': token.OR,
	'~': token.NOT,
	'>': token.GT,
	'<': token.LT,
	'=': token.EQ,
}

var llnTokens = map[string]token.Token{
	"null":        token.NULL,
	"this":        token.THIS,
	"true":        token.TRUE,
	"false":       token.FALSE,
	"class":       token.CLASS,
	"constructor": token.CONSTRUCTOR,
	"function":    token.FUNCTION,
	"method":      token.METHOD,
	"field":       token.FIELD,
	"static":      token.STATIC,
	"var":         token.VAR,
	"int":         token.INT,
	"char":        token.CHAR,
	"boolean":     token.BOOLEAN,
	"void":        token.VOID,
	"let":         token.LET,
	"do":          token.DO,
	"if":          token.IF,
	"else":        token.ELSE,
	"while":       token.WHILE,
	"return":      token.RETURN,
}
