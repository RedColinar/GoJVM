package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type L2D struct {base.NoOperandsInstruction}
type L2I struct {base.NoOperandsInstruction}
type L2F struct {base.NoOperandsInstruction}

func (self *L2F) Execute(frame *rtda.Frame){
	stack := frame.OperandSatck()
	l := satck.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}
func (self *L2I) Execute(frame *rtda.Frame){
	stack := frame.OperandSatck()
	stack.PushInt(int32(stack.PopLong()))
}
func (self *L2D) Execute(frame *rtda.Frame){
	stack := frame.OperandSatck()
	stack.PushDouble(float64(stack.PopLong()))
}