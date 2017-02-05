package stack
import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

type POP struct { base.NoOperandsInstruction }
type POP2 struct { base.NoOperandsInstruction }
//POP指令用于弹出int、float等占用一个操作数栈位置的变量
func (self *POP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PopSlot()
}
//POP2用于弹出double，long变量在操作数栈中占据两个位置的
func (self *POP2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

