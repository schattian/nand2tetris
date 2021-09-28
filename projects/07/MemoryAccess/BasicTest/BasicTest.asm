// push
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@LCL
D=M+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// push
@21
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@22
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@ARG
D=M+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// pop
@ARG
D=M+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// push
@36
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THIS
D=M+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// push
@42
D=A
@SP
A=M
M=D
@SP
M=M+1
// push
@45
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@THAT
D=M+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// pop
@THAT
D=M+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// push
@510
D=A
@SP
A=M
M=D
@SP
M=M+1
// pop
@6
D=A
@5
D=A+D
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D// push
@LCL
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push
@THAT
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push
@ARG
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push
@THIS
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push
@THIS
A=M+D
D=M
@SP
A=M
M=D
@SP
M=M+1
// push
@6
D=A
@5
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
