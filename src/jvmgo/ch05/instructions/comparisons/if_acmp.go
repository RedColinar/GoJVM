package comparisons

import "jvmgo/ch05/instructions/base"
import "jvm/ch05/rtda"

type IF_ACMPEQ struct { base.BranchInstruction }
type IF_ACMPNE struct { base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame){
	satck := frame,OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame,self.Offset)
	}
}
func (self *IF_ACMPNE) Execute(frame *rtda.Frame){
	satck := frame,OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame,self.Offset)
	}
}