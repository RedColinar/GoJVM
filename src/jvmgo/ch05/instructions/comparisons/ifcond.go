package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IFEQ struct { base.BranchInstruction }
type IFNE struct { base.BranchInstruction }
type IFLT struct { base.BranchInstruction }
type IFLE struct { base.BranchInstruction }
type IFGT struct { base.BranchInstruction }
type IFGE struct { base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val == 0{
		base.Branch(frame, self.Offset)
	}
}