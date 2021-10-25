package parse

import (
	"encoding/xml"
)

type Tree struct {
	Name string
	Root Node
}

type Node interface {
	Type() NodeType
	AddNode(node Node) bool
	Children() []Node
	Token() *Token
}

func (tree *Tree) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalXMLNode(e, start, tree.Root)
}

func marshalXMLNode(e *xml.Encoder, start xml.StartElement, node Node) error {
	elemTypeName := xml.Name{Local: node.Type().String()}
	if node.Token() != nil {
		elemTypeName.Local = string(node.Token().Token.Type())
	}
	if elemTypeName.Local != "" {
		err := e.EncodeToken(xml.StartElement{Name: elemTypeName})
		if err != nil {
			return err
		}
	}
	for _, child := range node.Children() {
		err := marshalXMLNode(e, start, child)
		if err != nil {
			return err
		}
	}
	if node.Token() != nil {
		err := e.EncodeToken(xml.CharData(" " + node.Token().Literal + " "))
		if err != nil {
			return err
		}
	}
	if elemTypeName.Local != "" {
		err := e.EncodeToken(xml.EndElement{Name: elemTypeName})
		if err != nil {
			return err
		}
	}
	return nil
}

type NodeType uint

const (
	NodeIllegal NodeType = iota

	NodeClass
	NodeClassVarDec
	NodeSubroutineDec
	NodeParameterList
	NodeSubroutineBody
	NodeVarDec

	NodeStatement
	NodeLetStatement
	NodeIfStatement
	NodeWhileStatement
	NodeDoStatement
	NodeReturnStatement

	NodeExpression
	NodeTerm
	NodeSubroutineCall
	NodeExpressionList

	NodeToken
)

var nodeTypeTemplateNames = map[NodeType]string{
	NodeClass:           "class",
	NodeClassVarDec:     "classVarDec",
	NodeSubroutineDec:   "subroutineDec",
	NodeParameterList:   "parameterList",
	NodeSubroutineBody:  "subroutineBody",
	NodeVarDec:          "varDec",
	NodeLetStatement:    "letStatement",
	NodeIfStatement:     "ifStatement",
	NodeWhileStatement:  "whileStatement",
	NodeDoStatement:     "doStatement",
	NodeReturnStatement: "returnStatement",
	NodeExpression:      "expression",
	NodeTerm:            "term",
	NodeExpressionList:  "expressionList",
	// NodeSubroutineCall:  "subroutineCall",
}

func (nt NodeType) String() string {
	return nodeTypeTemplateNames[nt]
}
