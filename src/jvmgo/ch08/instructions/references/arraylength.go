package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type ARRAY_LENGTH struct{ base.NoOperandsInstruction }
//arraylength指令只需要一个操作数，即数组引用
func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}