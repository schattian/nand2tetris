package vm

type operation string

const (
	OpPush     operation = "push"
	OpPop      operation = "pop"
	OpAdd      operation = "add"
	OpSub      operation = "sub"
	OpNeg      operation = "neg"
	OpEq       operation = "eq"
	OpGt       operation = "gt"
	OpLt       operation = "lt"
	OpAnd      operation = "and"
	OpOr       operation = "or"
	OpNot      operation = "not"
	OpFunction operation = "function"
	OpCall     operation = "call"
	OpReturn   operation = "return"
	OpLabel    operation = "label"
	OpGoto     operation = "goto"
	OpIfGoto   operation = "if-goto"
)

type memSegment string

const (
	SegLcl     memSegment = "local"
	SegArg     memSegment = "argument"
	SegPointer memSegment = "pointer"
	SegStatic  memSegment = "static"
	SegTemp    memSegment = "temp"
	SegConst   memSegment = "constant"
	SegThis    memSegment = "this"
	SegThat    memSegment = "that"
)
