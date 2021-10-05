package main

import (
	"errors"
)

const (
	SegLcl     VMMemSegment = "local"
	SegArg     VMMemSegment = "argument"
	SegPointer VMMemSegment = "pointer"
	SegStatic  VMMemSegment = "static"
	SegTemp    VMMemSegment = "temp"
	SegConst   VMMemSegment = "constant"
	SegThis    VMMemSegment = "this"
	SegThat    VMMemSegment = "that"

	OpPush VMOperation = "push"
	OpPop  VMOperation = "pop"
	OpAdd  VMOperation = "add"
	OpSub  VMOperation = "sub"
	OpNeg  VMOperation = "neg"
	OpEq   VMOperation = "eq"
	OpGt   VMOperation = "gt"
	OpLt   VMOperation = "lt"
	OpAnd  VMOperation = "and"
	OpOr   VMOperation = "or"
	OpNot  VMOperation = "not"

	OpFunction VMOperation = "function"
	OpCall     VMOperation = "call"
	OpReturn   VMOperation = "return"

	OpLabel  VMOperation = "label"
	OpGoto   VMOperation = "goto"
	OpIfGoto VMOperation = "if-goto"

	internalReg1 = "R13"
	internalReg2 = "R14"
	internalReg3 = "R15"

	ARegister Register = "A"
	MRegister Register = "M"
	DRegister Register = "D"

	initFunc string = "Sys.init"

	initSpValue uint16 = 256
)

var (
	segBaseAddress = map[VMMemSegment]uint16{
		SegTemp: 5,
	}

	segASMSymbol = map[VMMemSegment]string{
		SegLcl:  "LCL",
		SegArg:  "ARG",
		SegThis: "THIS",
		SegThat: "THAT",
	}
	pointerReferences = map[uint16]VMMemSegment{
		0: SegThis,
		1: SegThat,
	}
	errInvalidOperation = errors.New("invalid operation")

	initVMCommand = &callVMCommand{
		vmPc: 12121,
		funcName: initFunc,
	}
)

type Register string

type VMMemSegment string

type ASMSymbol string

func (seg VMMemSegment) ASMSymbol() string {
	return segASMSymbol[seg]
}

func (seg VMMemSegment) Dereference(i uint16) VMMemSegment {
	if !seg.IsPointer() {
		panic(seg + " is not a pointer memory segment")
	}
	return pointerReferences[i]
}

func (seg VMMemSegment) IsStateless() bool {
	return !seg.IsDynamic()
}

func (seg VMMemSegment) IsPointer() bool {
	return seg == SegPointer
}

func (seg VMMemSegment) IsVirtual() bool {
	return seg == SegConst
}

func (seg VMMemSegment) IsDynamic() bool {
	return map[VMMemSegment]bool{
		SegLcl:  true,
		SegArg:  true,
		SegThis: true,
		SegThat: true,
	}[seg]
}

func (seg VMMemSegment) IsStatic() bool {
	return seg == SegStatic
}

func (seg VMMemSegment) IsFixed() bool {
	return seg == SegTemp
}

func (seg VMMemSegment) BaseAddr() uint16 {
	return segBaseAddress[seg]
}

type VMOperation string

func (op VMOperation) IsMemoryAccess() bool {
	return op == OpPush || op == OpPop
}

func (op VMOperation) IsFlowControl() bool {
	return op == OpGoto || op == OpIfGoto || op == OpLabel
}

func (op VMOperation) IsCall() bool {
	return op == OpCall
}

func (op VMOperation) IsReturn() bool {
	return op == OpReturn
}

func (op VMOperation) IsFunction() bool {
	return op == OpFunction
}
