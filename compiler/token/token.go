package token

type Token uint

func (t Token) Type() Type {
	return tokenTypes[t]
}

func (t Token) String() string {
	return tokens[t]
}

func init() {
	for tok := keywords_start + 1; tok < keywords_end; tok++ {
		tokenTypes[tok] = TypeKw
	}
	for tok := symbols_start + 1; tok < symbols_end; tok++ {
		tokenTypes[tok] = TypeSymbol
	}
}

func IsLiteral(t Token) bool {
	return t > literals_start && t < literals_end
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

var tokenTypes = map[Token]Type{
	NULL:          TypeKw,
	FALSE:         TypeKw,
	TRUE:          TypeKw,
	THIS:          TypeKw,
	IDENT:         TypeIdent,
	STRING_CONST:  TypeStrConst,
	INTEGER_CONST: TypeIntConst,
}

type Type string

const (
	TypeKw       Type = "keyword"
	TypeSymbol   Type = "symbol"
	TypeIntConst Type = "integerConstant"
	TypeStrConst Type = "stringConstant"
	TypeIdent    Type = "identifier"
)
