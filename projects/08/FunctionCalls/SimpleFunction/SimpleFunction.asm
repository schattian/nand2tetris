//function SimpleFunction.test 2
(SimpleFunction.test)
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
//push local 0
@LCL
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//push local 1
@1
D=A
@LCL
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
//add
@SP
AM=M-1
D=M
A=A-1
M=M+D
//not

@SP
A=M-1
M=!M
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//add
@SP
AM=M-1
D=M
A=A-1
M=M+D
//push argument 1
@1
D=A
@ARG
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
//sub
@SP
AM=M-1
D=M
A=A-1
M=M-D
//return
// FRAME=LCL
@LCL
D=M
@FRAME
M=D
@5
D=A
@FRAME
A=M-D
D=M
@RET
M=D
//*ARG0 = pop()
@ARG
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// SP=ARG+1
@ARG
D=M
@SP
M=D+1
// save ARG = *(FRAME-3) 
@3
D=A
@FRAME
A=M-D
D=M
@ARG
M=D
// save LCL = *(FRAME-4) 
@4
D=A
@FRAME
A=M-D
D=M
@LCL
M=D
// save THAT = *(FRAME-1) 
@1
D=A
@FRAME
A=M-D
D=M
@THAT
M=D
// save THIS = *(FRAME-2) 
@2
D=A
@FRAME
A=M-D
D=M
@THIS
M=D
// goto RET
@RET
A=M
0;JMP
