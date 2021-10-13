package parser

import (
	"github.com/schattian/nand2tetris/compiler/parse"
)

type node struct {
	children []parse.Node
	closed   bool

	fieldsBySubset  [][]*field
	lastFieldSubset int

	schema *nodeSchema
}

func (n *node) AddNode(childNode parse.Node) (isAdded bool) {
	if n.closed || childNode == nil {
		return false
	}

	for _, field := range n.getFieldsBySubset()[n.lastFieldSubset] {
		isAdded = field.Add(childNode)
		if isAdded {
			n.lastFieldSubset = field.schema.subset
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

func (f *field) IsSatisfied() bool {
	if f.closed {
		return true
	}
	return !f.schema.required
}
