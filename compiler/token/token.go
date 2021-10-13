package token

type Token uint

func (t Token) String() string {
	return tokens[t]
}

func IsLiteral(t Token) bool {
	return t > literals_start && t < literals_end
}

func IsSymbol(t Token) bool {
	return t > symbols_start && t < symbols_end
}

func IsKeyword(t Token) bool {
	return t > keywords_start && t < keywords_end
}

func IsNativeType(t Token) bool {
	return t > types_start && t < types_end
}

func IsType(t Token) bool {
	return IsNativeType(t) || IsIdentifier(t)
}

func IsIdentifier(t Token) bool {
	return t == IDENT
}

func IsOperator(t Token) bool {
	return t > operators_start && t < operators_end
}

func IsBinaryOperator(t Token) bool {
	return IsOperator(t) && t != NOT
}

func IsUnaryOperator(t Token) bool {
	return t == NOT || t == SUB
}

func IsKeywordLiteral(t Token) bool {
	return t == NULL || t == FALSE || t == TRUE || t == THIS
}
