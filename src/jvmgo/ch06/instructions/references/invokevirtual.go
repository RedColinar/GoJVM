package references

import "fmt"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

//hack 用来观察结果
func(self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame){
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor(){
		case "(Z)V":	fmt.Prinf("%v\n",stack.PopInt() != 0)
		case "(C)V":	fmt.Prinf("%c\n",stack.PopInt())
		case "(B)V":	fmt.Prinf("%v\n",stack.PopInt())
		case "(S)V":	fmt.Prinf("%v\n",stack.PopInt())
		case "(I)V":	fmt.Prinf("%v\n",stack.PopInt())
		case "(J)V":	fmt.Prinf("%v\n",stack.PopLong())
		case "(F)V":	fmt.Prinf("%v\n",stack.PopFloat())
		case "(D)V":	fmt.Prinf("%v\n",stack.PopDouble())
		default: panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}