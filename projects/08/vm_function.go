package main

import "fmt"

type functionVMCommand struct {
	localSize uint16
	funcName  string
}

func (cmd *functionVMCommand) GetOp() VMOperation {
	return OpFunction
}

func (cmd *functionVMCommand) MarshalASM() (s string, err error) {
	s += translateDefLabel(cmd.funcName)
	push := &pushVMCommand{seg: SegConst, segIdx: 0}
	pushAsm, err := push.MarshalASM()
	if err != nil {
		return "", err
	}
	for i := 1; i <= int(cmd.localSize); i++ {
		s += pushAsm
	}
	return
}

func (cmd *functionVMCommand) String() string {
	return fmt.Sprintf("%s %s %d", cmd.GetOp(), cmd.funcName, cmd.localSize)
}

type callVMCommand struct {
	vmPc     uint16
	argSize  uint16
	funcName string
}

func (cmd *callVMCommand) GetOp() VMOperation {
	return OpCall
}

func (cmd *callVMCommand) MarshalASM() (s string, err error) {
	retLabel := fmt.Sprintf("RET_%d", cmd.vmPc)

	// fill FRAME
	s += fmt.Sprintf("// save RET label addr\n")
	s += translatePushRegister(retLabel, ARegister)

	for _, addr := range [4]string{"LCL", "ARG", "THIS", "THAT"} {
		s += fmt.Sprintf("//save %s\n", addr)
		s += translatePushRegister(addr, MRegister)
	}
	// moving ARG
	s += fmt.Sprintf(`// move ARG argSize times
@SP
D=M
@%d
D=D-A
@ARG
M=D
`, cmd.argSize+5)
	// LCL = SP
	s += `// LCL=SP
@SP
D=M
@LCL
M=D
`
	// GOTO funcName
	s += "// goto  funcName\n"
	s += translateGoto(cmd.funcName)
	// label RET_vmPc
	s += "// label  RET_vmPc\n"
	s += translateDefLabel(retLabel)
	s += "// end func\n"
	return
}

func (cmd *callVMCommand) String() string {
	return fmt.Sprintf("%s %s %d", cmd.GetOp(), cmd.funcName, cmd.argSize)
}

type returnVMCommand struct{}

func (cmd *returnVMCommand) GetOp() VMOperation {
	return OpReturn
}

func (cmd *returnVMCommand) MarshalASM() (s string, err error) {
	// FRAME = LCL
	s += `// FRAME=LCL
@LCL
D=M
@FRAME
M=D
`

	//	// RET = *(FRAME-5)
	//	translateAssignConstantD(5)
	//	s += `@FRAME
	//A=M-D
	//D=M
	//@RET
	//M=D
	//`
	s += translateAssignConstantD(5)
		s += fmt.Sprintf(`@FRAME
A=M-D
D=M
@RET
M=D
`)

	// *ARG0 = pop()
	popCmd := &popVMCommand{seg: SegArg}
	popAsm, err := popCmd.MarshalASM()
	if err != nil {
		return "", err
	}
	s += "//*ARG0 = pop()\n"
	s += popAsm

	// SP = ARG+1
	s += `// SP=ARG+1
@ARG
D=M
@SP
M=D+1
`
	// ASM_SYMBOL = *(FRAME - offset)
	for asmSymbol, offset := range map[string]uint16{
		"THAT": 1,
		"THIS": 2,
		"ARG":  3,
		"LCL":  4,
		//"RET":  5,
	} {
		s += fmt.Sprintf("// save %s = *(FRAME-%d) \n", asmSymbol, offset)
		s += translateAssignConstantD(offset)
		s += fmt.Sprintf(`@FRAME
A=M-D
D=M
@%s
M=D
`, asmSymbol)
	}
	// GOTO RET
	s += `// goto RET
@RET
A=M
0;JMP
`
	return
}
func (cmd *returnVMCommand) String() string {
	return string(cmd.GetOp())
}
