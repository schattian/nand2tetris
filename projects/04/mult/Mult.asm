// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.
// Put your code here.
(INIT)
    @i
    M=0
    @total
    M=0
(LOOP)
    // circuit breaker
    @i
    D=M
    @R1
    D=M-D  // factor - iteration
    @END
    D;JEQ
    @R0
    D=M
    @total
    M=M+D
    @i
    M=M+1
    @LOOP
    0;JMP
(END)
    @total
    D=M
    @R2
    M=D
    @END
    0;JMP