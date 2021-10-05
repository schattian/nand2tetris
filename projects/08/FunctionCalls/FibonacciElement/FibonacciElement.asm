@256
D=A
@SP
M=D
// save RET label addr
@RET_12121
D=A
@SP
A=M
M=D
@SP
M=M+1
//save LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
//save ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// move ARG argSize times
@SP
D=M
@5
D=D-A
@ARG
M=D
// LCL=SP
@SP
D=M
@LCL
M=D
// goto  funcName
@Sys.init
0;JMP
// label  RET_vmPc
(RET_12121)
// end func
//function Main.fibonacci 0
(Main.fibonacci)
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//push constant 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
//lt
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_LT_13
D;JLT

@SP
A=M-1
M=0
@END_LT_13
0;JMP

(IS_LT_13)
@SP
A=M-1
M=-1
(END_LT_13)
//if-goto IF_TRUE
@SP
AM=M-1
D=M
@Main.fibonacci$IF_TRUE
D;JNE
//goto IF_FALSE
@Main.fibonacci$IF_FALSE
0;JMP
//label IF_TRUE
(Main.fibonacci$IF_TRUE)
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
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
// goto RET
@RET
A=M
0;JMP
//label IF_FALSE
(Main.fibonacci$IF_FALSE)
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//push constant 2
@2
D=A
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
//call Main.fibonacci 1
// save RET label addr
@RET_23
D=A
@SP
A=M
M=D
@SP
M=M+1
//save LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
//save ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// move ARG argSize times
@SP
D=M
@6
D=D-A
@ARG
M=D
// LCL=SP
@SP
D=M
@LCL
M=D
// goto  funcName
@Main.fibonacci
0;JMP
// label  RET_vmPc
(RET_23)
// end func
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//push constant 1
@1
D=A
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
//call Main.fibonacci 1
// save RET label addr
@RET_27
D=A
@SP
A=M
M=D
@SP
M=M+1
//save LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
//save ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// move ARG argSize times
@SP
D=M
@6
D=D-A
@ARG
M=D
// LCL=SP
@SP
D=M
@LCL
M=D
// goto  funcName
@Main.fibonacci
0;JMP
// label  RET_vmPc
(RET_27)
// end func
//add
@SP
AM=M-1
D=M
A=A-1
M=M+D
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
// goto RET
@RET
A=M
0;JMP
//function Sys.init 0
(Sys.init)
//push constant 4
@4
D=A
@SP
A=M
M=D
@SP
M=M+1
//call Main.fibonacci 1
// save RET label addr
@RET_12
D=A
@SP
A=M
M=D
@SP
M=M+1
//save LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
//save ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
//save THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// move ARG argSize times
@SP
D=M
@6
D=D-A
@ARG
M=D
// LCL=SP
@SP
D=M
@LCL
M=D
// goto  funcName
@Main.fibonacci
0;JMP
// label  RET_vmPc
(RET_12)
// end func
//label WHILE
(Sys.init$WHILE)
//goto WHILE
@Sys.init$WHILE
0;JMP
