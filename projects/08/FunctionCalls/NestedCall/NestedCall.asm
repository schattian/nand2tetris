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
//function Sys.init 0
(Sys.init)
//push constant 4000
@4000
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop pointer 0
@SP
AM=M-1
D=M
@THIS
M=D
//push constant 5000
@5000
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop pointer 1
@SP
AM=M-1
D=M
@THAT
M=D
//call Sys.main 0
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
@Sys.main
0;JMP
// label  RET_vmPc
(RET_12)
// end func
//pop temp 1
@SP
AM=M-1
D=M
@6
M=D
//label LOOP
(Sys.init$LOOP)
//goto LOOP
@Sys.init$LOOP
0;JMP
//function Sys.main 5
(Sys.main)
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
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 4001
@4001
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop pointer 0
@SP
AM=M-1
D=M
@THIS
M=D
//push constant 5001
@5001
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop pointer 1
@SP
AM=M-1
D=M
@THAT
M=D
//push constant 200
@200
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop local 1
@1
D=A
@LCL
D=M+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
//push constant 40
@40
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop local 2
@2
D=A
@LCL
D=M+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
//push constant 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop local 3
@3
D=A
@LCL
D=M+D
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
//push constant 123
@123
D=A
@SP
A=M
M=D
@SP
M=M+1
//call Sys.add12 1
// save RET label addr
@RET_37
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
@Sys.add12
0;JMP
// label  RET_vmPc
(RET_37)
// end func
//pop temp 0
@SP
AM=M-1
D=M
@5
M=D
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
//push local 2
@2
D=A
@LCL
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
//push local 3
@3
D=A
@LCL
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
//push local 4
@4
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
//add
@SP
AM=M-1
D=M
A=A-1
M=M+D
//add
@SP
AM=M-1
D=M
A=A-1
M=M+D
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
//function Sys.add12 0
(Sys.add12)
//push constant 4002
@4002
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop pointer 0
@SP
AM=M-1
D=M
@THIS
M=D
//push constant 5002
@5002
D=A
@SP
A=M
M=D
@SP
M=M+1
//pop pointer 1
@SP
AM=M-1
D=M
@THAT
M=D
//push argument 0
@ARG
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//push constant 12
@12
D=A
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
