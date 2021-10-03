package main

import (
	"fmt"
	"path/filepath"
)

type VMCommand interface {
	fmt.Stringer
	GetOp() VMOperation
	MarshalASM() (string, error)
}

func NewVMCommand(absModuleName string, op VMOperation, seg VMMemSegment, segIdx, vmPc uint16) (VMCommand, error) {
	cmd, ok := map[VMOperation]VMCommand{
		OpPush: &pushVMCommand{moduleName: filepath.Base(absModuleName), seg: seg, segIdx: segIdx},
		OpPop:  &popVMCommand{moduleName: filepath.Base(absModuleName), seg: seg, segIdx: segIdx},
		OpAdd:  &addVMCommand{},
		OpSub:  &subVMCommand{},
		OpNeg:  &negVMCommand{},
		OpEq:   &eqVMCommand{vmPc: vmPc},
		OpLt:   &ltVMCommand{vmPc: vmPc},
		OpGt:   &gtVMCommand{vmPc: vmPc},
		OpAnd:  &andVMCommand{},
		OpOr:   &orVMCommand{},
		OpNot:  &notVMCommand{},
	}[op]
	if !ok {
		return nil, nil
		//return nil, fmt.Errorf("op not supported: %s", op)
	}
	return cmd, nil
}

type pushVMCommand struct {
	moduleName string
	seg        VMMemSegment
	segIdx     uint16
}

func (cmd *pushVMCommand) GetOp() VMOperation {
	return OpPush
}

func (cmd *pushVMCommand) MarshalASM() (s string, err error) {
	if cmd.seg.IsVirtual() {
		s += translateAssignConstantD(cmd.segIdx)
	}
	if cmd.seg.IsDynamic() {
		s += translateGetDynamicAddr(cmd.seg, cmd.segIdx, ARegister)
	}
	if cmd.seg.IsPointer() {
		s += translateGetPointerAddr(cmd.seg, cmd.segIdx)
	}
	if cmd.seg.IsFixed() {
		s += translateGetFixedAddr(cmd.seg, cmd.segIdx)
	}
	if cmd.seg.IsStatic() {
		s += translateGetStaticAddr(cmd.moduleName, cmd.segIdx)
	}
	if !cmd.seg.IsVirtual() {
		s += "D=M\n"
	}
	s += `@SP
A=M
M=D
@SP
M=M+1
`
	return
}

func (cmd *pushVMCommand) String() string {
	return fmt.Sprintf("%s %s %d", cmd.GetOp(), cmd.seg, cmd.segIdx)
}

type popVMCommand struct {
	moduleName string
	seg        VMMemSegment
	segIdx     uint16
}

func (cmd *popVMCommand) MarshalASM() (s string, err error) {
	if cmd.seg.IsVirtual() {
		return "", errInvalidOperation
	}
	if cmd.seg.IsStateless() {
		s = cmd.marshalASMStateless()
	} else {
		s = cmd.marshalASMStateful()
	}
	s += "M=D\n"
	return
}
func (cmd *popVMCommand) marshalASMStateful() (s string) {
	s += translateGetDynamicAddr(cmd.seg, cmd.segIdx, DRegister)
	s += fmt.Sprintf(`@%s
M=D
`, internalReg1)
	s += translatePopToD()
	s += fmt.Sprintf(`@%s
A=M
`, internalReg1)
	return
}

func (cmd *popVMCommand) marshalASMStateless() (s string) {
	s += translatePopToD()
	if cmd.seg.IsPointer() { // TODO: Add fixed
		s += translateGetPointerAddr(cmd.seg, cmd.segIdx)
	} else if cmd.seg.IsStatic() {
		s += translateGetStaticAddr(cmd.moduleName, cmd.segIdx)
	} else if cmd.seg.IsFixed() {
		s += translateGetFixedAddr(cmd.seg, cmd.segIdx)
	}
	return
}

func translatePopToD() string {
	return `@SP
AM=M-1
D=M
`
}

func (cmd *popVMCommand) String() string {
	return fmt.Sprintf("%s %s %d", cmd.GetOp(), cmd.seg, cmd.segIdx)
}

func (cmd *popVMCommand) GetOp() VMOperation {
	return OpPop
}

type addVMCommand struct{}

func (cmd *addVMCommand) GetOp() VMOperation {
	return OpAdd
}

func (cmd *addVMCommand) MarshalASM() (s string, err error) {
	return `@SP
AM=M-1
D=M
A=A-1
M=M+D
`, nil
}

func (cmd *addVMCommand) String() string {
	return string(cmd.GetOp())
}

type subVMCommand struct{}

func (cmd *subVMCommand) GetOp() VMOperation {
	return OpSub
}

