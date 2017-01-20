package stores

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type AASTORE struct{ base.NoOperandsInstruction }
type BASTORE struct{ base.NoOperandsInstruction }
type CASTORE struct{ base.NoOperandsInstruction }
type DASTORE struct{ base.NoOperandsInstruction }
type FASTORE struct{ base.NoOperandsInstruction }
type IASTORE struct{ base.NoOperandsInstruction }
type LASTORE struct{ base.NoOperandsInstruction }
type SASTORE struct{ base.NoOperandsInstruction }

func (self *IASTORE) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}
func (self *AASTORE) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = int32(val)
}
func (self *BASTORE) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
}