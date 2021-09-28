package main

// TODO: refactor to use maps when mapping constants and measure changes
import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	SegLcl     VMMemSegment = "local"
	SegArg     VMMemSegment = "argument"
	SegPointer VMMemSegment = "pointer"
	SegStatic  VMMemSegment = "static"
	SegTemp    VMMemSegment = "temp"
	SegConst   VMMemSegment = "constant"
	SegThis    VMMemSegment = "this"
	SegThat    VMMemSegment = "that"

	OpPush VMOperation = "push"
	OpPop  VMOperation = "pop"
	OpAdd  VMOperation = "add"
	OpSub  VMOperation = "sub"
	OpNeg  VMOperation = "neg"
	OpEq   VMOperation = "eq"
	OpGt   VMOperation = "gt"
	OpLt   VMOperation = "lt"
	OpAnd  VMOperation = "and"
	OpOr   VMOperation = "or"
	OpNot  VMOperation = "not"

	internalReg1 = "R13"
	internalReg2 = "R14"
	internalReg3 = "R15"

	ARegister Register = "A"
	DRegister Register = "D"
)

type Register string

var (
	segBaseAddress = map[VMMemSegment]uint16{
		SegTemp: 5,
	}

	segASMSymbol = map[VMMemSegment]string{
		SegLcl:  "LCL",
		SegArg:  "ARG",
		SegThis: "THIS",
		SegThat: "THAT",
	}

	errInvalidOperation = errors.New("invalid operation")
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename not given")
	}
	filename := os.Args[1]
	r, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}
	defer r.Close()
	moduleName := strings.TrimSuffix(filename, ".vm")
	w, err := os.Create(moduleName + ".asm")
	if err != nil {
		log.Fatalf("os.Create: %v", err)
	}
	defer w.Close()

	d := NewVMDecoder(moduleName, r)
	e := NewASMEncoder(w)
	for {
		vmc, err := d.Decode()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("e.Decode: %v", err)
		}
		if vmc == nil {
			continue
		}
		err = e.Encode(vmc)
		if err != nil {
			log.Fatalf("e.Encode: %v", err)
		}
	}
}

type VMOperation string

func (op VMOperation) IsMemoryAccess() bool {
	return op == OpPush || op == OpPop
}

//var ops = map[string]VMOperation{
//	"add": OP_ADD,
//	"sub": OP_SUB,
//	"neg": OP_NEG,
//	"eq": OP_EQ,
//	"lt": OP_LT,
//	"gt": OP_GT,
//	"and": OP_AND,
//	"or": OP_OR,
//	"not": OP_NOT,
//}

type VMMemSegment string

func (seg VMMemSegment) ASMSymbol() string {
	return segASMSymbol[seg]
}

func (m VMMemSegment) IsVirtual() bool {
	return m == SegConst
}

func (m VMMemSegment) IsDynamic() bool {
	return map[VMMemSegment]bool{
		SegLcl:  true,
		SegArg:  true,
		SegThis: true,
		SegThat: true,
	}[m]
}

func (m VMMemSegment) IsStatic() bool {
	return m == SegStatic
}

func (m VMMemSegment) IsFixed() bool {
	return m == SegTemp
}

func (m VMMemSegment) BaseAddr() uint16 {
	return segBaseAddress[m]
}

type pushVMCommand struct {
	moduleName string
	seg        VMMemSegment
	segIdx     uint16
}

func (cmd *pushVMCommand) GetOp() VMOperation {
	return OpPush
}

func (cmd *pushVMCommand) GetMemSegment() VMMemSegment {
	return cmd.seg
}

func (cmd *pushVMCommand) GetMemIndex() uint16 {
	return cmd.segIdx
}

func (cmd *pushVMCommand) init() {
	cmd.transpilePointer()
}

func (cmd *pushVMCommand) transpilePointer() {
	if cmd.seg != SegPointer {
		return
	}
	cmd.seg = map[uint16]VMMemSegment{0: SegThis, 1: SegThat}[cmd.segIdx]
	cmd.segIdx = 0
}

