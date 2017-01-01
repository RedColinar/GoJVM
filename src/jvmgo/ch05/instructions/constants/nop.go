package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type ACONST_NULL struct { base.NoOperandsInstruction }
type DCONST_0 struct { base.NoOperandsInstruction }
type ICONST_M1 struct { base.NoOperandsInstruction }
//aconst_null指令把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame){
	frame.OperandStack().PushRef(nil)
}
//dconst_0指令把double型0推入操作数栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushDouble(0.0)
}
//iconst_m1把int型-1推入操作数栈顶
func (self *ICONST_M1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(-1)
}