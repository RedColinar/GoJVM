package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
//把隐含在操作码中的常量推入操作数栈顶，不用FetchOperands()
type ACONST_NULL struct { base.NoOperandsInstruction }
type DCONST_0 struct { base.NoOperandsInstruction }
type DCONST_1 struct { base.NoOperandsInstruction }
type FCONST_0 struct { base.NoOperandsInstruction }
type FCONST_1 struct { base.NoOperandsInstruction }
type FCONST_2 struct { base.NoOperandsInstruction }
type ICONST_M1 struct { base.NoOperandsInstruction }
type ICONST_0 struct { base.NoOperandsInstruction }
type ICONST_1 struct { base.NoOperandsInstruction }
type ICONST_2 struct { base.NoOperandsInstruction }
type ICONST_3 struct { base.NoOperandsInstruction }
type ICONST_4 struct { base.NoOperandsInstruction }
type ICONST_5 struct { base.NoOperandsInstruction }
type LCONST_0 struct { base.NoOperandsInstruction }
type LCONST_1 struct { base.NoOperandsInstruction }
//aconst_null指令把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame){
	frame.OperandStack().PushRef(nil)
}
//dconst_0指令把double型0推入操作数栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushDouble(0.0)
}
func (self *DCONST_1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushDouble(1.0)
}
//float类型
func (self *FCONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(0.0)
}
func (self *FCONST_1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(1.0)
}
func (self *FCONST_2) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(2.0)
}
//iconst_m1把int型-1推入操作数栈顶
func (self *ICONST_M1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(-1)
}
func (self *ICONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(0)
}
func (self *ICONST_1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(1)
}
func (self *ICONST_2) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(2)
}
func (self *ICONST_3) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(3)
}
func (self *ICONST_4) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(4)
}
func (self *ICONST_5) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(5)
}

func (self *LCONST_0) Execute(frame *rtda.Frame){
	frame.OperandStack().PushLong(0)
}
func (self *LCONST_1) Execute(frame *rtda.Frame){
	frame.OperandStack().PushLong(1)
}