// TODO: refactor as bytes.Buffer
func (cmd *pushVMCommand) MarshalASM() (s string, err error) {
	s += "// push\n"
	if cmd.seg.IsVirtual() {
		s += translateAssignConstantD(cmd.segIdx)
	}
	if cmd.seg.IsDynamic() {
		s += translateGetDynamicAddr(cmd.seg, cmd.segIdx, ARegister)
	}
	if cmd.seg.IsFixed() {
		s += translateGetFixedAddr(cmd.seg, cmd.segIdx, ARegister)
	}
	if cmd.seg.IsStatic() {
		s += translateGetStaticAddr(cmd.moduleName, cmd.segIdx)
	}
	if !cmd.seg.IsVirtual() {
		s += "D=M\n"
	}
	s += `@SP
A=M
M=D
@SP
M=M+1
`
	return
}

type popVMCommand struct {
	moduleName string
	seg        VMMemSegment
	segIdx     uint16
}

func (cmd *popVMCommand) init() {
	cmd.transpilePointer()
}

func (cmd *popVMCommand) transpilePointer() {
	if cmd.seg != SegPointer {
		return
	}
	cmd.seg = map[uint16]VMMemSegment{0: SegThis, 1: SegThat}[cmd.segIdx]
	cmd.segIdx = 0
}

func translateAssignConstantD(constant uint16) string {
	s := fmt.Sprintf(`@%d
D=A`, constant)
	s += "\n"
	return s
}
func translateGetDynamicAddr(seg VMMemSegment, index uint16, destRegister Register) string {
	var s string
	s += translateAssignConstantD(index)
	s = fmt.Sprintf(`@%s
%s=M+D`, seg.ASMSymbol(), destRegister)
	s += "\n"
	return s

}
func translateGetFixedAddr(seg VMMemSegment, index uint16, destRegister Register) string {
	var s string
	s += translateAssignConstantD(index)
	s += fmt.Sprintf(`@%d
%s=A+D`, seg.BaseAddr(), destRegister)
	s += "\n"
	return s
}
func translateGetStaticAddr(moduleName string, index uint16) string {
	var s string
	s += fmt.Sprintf(`@%s.%d`, moduleName, index)
	s += "\n"
	return s
}
func (cmd *popVMCommand) MarshalASM() (s string, err error) {
	if cmd.seg.IsVirtual() {
		return "", errInvalidOperation
	}
	s += "// pop\n"
	if cmd.seg.IsDynamic() {
		s += translateGetDynamicAddr(cmd.seg, cmd.segIdx, DRegister)
	}
	if cmd.seg.IsFixed() {
		s += translateGetFixedAddr(cmd.seg, cmd.segIdx, DRegister)
	}
	if cmd.seg.IsStatic() {
		s += translateGetStaticAddr(cmd.moduleName, cmd.segIdx)
		s += fmt.Sprintf("D=A\n")
	}
	s += fmt.Sprintf(`@%s
M=D
@SP
M=M-1
A=M
D=M
@%s
A=M
M=D`, internalReg1, internalReg1)
	return
}

func (cmd *popVMCommand) GetOp() VMOperation          { return OpPop }
func (cmd *popVMCommand) GetMemSegment() VMMemSegment { return cmd.seg }
func (cmd *popVMCommand) GetMemIndex() uint16 {
	return cmd.segIdx
}

type VMCommand interface {
	init()
	GetOp() VMOperation
	MarshalASM() (string, error)
}

func NewVMCommand(absModuleName string, op VMOperation, seg VMMemSegment, segIdx uint16) (VMCommand, error) {
	cmd, ok := map[VMOperation]VMCommand{
		OpPush: &pushVMCommand{moduleName: filepath.Base(absModuleName), seg: seg, segIdx: segIdx},
		OpPop:  &popVMCommand{moduleName: filepath.Base(absModuleName), seg: seg, segIdx: segIdx},
	}[op]
	if !ok {
		return nil, nil
		//return nil, fmt.Errorf("op not supported: %s", op)
	}
	cmd.init()
	return cmd, nil
}
