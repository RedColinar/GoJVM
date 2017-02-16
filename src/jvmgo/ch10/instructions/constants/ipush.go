package constants

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
//从操作数中获取byte型或short型整数，然后拓展成int型，推入栈顶
type BIPUSH struct { val int8 } //push byte
type SIPUSH struct { val int16 }//push short

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}