func (cmd *subVMCommand) MarshalASM() (s string, err error) {
	return `@SP
AM=M-1
D=M
A=A-1
M=M-D
`, nil
}

func (cmd *subVMCommand) String() string {
	return string(cmd.GetOp())
}

type negVMCommand struct{}

func (cmd *negVMCommand) GetOp() VMOperation {
	return OpNeg
}

func (cmd *negVMCommand) MarshalASM() (s string, err error) {
	return `@SP
A=M-1
M=-M
`, nil
}

func (cmd *negVMCommand) String() string {
	return string(cmd.GetOp())
}

type eqVMCommand struct {
	vmPc uint16
}

func (cmd *eqVMCommand) GetOp() VMOperation {
	return OpEq
}

func (cmd *eqVMCommand) MarshalASM() (s string, err error) {
	return fmt.Sprintf(`@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_EQ_%d
D;JEQ

@SP
A=M-1
M=0
@END_EQ_%d
0;JMP

(IS_EQ_%d)
@SP
A=M-1
M=-1
(END_EQ_%d)
`, cmd.vmPc, cmd.vmPc, cmd.vmPc, cmd.vmPc), nil
}

func (cmd *eqVMCommand) String() string {
	return string(cmd.GetOp())
}

type gtVMCommand struct {
	vmPc uint16
}

func (cmd *gtVMCommand) GetOp() VMOperation {
	return OpGt
}

func (cmd *gtVMCommand) MarshalASM() (s string, err error) {
	return fmt.Sprintf(`@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_GT_%d
D;JGT

@SP
A=M-1
M=0
@END_GT_%d
0;JMP

(IS_GT_%d)
@SP
A=M-1
M=-1
(END_GT_%d)
`, cmd.vmPc, cmd.vmPc, cmd.vmPc, cmd.vmPc), nil
}

func (cmd *gtVMCommand) String() string {
	return string(cmd.GetOp())
}

type ltVMCommand struct {
	vmPc uint16
}

func (cmd *ltVMCommand) GetOp() VMOperation {
	return OpLt
}

func (cmd *ltVMCommand) MarshalASM() (s string, err error) {
	return fmt.Sprintf(`@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_LT_%d
D;JLT

@SP
A=M-1
M=0
@END_LT_%d
0;JMP

(IS_LT_%d)
@SP
A=M-1
M=-1
(END_LT_%d)
`, cmd.vmPc, cmd.vmPc, cmd.vmPc, cmd.vmPc), nil
}

func (cmd *ltVMCommand) String() string {
	return string(cmd.GetOp())
}

type notVMCommand struct{}

func (cmd *notVMCommand) GetOp() VMOperation {
	return OpNot
}

func (cmd *notVMCommand) MarshalASM() (s string, err error) {
	return `
@SP
A=M-1
M=!M
`, nil
}

func (cmd *notVMCommand) String() string {
	return string(cmd.GetOp())
}

type andVMCommand struct{}

func (cmd *andVMCommand) GetOp() VMOperation {
	return OpAnd
}

func (cmd *andVMCommand) MarshalASM() (s string, err error) {
	return `@SP
AM=M-1
D=M
A=A-1

M=M&D
`, nil
}

func (cmd *andVMCommand) String() string {
	return string(cmd.GetOp())
}

type orVMCommand struct{}

func (cmd *orVMCommand) GetOp() VMOperation {
	return OpOr
}

func (cmd *orVMCommand) MarshalASM() (s string, err error) {
	return `@SP
M=M-1
A=M
D=M
A=A-1

M=M|D
`, nil
}

func (cmd *orVMCommand) String() string {
	return string(cmd.GetOp())
}

func translateAssignConstantD(constant uint16) string {
	return fmt.Sprintf(`@%d
D=A
`, constant)
}

func translateGetPointerAddr(seg VMMemSegment, index uint16) string {
	return fmt.Sprintf("@%s\n", seg.Dereference(index).ASMSymbol())
}

// Several optimizations can be done here and in fixed addr when using index=0
func translateGetDynamicAddr(seg VMMemSegment, index uint16, destRegister Register) (s string) {
	if index != 0 {
		s += translateAssignConstantD(index)
		s += fmt.Sprintf(`@%s
%s=M+D`, seg.ASMSymbol(), destRegister)
	} else {
		s += fmt.Sprintf(`@%s
%s=M`, seg.ASMSymbol(), destRegister)
	}
	s += "\n"
	return
}

func translateGetFixedAddr(seg VMMemSegment, index uint16) (s string) {
	return fmt.Sprintf("@%d\n", seg.BaseAddr()+index)
}

func translateGetStaticAddr(moduleName string, index uint16) string {
	return fmt.Sprintf("@%s.%d\n", moduleName, index)
}
