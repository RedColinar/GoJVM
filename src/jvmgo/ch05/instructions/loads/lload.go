package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct { base.NoOperandsInstruction }
type LLOAD_1 struct { base.NoOperandsInstruction }
type LLOAD_2 struct { base.NoOperandsInstruction }
type LLOAD_3 struct { base.NoOperandsInstruction }

func _lload(frame *rtda.Frame ,index uint){
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
func LLOAD_0(frame *rtda.Frame){
	_lload(frame,0)
}
func LLOAD_1(frame *rtda.Frame){
	_lload(frame,1)
}
func LLOAD_2(frame *rtda.Frame){
	_lload(frame,2)
}
func LLOAD_3(frame *rtda.Frame){
	_lload(frame,3)
}