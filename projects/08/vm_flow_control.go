package main

import "fmt"

type labelVMCommand struct {
	labelName string
	ctxName   string
}

func (cmd *labelVMCommand) GetOp() VMOperation {
	return OpLabel
}

func (cmd *labelVMCommand) MarshalASM() (s string, err error) {
	s += translateDefLabel(cmd.absoluteLabelName())
	return
}

func (cmd *labelVMCommand) absoluteLabelName() string {
	return translateAbsLabelName(cmd.ctxName, cmd.labelName)
}

func (cmd *labelVMCommand) String() string {
	return fmt.Sprintf("%s %s", cmd.GetOp(), cmd.labelName)
}

type gotoVMCommand struct {
	labelName string
	ctxName   string
}

func (cmd *gotoVMCommand) GetOp() VMOperation {
	return OpGoto
}

func (cmd *gotoVMCommand) MarshalASM() (s string, err error) {
	s += translateGoto(cmd.absoluteLabelName())
	return
}

func (cmd *gotoVMCommand) absoluteLabelName() string {
	return translateAbsLabelName(cmd.ctxName, cmd.labelName)
}

func (cmd *gotoVMCommand) String() string {
	return fmt.Sprintf("%s %s", cmd.GetOp(), cmd.labelName)
}

type ifGotoVMCommand struct {
	labelName string
	ctxName   string
}

func (cmd *ifGotoVMCommand) GetOp() VMOperation {
	return OpIfGoto
}

func (cmd *ifGotoVMCommand) MarshalASM() (s string, err error) {
	s += translatePopToD()
	s += fmt.Sprintf(`@%s
D;JNE
`, cmd.absoluteLabelName())
	return
}

func (cmd *ifGotoVMCommand) absoluteLabelName() string {
	return translateAbsLabelName(cmd.ctxName, cmd.labelName)
}

func (cmd *ifGotoVMCommand) String() string {
	return fmt.Sprintf("%s %s", cmd.GetOp(), cmd.labelName)
}
