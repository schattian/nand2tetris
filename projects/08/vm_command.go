package main

import (
	"fmt"
)

type VMCommand interface {
	fmt.Stringer
	GetOp() VMOperation
	MarshalASM() (string, error)
}

type NewVMCommandInput struct {
	op        VMOperation
	seg       VMMemSegment
	funcName  string
	segIdx    uint16
	argSize   uint16
	localSize uint16
	vmPc      uint16
	labelName string
	module    string
	ctx       string
}

func NewVMCommand(input NewVMCommandInput) (VMCommand, error) {
	cmd, ok := map[VMOperation]VMCommand{
		OpPush:     &pushVMCommand{moduleName: input.module, seg: input.seg, segIdx: input.segIdx},
		OpPop:      &popVMCommand{moduleName: input.module, seg: input.seg, segIdx: input.segIdx},
		OpAdd:      &addVMCommand{},
		OpSub:      &subVMCommand{},
		OpNeg:      &negVMCommand{},
		OpEq:       &eqVMCommand{vmPc: input.vmPc},
		OpLt:       &ltVMCommand{vmPc: input.vmPc},
		OpGt:       &gtVMCommand{vmPc: input.vmPc},
		OpAnd:      &andVMCommand{},
		OpOr:       &orVMCommand{},
		OpNot:      &notVMCommand{},
		OpFunction: &functionVMCommand{funcName: input.funcName, localSize: input.localSize},
		OpCall:     &callVMCommand{funcName: input.funcName, argSize: input.argSize, vmPc: input.vmPc},
		OpReturn:   &returnVMCommand{},
		OpLabel:    &labelVMCommand{labelName: input.labelName, ctxName: input.ctx},
		OpGoto:     &gotoVMCommand{labelName: input.labelName, ctxName: input.ctx},
		OpIfGoto:   &ifGotoVMCommand{labelName: input.labelName, ctxName: input.ctx},
	}[input.op]
	if !ok {
		return nil, fmt.Errorf("op not supported: %s", input.op)
	}
	return cmd, nil
}
