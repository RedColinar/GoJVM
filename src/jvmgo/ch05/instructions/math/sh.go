package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type ISHL struct { base.NoOperandsInstruction }
type ISHR struct { base.NoOperandsInstruction }
type IUSHR struct { base.NoOperandsInstruction }
type LSHL struct { base.NoOperandsInstruction }
type LSHR struct { base.NoOperandsInstruction }
type LUSHR struct { base.NoOperandsInstruction }
//左移指令
func (self *ISHL) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}