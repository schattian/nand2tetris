package parse

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

type NodeType uint

const (
	NodeIllegal NodeType = iota

	NodeClass
	NodeClassVarDec
	NodeDataType
	NodeSubroutineDec
	NodeParameterList
	NodeSubroutineBody
	NodeVarDec
	NodeClassName
	NodeVarName

	NodeStatements
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
	NodeBinaryOp
	NodeUnaryOp
	NodeKeywordConstant

	NodeLeaf
)
