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
//function Class1.set 0
(Class1.set)
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//pop static 0
@SP
AM=M-1
D=M
@Class1.vm.0
M=D
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
//pop static 1
@SP
AM=M-1
D=M
@Class1.vm.1
M=D
//push constant 0
@0
D=A
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
//function Class1.get 0
(Class1.get)
//push static 0
@Class1.vm.0
D=M
@SP
A=M
M=D
@SP
M=M+1
//push static 1
@Class1.vm.1
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
//function Class2.set 0
(Class2.set)
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//pop static 0
@SP
AM=M-1
D=M
@Class2.vm.0
M=D
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
//pop static 1
@SP
AM=M-1
D=M
@Class2.vm.1
M=D
//push constant 0
@0
D=A
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
//function Class2.get 0
(Class2.get)
//push static 0
@Class2.vm.0
D=M
@SP
A=M
M=D
@SP
M=M+1
//push static 1
@Class2.vm.1
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
//function Sys.init 0
(Sys.init)
//push constant 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 8
@8
D=A
@SP
A=M
M=D
@SP
M=M+1
//call Class1.set 2
// save RET label addr
@RET_10
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
@7
D=D-A
@ARG
M=D
// LCL=SP
@SP
D=M
@LCL
M=D
// goto  funcName
@Class1.set
0;JMP
// label  RET_vmPc
(RET_10)
// end func
//pop temp 0
@SP
AM=M-1
D=M
@5
M=D
//push constant 23
@23
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 15
@15
D=A
@SP
A=M
M=D
@SP
M=M+1
//call Class2.set 2
// save RET label addr
@RET_14
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
@7
D=D-A
@ARG
M=D
// LCL=SP
@SP
D=M
@LCL
M=D
// goto  funcName
@Class2.set
0;JMP
// label  RET_vmPc
(RET_14)
// end func
//pop temp 0
@SP
AM=M-1
D=M
@5
M=D
//call Class1.get 0
// save RET label addr
@RET_16
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
@Class1.get
0;JMP
// label  RET_vmPc
(RET_16)
// end func
//call Class2.get 0
// save RET label addr
@RET_17
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
@Class2.get
0;JMP
// label  RET_vmPc
(RET_17)
// end func
//label WHILE
(Sys.init$WHILE)
//goto WHILE
@Sys.init$WHILE
0;JMP
