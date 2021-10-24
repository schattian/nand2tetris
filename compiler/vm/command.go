package vm

import "fmt"

type command struct {
	op  operation
	seg memSegment

	arg uint16

	argSz   uint16
	localSz uint16

	label  string
	fnName string
}

func (c *command) String() (s string) {
	switch c.op {
	case OpPush, OpPop:
		s += fmt.Sprintf("%s %s %d", c.op, c.seg, c.arg)
	case OpAdd, OpSub, OpNeg, OpEq, OpGt, OpLt, OpAnd, OpOr, OpNot:
		s += string(c.op)
	case OpLabel, OpGoto, OpIfGoto:
		s += fmt.Sprintf("%s %s", c.op, c.label)
	case OpCall:
		s += fmt.Sprintf("%s %s", c.op, c.fnName)
	case OpFunction:
		s += fmt.Sprintf("%s %s %d", c.op, c.fnName, c.localSz)
	case OpReturn:
		s += fmt.Sprintf("%s", c.op)
	}
	s += "\n"
	return
}

func NewAccessCommand(op operation, seg memSegment, index uint16) *command {
	return &command{seg: seg, arg: index, op: op}
}

func NewArithmeticCommand(op operation) *command {
	return &command{op: op}
}

func NewFlowControlCommand(op operation, label string) *command {
	return &command{op: op, label: label}
}

func NewCallCommand(fnName string, argSz uint16) *command {
	return &command{op: OpCall, argSz: argSz, fnName: fnName}
}

func NewFunctionCommand(name string, localSz uint16) *command {
	return &command{fnName: name, localSz: localSz}
}

func NewReturnCommand() *command {
	return &command{op: OpReturn}
}
