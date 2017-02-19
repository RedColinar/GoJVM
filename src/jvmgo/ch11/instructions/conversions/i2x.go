package conversions

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

type I2D struct {base.NoOperandsInstruction}
type I2L struct {base.NoOperandsInstruction}
type I2F struct {base.NoOperandsInstruction}
type I2B struct {base.NoOperandsInstruction}
type I2S struct {base.NoOperandsInstruction}
type I2C struct {base.NoOperandsInstruction}
//byte对应int8
func (self *I2B) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(int32(int8(stack.PopInt())))
}
//char 对应uint16
func (self *I2C) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(int32(uint16(stack.PopInt())))
}
//short 对应int16
func (self *I2S) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushInt(int32(int16(stack.PopInt())))
}
//
func (self *I2F) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushFloat(float32(stack.PopInt()))
}
//
func (self *I2L) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushLong(int64(stack.PopInt()))
}
//
func (self *I2D) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PushDouble(float64(stack.PopInt()))
}