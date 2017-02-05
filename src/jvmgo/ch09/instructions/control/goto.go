package control

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
//goto指令无条件跳转
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame){
	base.Branch(frame,self.Offset)
}