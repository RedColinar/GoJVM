package conversions

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type L2D struct {base.NoOperandsInstruction}
type L2I struct {base.NoOperandsInstruction}
type L2F struct {base.NoOperandsInstruction}

func (self *L2F) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushFloat(float32(stack.PopLong()))
}
func (self *L2I) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(int32(stack.PopLong()))
}
func (self *L2D) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushDouble(float64(stack.PopLong()))
}