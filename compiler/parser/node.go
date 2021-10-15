package parser

import (
	"fmt"

	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/token"
)

type node struct {
	children []parse.Node
	closed   bool

	state state

	fieldsBySubset  [][]*field
	lastFieldSubset int

	schema *nodeSchema
}

func (n *node) popChild() parse.Node {
	child := n.children[len(n.children)-1]
	n.children = n.children[0 : len(n.children)-1]
	return child
}

type state int

func (n *node) NextNodeSchema(tok *parse.Token) *nodeSchema {
	nextSchema, ok := n.getSchemaByCtx()[tok.Token]
	if !ok {
		return nil
		// return &nodeSchema{token: tok}
	}
	return nextSchema
	// for _, fields := range n.getFieldsBySubset() {
	// 	for _, field := range fields {
	// 		if field.closed {
	// 			continue
	// 		}
	// 		if nextSchema.nodeType == field.schema.mustNodeType {

	// 		}

	// 	}
	// }
	// return nil
}

func (n *node) getSchemaByCtx() map[token.Token]*nodeSchema {
	return schemaByCtx[n.schema.nodeType][n.state]
}

var schemaByCtx = map[parse.NodeType]map[state]map[token.Token]*nodeSchema{
	parse.NodeIllegal: {
		// class
		0: {
			token.CLASS: nodeClass,
		},
	},
	parse.NodeClass: {
		0: {
			// varDec
			token.STATIC: nodeClassVarDec,
			token.FIELD:  nodeClassVarDec,
			// subroutineDec
			token.CONSTRUCTOR: nodeSubroutineDec,
			token.FUNCTION:    nodeSubroutineDec,
			token.METHOD:      nodeSubroutineDec,
		},
	},

	parse.NodeClassVarDec: {
		0: {},
	},

	parse.NodeSubroutineDec: {
		0: {},
		1: {
			// paramList
			token.INT:     nodeParameterList,
			token.CHAR:    nodeParameterList,
			token.BOOLEAN: nodeParameterList,
			token.IDENT:   nodeParameterList,
			// subroutineBody
			token.LBRACE: nodeSubroutineBody,
		},
	},

	parse.NodeSubroutineBody: {
		0: {
			// varDec
			token.VAR: nodeVarDec,
			// statements
			token.LET:    nodeLetStatement,
			token.IF:     nodeIfStatement,
			token.WHILE:  nodeWhileStatement,
			token.DO:     nodeDoStatement,
			token.RETURN: nodeReturnStatement,
		},
	},

	parse.NodeLetStatement: {
		0: {},
		1: matchExpression,
	},

	// TODO: this breaks the mapping since we are using parse.Statement instead.
	parse.NodeIfStatement: {
		0: {},
		1: matchExpression,
		2: {
			token.LET:    nodeLetStatement,
			token.IF:     nodeIfStatement,
			token.WHILE:  nodeWhileStatement,
			token.DO:     nodeDoStatement,
			token.RETURN: nodeReturnStatement,
			//match statements
		},
	},

	parse.NodeWhileStatement: {
		0: {},
		1: matchExpression,
		2: {
			token.LET:    nodeLetStatement,
			token.IF:     nodeIfStatement,
			token.WHILE:  nodeWhileStatement,
			token.DO:     nodeDoStatement,
			token.RETURN: nodeReturnStatement,
			//match statements
		},
	},

	parse.NodeDoStatement: {
		0: {
			token.IDENT: nodeSubroutineCall,
		},
	},

	parse.NodeReturnStatement: {
		0: matchExpression,
	},

	parse.NodeExpression: {
		0: matchTerm,
		1: {},
		2: matchTerm,
	},

	parse.NodeExpressionList: {
		0: matchExpression,
	},

	parse.NodeSubroutineCall: {
		0: {},
		1: {
			token.IDENT:         nodeExpressionList,
			token.LPAREN:        nodeExpressionList,
			token.INTEGER_CONST: nodeExpressionList,
			token.STRING_CONST:  nodeExpressionList,
			token.NOT:           nodeExpressionList,
			token.SUB:           nodeExpressionList,
		},
	},

	parse.NodeTerm: {
		0: {},
		// unaryOp
		1: matchTerm,
		// isParen
		2: matchExpression,
		// isBrack
		3: matchExpression,
		// isIdent
		4: {
			token.LPAREN: nodeSubroutineCall,
		},
	},
}

