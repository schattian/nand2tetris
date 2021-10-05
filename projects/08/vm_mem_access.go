package main

import "fmt"

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


func (cmd *popVMCommand) String() string {
	return fmt.Sprintf("%s %s %d", cmd.GetOp(), cmd.seg, cmd.segIdx)
}

func (cmd *popVMCommand) GetOp() VMOperation {
	return OpPop
}
