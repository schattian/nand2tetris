package main

import "fmt"

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
	s += translatePopToD()
	s += `A=A-1
M=M&D
`
	return
}

func (cmd *andVMCommand) String() string {
	return string(cmd.GetOp())
}

type orVMCommand struct{}

func (cmd *orVMCommand) GetOp() VMOperation {
	return OpOr
}

func (cmd *orVMCommand) MarshalASM() (s string, err error) {
	s += translatePopToD()
	s += `A=A-1
M=M|D
`
	return
}

func (cmd *orVMCommand) String() string {
	return string(cmd.GetOp())
}
func translatePopToD() string {
	return `@SP
AM=M-1
D=M
`
}
