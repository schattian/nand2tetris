package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var VARIABLES_OFFSET = 15

var SYMBOLS_MAP = map[string]string{
	"SP":     "0",
	"LCL":    "1",
	"ARG":    "2",
	"THIS":   "3",
	"THAT":   "4",
	"R0":     "0",
	"R1":     "1",
	"R2":     "2",
	"R3":     "3",
	"R4":     "4",
	"R5":     "5",
	"R6":     "6",
	"R7":     "7",
	"R8":     "8",
	"R9":     "9",
	"R10":    "10",
	"R11":    "11",
	"R12":    "12",
	"R13":    "13",
	"R14":    "14",
	"R15":    "15",
	"SCREEN": "16384",
	"KBD":    "24576",
}

var COMP_TO_BINARY = map[string]string{
	"0":  "0101010",
	"1":  "0111111",
	"-1": "0111010",

	"D":  "0001100",
	"A":  "0110000",
	"!D": "0001101",
	"!A": "0110001",
	"-D": "0001111",
	"-A": "0110011",

	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",

	"M":  "1110000",
	"!M": "1110001",
	"-M": "1110011",

	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

var DEST_TO_BINARY = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"A":   "100",
	"MD":  "011",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

var JMP_TO_BINARY = map[string]string{
	"":    "000",
	"JEQ": "010",
	"JGE": "011",
	"JGT": "001",
	"JLT": "100",
	"JLE": "110",
	"JNE": "101",
	"JMP": "111",
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename not specified in args")
	}
	filename := os.Args[1]
	asm, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer asm.Close()

	bin, err := os.Create(strings.Split(filename, ".")[0])
	if err != nil {
		log.Fatal(err)
	}
	defer bin.Close()
	w := bufio.NewWriter(bin)

	lookupSymbols(asm)
	_, err = asm.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(asm)
	for scanner.Scan() {
		text := stripComments(scanner.Text())
		if len(text) == 0 || isLabelDef(text) {
			continue
		}
		op := NewCommand(text).Parse()
		op += "\n"
		_, err = w.WriteString(op)
		if err != nil {
			break
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}
}

type Command struct {
	Value string
	Type  CommandType
}

func NewCommand(instruction string) *Command {
	c := &Command{Value: instruction}
	if instruction[0] == '@' {
		c.Type = COMMAND_A
	} else {
		c.Type = COMMAND_C
	}

	return c
}

func (c *Command) Parse() (b string) {
	switch c.Type {
	case COMMAND_A:
		b = c.parseA()
	case COMMAND_C:
		b = c.parseC()
	}
	return
}

func (c *Command) parseC() string {
	op := c.Value
	var jmp string
	if dotComma := strings.Split(op, ";"); len(dotComma) == 2 {
		op = dotComma[0]
		jmp = dotComma[1]
	}
	var comp, dest string
	if equals := strings.Split(op, "="); len(equals) == 2 {
		comp = equals[1]
		dest = equals[0]
	} else {
		comp = op
	}
	return "111" + COMP_TO_BINARY[comp] + DEST_TO_BINARY[dest] + JMP_TO_BINARY[jmp]
}
func (c *Command) parseA() string {
	addr := replaceSymbols(c.Value[1:])
	addrNum, err := strconv.Atoi(addr)
	if err != nil {
		log.Fatalf("a instruction is not correct: %v", err)
	}
	return "0" + fmt.Sprintf("%015b", addrNum)
}

type CommandType string

var (
	COMMAND_A CommandType = "A"
	COMMAND_C CommandType = "C"
)

func lookupSymbols(r io.Reader) {
	s := bufio.NewScanner(r)
	varsMap := make(map[string]bool)
	var vars []string
	for i := 0; s.Scan(); i++ {
		op := stripComments(s.Text())
		if len(op) == 0 {
			i--
			continue
		}
		_, ok := SYMBOLS_MAP[op[1:]]
		if ok {
			continue
		}
		if isLabelDef(op) {
			label := op[1 : len(op)-1]
			SYMBOLS_MAP[label] = strconv.Itoa(i)
			i--
		} else if varname := getVariable(op); varname != "" {
			if varsMap[varname] {
				continue
			}
			varsMap[varname] = true
			vars = append(vars, varname)
		}
	}
	var j int
	for _, varname := range vars {
		if _, ok := SYMBOLS_MAP[varname]; ok {
			continue
		}
		SYMBOLS_MAP[varname] = strconv.Itoa(VARIABLES_OFFSET + j + 1)
		j++
	}
}

func getVariable(s string) string {
	if s[0] != '@' {
		return ""
	}
	s = s[1:]
	if _, err := strconv.Atoi(s); err == nil {
		return ""
	}
	return s
}

func isLabelDef(s string) bool {
	return s[0] == '(' && s[len(s)-1] == ')'
}

func replaceSymbols(s string) string {
	sVal, ok := SYMBOLS_MAP[s]
	if !ok {
		return s
	}
	return sVal
}

func stripComments(s string) string {
	return strings.TrimSpace(strings.Split(s, "//")[0])
}
