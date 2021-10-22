package parser

import (
	"encoding/json"

	"github.com/schattian/nand2tetris/compiler/parse"
	"github.com/schattian/nand2tetris/compiler/token"
)

type node struct {
	Child  []parse.Node `json:"child,omitempty"`
	Closed bool         `json:"closed,omitempty"`
	State  state        `json:"state,omitempty"`

	FieldsBySubset  [][]*field `json:"-"`
	LastFieldSubset int        `json:"last_field_subset"`

	Schema *nodeSchema `json:"schema,omitempty"`
}

func (n *nodeSchema) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}

func (n *node) String() string {
	s, _ := json.MarshalIndent(n, "", "  ")
	return string(s)
}

func (n *node) popChild() parse.Node {
	child := n.Child[len(n.Child)-1]
	n.Child = n.Child[0 : len(n.Child)-1]
	return child
}

type state int

func (n *node) NextNodeSchema(tok *parse.Token) *nodeSchema {
	return n.getSchemaByCtx()[tok.Token]
}

func (n *node) getSchemaByCtx() map[token.Token]*nodeSchema {
	return schemaByCtx[n.Schema.NodeType][n.State]
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
		2: {},
	},

	parse.NodeIfStatement: {
		0: {},
		1: matchExpression,
		2: {
			token.LET:    nodeLetStatement,
			token.IF:     nodeIfStatement,
			token.WHILE:  nodeWhileStatement,
			token.DO:     nodeDoStatement,
			token.RETURN: nodeReturnStatement,
		},
		3: {},
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
			token.TRUE:          nodeExpressionList,
			token.NULL:          nodeExpressionList,
			token.FALSE:         nodeExpressionList,
			token.THIS:          nodeExpressionList,
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
			token.DOT:    nodeSubroutineCall,
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
	if n.Closed || childNode == nil {
		return false
	}

	prevField := &field{}
	for _, field := range n.getFieldsBySubset()[n.LastFieldSubset] {
		isAdded = field.Add(childNode)
		if isAdded {
			if field.schema.nextState != 0 {
				n.State = field.schema.nextState
			}
			n.LastFieldSubset = field.schema.subset
			if field.schema.isSubsetCloser {
				n.LastFieldSubset += 1
			}
			if field.schema.isCloser {
				n.close()
			}
			break
		}
		if !field.IsSatisfied() {
			return false
		}
		prevField = field
		if prevField.schema != nil && prevField.schema.isChainer {
			if (prevField.valueCount-field.valueCount) > 1 || field.valueCount == 0 {
				break
			}
		}
	}

	if isAdded {
		n.Child = append(n.Child, childNode)
	} else {
		if len(n.getFieldsBySubset())-1 > n.LastFieldSubset {
			n.LastFieldSubset += 1
			return n.AddNode(childNode)
		}
	}

	return
}

func (n *node) getFieldsBySubset() (subsets [][]*field) {
	if n.FieldsBySubset != nil {
		return n.FieldsBySubset
	}

	var subset []*field
	var currentSubset int
	for _, fieldSchema := range n.Schema.FieldsSchema {
		if fieldSchema.subset != currentSubset {
			subsets = append(subsets, subset)
			subset = nil
		}
		subset = append(subset, &field{schema: fieldSchema})
		currentSubset = fieldSchema.subset
	}
	subsets = append(subsets, subset)
	n.FieldsBySubset = subsets
	return
}

func (n *node) close() {
	n.Closed = true
}

func (n *node) Token() *parse.Token {
	return n.Schema.Token
}
func (n *node) Type() parse.NodeType {
	return n.Schema.NodeType
}

func (n *node) Children() []parse.Node {
	return n.Child
}

type field struct {
	schema     *fieldSchema
	state      state
	valueCount int
	closed     bool
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
	f.valueCount += 1
	return true
}

func (f *field) IsSatisfied() bool {
	if f.closed {
		return true
	}
	return !f.schema.required
}
