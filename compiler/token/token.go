package token

type Token uint

func (t Token) String() string {
	return tokens[t]
}

func (t Token) IsLiteral() bool {
	return t > literals_start && t < literals_end
}

func (t Token) IsSymbol() bool {
	return t > symbols_start && t < symbols_end
}

func (t Token) IsKeyword() bool {
	return t > keywords_start && t < keywords_end
}

func (t Token) IsNativeType() bool {
	return t > types_start && t < types_end
}

func (t Token) IsType() bool {
	return t.IsNativeType() || t.IsIdentifier()
}

func (t Token) IsIdentifier() bool {
	return t == IDENT
}

func (t Token) IsOperator() bool {
	return t > operators_start && t < operators_end
}

func (t Token) IsBinaryOperator() bool {
	return t.IsOperator() && t != NOT
}

func IsUnaryOperator(t Token) bool {
	return t == NOT || t == SUB
}

func IsKeywordLiteral(t Token) bool {
	return t == NULL || t == FALSE || t == TRUE || t == THIS
}
