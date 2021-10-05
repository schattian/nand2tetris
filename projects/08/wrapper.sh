export DEBUG=1
go run *.go ProgramFlow/BasicLoop/BasicLoop.vm  ProgramFlow/BasicLoop/BasicLoop.asm
go run *.go ProgramFlow/FibonacciSeries/FibonacciSeries.vm ProgramFlow/FibonacciSeries/FibonacciSeries.asm


go run *.go MemoryAccess/BasicTest/BasicTest.vm MemoryAccess/BasicTest/BasicTest.asm
go run *.go MemoryAccess/StaticTest/StaticTest.vm MemoryAccess/StaticTest/StaticTest.asm
go run *.go MemoryAccess/PointerTest/PointerTest.vm MemoryAccess/PointerTest/PointerTest.asm
go run *.go StackArithmetic/SimpleAdd/SimpleAdd.vm StackArithmetic/SimpleAdd/SimpleAdd.asm
go run *.go StackArithmetic/StackTest/StackTest.vm StackArithmetic/StackTest/StackTest.asm


go run *.go FunctionCalls/SimpleFunction/SimpleFunction.vm FunctionCalls/SimpleFunction/SimpleFunction.asm

INIT=1 go run *.go FunctionCalls/NestedCall FunctionCalls/NestedCall/NestedCall.asm
INIT=1 go run *.go FunctionCalls/FibonacciElement FunctionCalls/FibonacciElement/FibonacciElement.asm
INIT=1 go run *.go FunctionCalls/StaticsTest FunctionCalls/StaticsTest/StaticsTest.asm
