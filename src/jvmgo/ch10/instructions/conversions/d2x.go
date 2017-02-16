package conversions

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
func (self *D2F) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := float32(d)
	stack.PushFloat(i)
}
func (self *D2L) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int64(d)
	stack.PushLong(i)
}