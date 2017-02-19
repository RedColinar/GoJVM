package references

import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"
import "jvmgo/ch11/rtda/heap"

type CHECK_CAST struct { base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	//checkcast 和instanceof的区别在于
	//instanceof会改变操作数栈(弹出的对象引用，推入判断结果)
	//checkcast不改变操作数栈(弹出再推入，接着进行判断，如果为空，抛异常)
	stack.PushRef(ref)
	if ref == nil{
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
