package main

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
