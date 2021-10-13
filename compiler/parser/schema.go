package parser

import (
	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/token"
)

var (
	nodeClass = &nodeSchema{
		nodeType: parse.NodeClass,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.CLASS),
			fieldIdentifier,
			fieldLBrace,
			{required: false, mustNodeType: parse.NodeClassVarDec, multiple: true},
			{required: false, mustNodeType: parse.NodeSubroutineDec, multiple: true},
			fieldRBraceCloser,
		},
	}

	nodeClassVarDec = &nodeSchema{
		nodeType: parse.NodeClassVarDec,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.STATIC, token.FIELD),
			fieldType,
			fieldIdentifier,
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},
			fieldSemicolon,
		},
	}

	nodeSubroutineDec = &nodeSchema{
		nodeType: parse.NodeSubroutineDec,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.CONSTRUCTOR, token.FUNCTION, token.METHOD),
			fieldType,
			fieldIdentifier,
			{required: true, mustOneOfTokens: []token.Token{token.LPAREN}},
			{required: false, mustNodeType: parse.NodeParameterList},
			{required: true, mustOneOfTokens: []token.Token{token.RPAREN}},
			{required: true, mustNodeType: parse.NodeSubroutineBody, isCloser: true},
		},
	}

	nodeParameterList = &nodeSchema{
		nodeType: parse.NodeParameterList,
		fieldsSchema: []*fieldSchema{
			fieldType,
			fieldIdentifier,
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, mustTokenRule: token.IsType, subset: 1},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},
		},
	}

	nodeSubroutineBody = &nodeSchema{
		nodeType: parse.NodeSubroutineBody,
		fieldsSchema: []*fieldSchema{
			fieldLBrace,
			{required: false, multiple: true, mustNodeType: parse.NodeVarDec},
			fieldStatements,
			fieldRBraceCloser,
		},
	}

	nodeVarDec = &nodeSchema{
		nodeType: parse.NodeVarDec,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.VAR),
			fieldType,
			fieldIdentifier,
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},
			fieldSemicolon,
		},
	}

	nodeLetStatement = &nodeSchema{
		nodeType: parse.NodeStatement,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.LET),
			fieldIdentifier,
		},
	}

	nodeIfStatement = &nodeSchema{
		nodeType: parse.NodeStatement,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.IF),

			fieldLParen,
			fieldExpression,
			fieldRParen,

			fieldLBrace,
			fieldStatements,
			fieldRBrace,

			{required: false, mustOneOfTokens: []token.Token{token.ELSE}, subset: 1},
			{required: false, mustOneOfTokens: []token.Token{token.LBRACE}, subset: 1},
			{required: false, multiple: true, mustNodeType: parse.NodeStatement, subset: 1},
			{required: false, mustOneOfTokens: []token.Token{token.RBRACE}, subset: 1},
		},
	}

	nodeWhileStatement = &nodeSchema{
		nodeType: parse.NodeStatement,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.WHILE),

			fieldLParen,
			fieldExpression,
			fieldRParen,

			fieldLBrace,
			fieldStatements,
			fieldRBraceCloser,
		},
	}

	nodeDoStatement = &nodeSchema{
		nodeType: parse.NodeStatement,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.DO),
			fieldMustType(parse.NodeSubroutineCall),
			fieldSemicolon,
		},
	}

	nodeReturnStatement = &nodeSchema{
		nodeType: parse.NodeStatement,
		fieldsSchema: []*fieldSchema{
			fieldMustTokens(token.RETURN),
			{required: false, mustNodeType: parse.NodeExpression},
			fieldSemicolon,
		},
	}

	nodeExpression = &nodeSchema{
		nodeType: parse.NodeExpression,
		fieldsSchema: []*fieldSchema{
			{required: true, mustNodeType: parse.NodeTerm},
			{required: false, multiple: true, mustTokenRule: token.IsBinaryOperator, subset: 1},
			{required: false, multiple: true, mustNodeType: parse.NodeTerm, subset: 1},
		},
	}

	nodeTerm = &nodeSchema{
		nodeType: parse.NodeTerm,
		fieldsSchema: []*fieldSchema{
			// unaryOp term
			{required: false, mustTokenRule: token.IsUnaryOperator},
			{required: false, mustNodeType: parse.NodeTerm, isCloser: true},

			// ( expression )
			{required: false, mustOneOfTokens: []token.Token{token.LPAREN}, subset: 1},
			{required: false, mustNodeType: parse.NodeExpression, subset: 1},
			{required: false, mustOneOfTokens: []token.Token{token.RPAREN}, isCloser: true, subset: 1},

			// literal
			{required: false, mustTokenRule: token.IsLiteral, isCloser: true, subset: 2},

			// IDENT [ expression ]
			{required: false, mustTokenRule: token.IsIdentifier, subset: 3},
			{required: false, mustOneOfTokens: []token.Token{token.LBRACK}, subset: 3},
			{required: false, mustNodeType: parse.NodeExpression, subset: 3},
			{required: false, mustOneOfTokens: []token.Token{token.RBRACK}, isCloser: true, subset: 3},

			// literal
			{required: false, mustNodeType: parse.NodeSubroutineCall, isCloser: true, subset: 4},
		},
	}

	nodeExpressionList = &nodeSchema{
		nodeType: parse.NodeExpressionList,
		fieldsSchema: []*fieldSchema{
			{required: false, mustNodeType: parse.NodeExpression},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, multiple: true, mustNodeType: parse.NodeExpression, subset: 1},
		},
	}
)

