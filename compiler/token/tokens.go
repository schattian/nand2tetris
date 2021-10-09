package token

const (
	EOF Token = iota

	// Keywords
	keywords_start
	CLASS
	CONSTRUCTOR
	FUNCTION
	METHOD
	FIELD
	STATIC
	VAR
	INT
	CHAR
	BOOLEAN
	VOID
	LET
	DO
	IF
	ELSE
	WHILE
	RETURN
	keywords_end

	// Symbols
	symbols_start
	// Delimiters - L
	LPAREN
	LBRACE
	LBRACK
	COMMA
	DOT
	// Delimiters - R
	RPAREN
	RBRACE
	RBRACK
	SEMICOLON
	COLON
	// Operators - Arithmetic
	ADD
	SUB
	MUL
	DIV
	// Operators - Logic
	AND
	OR
	NOT
	// Operators - Comparison
	GT
	LT
	EQ
	symbols_end

	// Literals
	literals_start
	TRUE
	FALSE
	NULL
	THIS
	INTEGER_CONST
	STRING_CONST
	IDENT
	literals_end
)

var tokens = [...]string{
	EOF: "<eof>",

	CLASS:       "class",
	CONSTRUCTOR: "constructor",
	FUNCTION:    "function",
	METHOD:      "method",
	FIELD:       "field",
	STATIC:      "static",
	VAR:         "var",
	INT:         "int",
	CHAR:        "char",
	BOOLEAN:     "boolean",
	VOID:        "void",

	LET:    "let",
	DO:     "do",
	IF:     "if",
	ELSE:   "else",
	WHILE:  "while",
	RETURN: "return",

	LPAREN:    "(",
	LBRACE:    "{",
	LBRACK:    "[",
	COMMA:     ",",
	DOT:       ".",
	RPAREN:    ")",
	RBRACE:    "}",
	RBRACK:    "]",
	SEMICOLON: ";",
	COLON:     ":",
	ADD:       "+",
	SUB:       "-",
	MUL:       "*",
	DIV:       "/",
	AND:       "&",
	OR:        "|",
	NOT:       "~",
	GT:        ">",
	LT:        "<",
	EQ:        "=",

	NULL:  "null",
	THIS:  "this",
	TRUE:  "true",
	FALSE: "false",

	INTEGER_CONST: "INT_LITERAL",
	STRING_CONST:  "STRING_LITERAL",
	IDENT:         "IDENT",
}