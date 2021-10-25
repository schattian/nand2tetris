package parser

import (
	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/token"
)

var (
	nodeClass = &nodeSchema{
		NodeType: parse.NodeClass,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.CLASS),
			fieldIdentifier,
			fieldLBrace,
			{required: false, mustNodeType: parse.NodeClassVarDec, multiple: true},
			{required: false, mustNodeType: parse.NodeSubroutineDec, multiple: true},
			fieldRBraceCloser,
		},
	}

	nodeClassVarDec = &nodeSchema{
		NodeType: parse.NodeClassVarDec,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.STATIC, token.FIELD),
			fieldType,
			fieldIdentifier,
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},
			fieldSemicolon,
		},
	}

	nodeSubroutineDec = &nodeSchema{
		NodeType: parse.NodeSubroutineDec,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.CONSTRUCTOR, token.FUNCTION, token.METHOD),

			{required: false, mustTokenRule: token.IsType, subset: 1, isSubsetCloser: true},
			{required: false, mustOneOfTokens: []token.Token{token.VOID}, subset: 1, isSubsetCloser: true},

			{required: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 2, nextState: 1},
			{required: true, mustOneOfTokens: []token.Token{token.LPAREN}, subset: 2},
			{required: false, mustNodeType: parse.NodeParameterList, subset: 2},
			{required: true, mustOneOfTokens: []token.Token{token.RPAREN}, subset: 2},
			{required: true, mustNodeType: parse.NodeSubroutineBody, isCloser: true, subset: 2},
		},
	}

	nodeParameterList = &nodeSchema{
		NodeType: parse.NodeParameterList,
		FieldsSchema: []*fieldSchema{
			fieldType,
			fieldIdentifier,
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, multiple: true, mustTokenRule: token.IsType, subset: 1},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},
		},
	}

	nodeSubroutineBody = &nodeSchema{
		NodeType: parse.NodeSubroutineBody,
		FieldsSchema: []*fieldSchema{
			fieldLBrace,
			{required: false, multiple: true, mustNodeType: parse.NodeVarDec},
			fieldStatements,
			fieldRBraceCloser,
		},
	}

	nodeVarDec = &nodeSchema{
		NodeType: parse.NodeVarDec,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.VAR),
			fieldType,
			fieldIdentifier,
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.COMMA}, subset: 1},
			{required: false, multiple: true, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},
			fieldSemicolon,
		},
	}

	nodeLetStatement = &nodeSchema{
		NodeType: parse.NodeLetStatement,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.LET),
			{required: true, mustOneOfTokens: []token.Token{token.IDENT}, nextState: 1},

			{required: false, mustOneOfTokens: []token.Token{token.LBRACK}, subset: 1, isChainer: false},
			{required: false, mustNodeType: parse.NodeExpression, subset: 1, nextState: 2},
			{required: false, mustOneOfTokens: []token.Token{token.RBRACK}, subset: 1},

			{required: true, mustOneOfTokens: []token.Token{token.EQ}, subset: 2, nextState: 1},
			{required: true, mustNodeType: parse.NodeExpression, subset: 2},
			{required: true, mustOneOfTokens: []token.Token{token.SEMICOLON}, isCloser: true, subset: 2},
		},
	}

	nodeIfStatement = &nodeSchema{
		NodeType: parse.NodeIfStatement,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.IF),

			{required: true, mustOneOfTokens: []token.Token{token.LPAREN}, nextState: 1},
			{required: true, mustNodeType: parse.NodeExpression, nextState: 2},
			{required: true, mustOneOfTokens: []token.Token{token.RPAREN}},

			fieldLBrace,
			fieldStatements,
			fieldRBrace,

			{required: false, mustOneOfTokens: []token.Token{token.ELSE}, subset: 1, isChainer: true},
			{required: false, mustOneOfTokens: []token.Token{token.LBRACE}, subset: 1},
			{required: false, multiple: true, mustNodeTypeRule: isStatement, subset: 1},
			{required: false, mustOneOfTokens: []token.Token{token.RBRACE}, isCloser: true, subset: 1},
		},
	}

	nodeWhileStatement = &nodeSchema{
		NodeType: parse.NodeWhileStatement,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.WHILE),

			{required: true, mustOneOfTokens: []token.Token{token.LPAREN}, nextState: 1},
			{required: true, mustNodeType: parse.NodeExpression, nextState: 2},
			{required: true, mustOneOfTokens: []token.Token{token.RPAREN}},

			fieldLBrace,
			fieldStatements,
			fieldRBraceCloser,
		},
	}

	nodeDoStatement = &nodeSchema{
		NodeType: parse.NodeDoStatement,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.DO),
			fieldMustType(parse.NodeSubroutineCall),
			fieldSemicolon,
		},
	}

	nodeReturnStatement = &nodeSchema{
		NodeType: parse.NodeReturnStatement,
		FieldsSchema: []*fieldSchema{
			fieldMustTokens(token.RETURN),
			{required: false, mustNodeType: parse.NodeExpression},
			fieldSemicolon,
		},
	}

	nodeExpression = &nodeSchema{
		NodeType: parse.NodeExpression,
		FieldsSchema: []*fieldSchema{
			{required: true, mustNodeType: parse.NodeTerm, nextState: 1},
			{required: false, multiple: true, mustTokenRule: token.IsBinaryOperator, nextState: 2, subset: 1, isChainer: true},
			{required: false, multiple: true, mustNodeType: parse.NodeTerm, nextState: 1, subset: 1},
		},
	}

	nodeTerm = &nodeSchema{
		NodeType: parse.NodeTerm,
		FieldsSchema: []*fieldSchema{
			// unaryOp term
			{required: false, mustTokenRule: token.IsUnaryOperator, nextState: 1, isChainer: true},
			{required: false, mustNodeType: parse.NodeTerm, isCloser: true},

			// ( expression )
			{required: false, mustOneOfTokens: []token.Token{token.LPAREN}, subset: 1, nextState: 2, isChainer: true},
			{required: false, mustNodeType: parse.NodeExpression, subset: 1},
			{required: false, mustOneOfTokens: []token.Token{token.RPAREN}, isCloser: true, subset: 1},

			// ident
			{required: false, mustTokenRule: token.IsIdentifier, subset: 2, nextState: 4},

			// ident [ expression ]
			{required: false, mustOneOfTokens: []token.Token{token.LBRACK}, subset: 3, nextState: 3, isChainer: true},
			{required: false, mustNodeType: parse.NodeExpression, subset: 3},
			{required: false, mustOneOfTokens: []token.Token{token.RBRACK}, isCloser: true, subset: 3},

			// subroutineCall
			{required: false, mustNodeType: parse.NodeSubroutineCall, isCloser: true, subset: 4, nextState: 5},

			// literal
			{required: false, mustTokenRule: token.IsLiteral, isCloser: true, subset: 5},
		},
	}

	nodeSubroutineCall = &nodeSchema{
		NodeType: parse.NodeSubroutineCall,
		FieldsSchema: []*fieldSchema{
			fieldIdentifier,

			{required: false, mustOneOfTokens: []token.Token{token.DOT}, subset: 1, isChainer: true},
			{required: false, mustOneOfTokens: []token.Token{token.IDENT}, subset: 1},

			{required: true, mustOneOfTokens: []token.Token{token.LPAREN}, nextState: 1, subset: 2, isChainer: true},
			{required: false, mustNodeType: parse.NodeExpressionList, subset: 2},
			{required: true, mustOneOfTokens: []token.Token{token.RPAREN}, isCloser: true, subset: 2},
		},
	}
	nodeExpressionList = &nodeSchema{
		NodeType: parse.NodeExpressionList,
		FieldsSchema: []*fieldSchema{
			{required: false, mustNodeType: parse.NodeExpression, isChainer: true},
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
	fieldStatements = &fieldSchema{required: false, multiple: true, mustNodeTypeRule: isStatement}

	fieldSemicolon  = &fieldSchema{required: true, mustOneOfTokens: []token.Token{token.SEMICOLON}, isCloser: true}
	fieldIdentifier = fieldMustTokens(token.IDENT)
	fieldType       = fieldMustTokenRule(token.IsType)

	fieldLBrace       = fieldMustTokens(token.LBRACE)
	fieldRBrace       = fieldMustTokens(token.RBRACE)
	fieldRBraceCloser = &fieldSchema{required: true, mustOneOfTokens: []token.Token{token.RBRACE}, isCloser: true}
)

func isStatement(n parse.NodeType) bool {
	return n == parse.NodeIfStatement || n == parse.NodeLetStatement || n == parse.NodeWhileStatement || n == parse.NodeDoStatement || n == parse.NodeReturnStatement
}

type nodeSchema struct {
	FieldsSchema []*fieldSchema `json:"-"`
	Token        *parse.Token   `json:"token,omitempty"`
	NodeType     parse.NodeType `json:"node_type,omitempty"`
}

type fieldSchema struct {
	isChainer      bool
	subset         int
	required       bool
	multiple       bool
	isSubsetCloser bool
	isCloser       bool
	nextState      state

	mustTokenRule    func(t token.Token) bool
	mustNodeTypeRule func(t parse.NodeType) bool
	mustNodeType     parse.NodeType
	mustOneOfTokens  []token.Token
}

func (f *fieldSchema) validate(node parse.Node) bool {
	if f.mustNodeTypeRule != nil {
		return f.mustNodeTypeRule(node.Type())
	}
	if f.mustNodeType != parse.NodeIllegal {
		if node.Type() == f.mustNodeType {
			return true
		}
	}
	if node.Token() == nil {
		return false
	}
	if f.mustOneOfTokens != nil {
		for _, tok := range f.mustOneOfTokens {
			if node.Token().Token == tok {
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
	return &node{Schema: n}
}

func newTokenNode(token *parse.Token) *node {
	schema := &nodeSchema{Token: token, NodeType: parse.NodeToken}
	return schema.newNode()
}
