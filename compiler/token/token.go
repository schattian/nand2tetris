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
