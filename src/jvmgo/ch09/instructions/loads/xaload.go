package loads

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"
//<t>aload系列指令按索引取数组元素值，然后推入操作数栈
type AALOAD struct{ base.NoOperandsInstruction }
type BALOAD struct{ base.NoOperandsInstruction }
type CALOAD struct{ base.NoOperandsInstruction }
type DALOAD struct{ base.NoOperandsInstruction }
type FALOAD struct{ base.NoOperandsInstruction }
type IALOAD struct{ base.NoOperandsInstruction }
type LALOAD struct{ base.NoOperandsInstruction }
type SALOAD struct{ base.NoOperandsInstruction }

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs),index)
	stack.PushRef(refs[index])
}
func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}
func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars),index)
	stack.PushInt(int32(chars[index]))
}
func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles),index)
	stack.PushDouble(doubles[index])
}
func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats),index)
	stack.PushFloat(floats[index])
}
func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)

	ints := arrRef.Ints()
	checkIndex(len(ints),index)
	stack.PushInt(ints[index])
}
func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs),index)
	stack.PushLong(longs[index])
}
func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//数组索引
	index := stack.PopInt()
	//数组引用
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts),index)
	stack.PushInt(int32(shorts[index]))
}


func checkNotNil(ref *heap.Object){
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrlen int, index int32){
	if index < 0 || index >= int32(arrlen){
		panic("ArrayIndexOutOfBoundsException")
	}
}