var matchTerm = map[token.Token]*nodeSchema{
	token.IDENT: nodeTerm,

	token.LPAREN: nodeTerm,

	token.TRUE:  nodeTerm,
	token.NULL:  nodeTerm,
	token.FALSE: nodeTerm,
	token.THIS:  nodeTerm,

	token.INTEGER_CONST: nodeTerm,
	token.STRING_CONST:  nodeTerm,

	token.NOT: nodeTerm,
	token.SUB: nodeTerm,
}
var matchExpression = map[token.Token]*nodeSchema{
	token.IDENT: nodeExpression,

	token.TRUE:  nodeExpression,
	token.NULL:  nodeExpression,
	token.FALSE: nodeExpression,
	token.THIS:  nodeExpression,

	token.LPAREN: nodeExpression,

	token.INTEGER_CONST: nodeExpression,
	token.STRING_CONST:  nodeExpression,

	token.NOT: nodeExpression,
	token.SUB: nodeExpression,
}

func (n *node) AddNode(childNode parse.Node) (isAdded bool) {

	if n.closed || childNode == nil {
		return false
	}
	if childNode.Token() != nil {
		if childNode.Token().Token == token.RBRACE {
			fmt.Println()
		}
	}
	for _, field := range n.getFieldsBySubset()[n.lastFieldSubset] {
		isAdded = field.Add(childNode)
		if isAdded {
			if field.schema.nextState != 0 {
				// n.state += field.nextStateChange()
				n.state = field.schema.nextState
			}
			n.lastFieldSubset = field.schema.subset
			if field.schema.isSubsetCloser {
				n.lastFieldSubset += 1
			}
			if field.schema.isCloser {
				n.close()
			}
			break
		}
		if !field.IsSatisfied() {
			return false
		}

	}

	if isAdded {
		n.children = append(n.children, childNode)
	} else {
		if len(n.getFieldsBySubset())-1 > n.lastFieldSubset {
			n.lastFieldSubset += 1
			return n.AddNode(childNode)
		}
	}

	return
}

func (n *node) getFieldsBySubset() (subsets [][]*field) {
	if n.fieldsBySubset != nil {
		return n.fieldsBySubset
	}

	var subset []*field
	var currentSubset int
	for _, fieldSchema := range n.schema.fieldsSchema {
		if fieldSchema.subset != currentSubset {
			subsets = append(subsets, subset)
			subset = nil
		}
		subset = append(subset, &field{schema: fieldSchema})
		currentSubset = fieldSchema.subset
	}
	subsets = append(subsets, subset)
	n.fieldsBySubset = subsets
	return
}

func (n *node) close() {
	n.closed = true
}

func (n *node) Token() *parse.Token {
	return n.schema.token
}
func (n *node) Type() parse.NodeType {
	return n.schema.nodeType
}

func (n *node) Children() []parse.Node {
	return n.children
}

type field struct {
	schema *fieldSchema
	state  state
	closed bool
}

func (f *field) Add(node parse.Node) bool {
	if f.closed {
		return false
	}
	if v := f.schema.validate(node); !v {
		return false
	}
	if !f.schema.multiple {
		f.closed = true
	}
	return true
}

// func (f *field) nextStateChange() state {
// 	if f.state == 0 {
// 		f.state += 1
// 		return 1
// 	}
// 	f.state -= 1
// 	return -1
// }

func (f *field) IsSatisfied() bool {
	if f.closed {
		return true
	}
	return !f.schema.required
}
