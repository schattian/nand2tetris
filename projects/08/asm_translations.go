package main

import "fmt"

func translateDefLabel(name string) string {
	return fmt.Sprintf("(%s)\n", name)
}

func translatePushRegister(addr string, r Register) (s string) {
	return fmt.Sprintf(`@%s
D=%s
@SP
A=M
M=D
@SP
M=M+1
`, addr,r)
}

func translateGoto(label string) string {
	return fmt.Sprintf(`@%s
0;JMP
`, label)
}

func translateAssignConstantD(constant uint16) string {
	return fmt.Sprintf(`@%d
D=A
`, constant)
}

func translateGetPointerAddr(seg VMMemSegment, index uint16) string {
	return fmt.Sprintf("@%s\n", seg.Dereference(index).ASMSymbol())
}

// Several optimizations can be done here and in fixed addr when using index=0
func translateGetDynamicAddr(seg VMMemSegment, index uint16, destRegister Register) (s string) {
	if index != 0 {
		s += translateAssignConstantD(index)
		s += fmt.Sprintf(`@%s
%s=M+D`, seg.ASMSymbol(), destRegister)
	} else {
		s += fmt.Sprintf(`@%s
%s=M`, seg.ASMSymbol(), destRegister)
	}
	s += "\n"
	return
}

func translateGetFixedAddr(seg VMMemSegment, index uint16) (s string) {
	return fmt.Sprintf("@%d\n", seg.BaseAddr()+index)
}

func translateGetStaticAddr(moduleName string, index uint16) string {
	return fmt.Sprintf("@%s.%d\n", moduleName, index)
}

func translateAbsLabelName(ctx, labelName string) string {
	if ctx != ""{
		return fmt.Sprintf("%s$%s", ctx, labelName)
	}
	return labelName
}
