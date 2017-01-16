package control

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
//goto指令无条件跳转
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame){
	base.Branch(frame,self.Offset)
}