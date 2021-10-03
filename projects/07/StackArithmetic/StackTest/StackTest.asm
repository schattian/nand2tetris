//push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
//eq
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_EQ_9
D;JEQ

@SP
A=M-1
M=0
@END_EQ_9
0;JMP

(IS_EQ_9)
@SP
A=M-1
M=-1
(END_EQ_9)
//push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
//eq
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_EQ_12
D;JEQ

@SP
A=M-1
M=0
@END_EQ_12
0;JMP

(IS_EQ_12)
@SP
A=M-1
M=-1
(END_EQ_12)
//push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1
//eq
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_EQ_15
D;JEQ

@SP
A=M-1
M=0
@END_EQ_15
0;JMP

(IS_EQ_15)
@SP
A=M-1
M=-1
(END_EQ_15)
//push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 891
@891
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
@IS_LT_18
D;JLT

@SP
A=M-1
M=0
@END_LT_18
0;JMP

(IS_LT_18)
@SP
A=M-1
M=-1
(END_LT_18)
//push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 892
@892
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
@IS_LT_21
D;JLT

@SP
A=M-1
M=0
@END_LT_21
0;JMP

(IS_LT_21)
@SP
A=M-1
M=-1
(END_LT_21)
//push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 891
@891
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
@IS_LT_24
D;JLT

@SP
A=M-1
M=0
@END_LT_24
0;JMP

(IS_LT_24)
@SP
A=M-1
M=-1
(END_LT_24)
//push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
//gt
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_GT_27
D;JGT

@SP
A=M-1
M=0
@END_GT_27
0;JMP

(IS_GT_27)
@SP
A=M-1
M=-1
(END_GT_27)
//push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1
//gt
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_GT_30
D;JGT

@SP
A=M-1
M=0
@END_GT_30
0;JMP

(IS_GT_30)
@SP
A=M-1
M=-1
(END_GT_30)
//push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1
//gt
@SP
AM=M-1
D=M
A=A-1

D=M-D
@IS_GT_33
D;JGT

@SP
A=M-1
M=0
@END_GT_33
0;JMP

(IS_GT_33)
@SP
A=M-1
M=-1
(END_GT_33)
//push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1
//push constant 53
@53
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
//push constant 112
@112
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
//neg
@SP
A=M-1
M=-M
//and
@SP
AM=M-1
D=M
A=A-1

M=M&D
//push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1
//or
@SP
M=M-1
A=M
D=M
A=A-1

M=M|D
//not

@SP
A=M-1
M=!M