func fieldMustType(nodeType parse.NodeType) *fieldSchema {
	return &fieldSchema{required: true, mustNodeType: nodeType}
}

func fieldMustTokens(tokens ...token.Token) *fieldSchema {
	return &fieldSchema{required: true, mustOneOfTokens: tokens}
}

func fieldMustTokenRule(tokenRule func(t token.Token) bool) *fieldSchema {
	return &fieldSchema{required: true, mustTokenRule: tokenRule}
}

var (
	fieldStatements = &fieldSchema{required: false, multiple: true, mustNodeType: parse.NodeStatement}
	fieldExpression = fieldMustType(parse.NodeExpression)

	fieldSemicolon  = &fieldSchema{required: true, mustOneOfTokens: []token.Token{token.SEMICOLON}, isCloser: true}
	fieldIdentifier = fieldMustTokens(token.IDENT)
	fieldType       = fieldMustTokenRule(token.IsType)

	fieldLBrace       = fieldMustTokens(token.LBRACE)
	fieldRBrace       = fieldMustTokens(token.RBRACE)
	fieldRBraceCloser = &fieldSchema{required: true, mustOneOfTokens: []token.Token{token.RBRACE}, isCloser: true}

	fieldLParen       = fieldMustTokens(token.LPAREN)
	fieldRParen       = fieldMustTokens(token.RPAREN)
	fieldRParenCloser = &fieldSchema{required: true, mustOneOfTokens: []token.Token{token.RPAREN}, isCloser: true}
)

type nodeSchema struct {
	fieldsSchema []*fieldSchema
	token        *parse.Token
	nodeType     parse.NodeType
}

type fieldSchema struct {
	subset   int
	required bool
	multiple bool
	isCloser bool

	mustTokenRule   func(t token.Token) bool
	mustNodeType    parse.NodeType
	mustOneOfTokens []token.Token
}

func (f *fieldSchema) validate(node parse.Node) bool {
	if f.mustNodeType != parse.NodeIllegal {
		if node.Type() == f.mustNodeType {
			return true
		}
	}
	if node.Token() == nil {
		return false
	}
	if f.mustOneOfTokens != nil {
		for _, token := range f.mustOneOfTokens {
			if node.Token().Token == token {
				return true
			}
		}
	}
	if f.mustTokenRule != nil {
		return f.mustTokenRule(node.Token().Token)
	}

	return false
}

func (n *nodeSchema) newNode() *node {
	return &node{schema: n}
}

func newTokenNode(token *parse.Token) *node {
	schema := &nodeSchema{token: token}
	return schema.newNode()
}