package stack

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type SWAP struct{ base.NoOperandsInstruction }
//SWAP用来交换栈顶两个变量
func (self *SWAP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}