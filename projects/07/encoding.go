package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type VMDecoder struct {
	moduleName string
	r          io.Reader
	s          *bufio.Scanner
	pc         uint16
}

func NewVMDecoder(moduleName string, r io.Reader) *VMDecoder {
	return &VMDecoder{r: r, s: bufio.NewScanner(r), moduleName: moduleName}
}

func (d *VMDecoder) decodeOperation(op string) (VMOperation, error) {
	vmOp := VMOperation(op)
	switch vmOp {
	case OpPush:
	case OpPop:
	case OpAdd:
	case OpSub:
	case OpNeg:
	case OpEq:
	case OpGt:
	case OpLt:
	case OpAnd:
	case OpOr:
	case OpNot:
	default:
		return "", fmt.Errorf("unsupported operation: %s", op)
	}
	return vmOp, nil
}

func (d *VMDecoder) decodeMemSegment(memSegment string) (VMMemSegment, error) {
	vmMemSegment := VMMemSegment(memSegment)
	switch vmMemSegment {
	case SegArg:
	case SegLcl:
	case SegThis:
	case SegThat:
	case SegPointer:
	case SegStatic:
	case SegTemp:
	case SegConst:
	default:
		return "", fmt.Errorf("unsupported mem segment: %s", memSegment)
	}
	return vmMemSegment, nil
}

func (d *VMDecoder) decodeMemIndex(memIndex string) (uint16, error) {
	i, err := strconv.Atoi(memIndex)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}

func (d *VMDecoder) decode(ln string) (VMCommand, error) {
	defer d.incPc()
	ln = d.stripComments(ln)
	if ln == "" {
		return nil, nil
	}
	tokens := strings.Split(ln, " ") // op **segment **index
	op, err := d.decodeOperation(tokens[0])
	if err != nil {
		return nil, err
	}
	var memSegment VMMemSegment
	var memIndex uint16
	if op.IsMemoryAccess() {
		memSegment, err = d.decodeMemSegment(tokens[1])
		if err != nil {
			return nil, err
		}
		memIndex, err = d.decodeMemIndex(tokens[2])
		if err != nil {
			return nil, err
		}
	}
	return NewVMCommand(d.moduleName, op, memSegment, memIndex, d.pc)
}

func (d *VMDecoder) incPc() {
	d.pc += 1
}

func (d *VMDecoder) Decode() (VMCommand, error) {
	if !d.s.Scan() {
		err := d.s.Err()
		if err == nil {
			err = io.EOF
		}
		return nil, err
	}
	ln := d.s.Text()
	return d.decode(ln)
}

func (d *VMDecoder) stripComments(s string) string {
	return strings.TrimSpace(strings.Split(s, "//")[0])
}

type ASMEncoder struct {
	w io.Writer
}

func NewASMEncoder(w io.Writer) *ASMEncoder {
	return &ASMEncoder{w: w}
}

func (e *ASMEncoder) Encode(vmc VMCommand) error {
	b, err := vmc.MarshalASM()
	if err != nil {
		return err
	}
	if _, debug := os.LookupEnv("DEBUG"); debug {
		b = fmt.Sprintf("//%s\n", vmc) + b
	}
	_, err = e.w.Write([]byte(b))
	if err != nil {
		return err
	}
	return err
}
