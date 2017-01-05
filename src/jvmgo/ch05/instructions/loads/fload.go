package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct { base.NoOperandsInstruction }
type FLOAD_1 struct { base.NoOperandsInstruction }
type FLOAD_2 struct { base.NoOperandsInstruction }
type FLOAD_3 struct { base.NoOperandsInstruction }

func _fload(frame *rtda.Frame ,index uint){
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
func FLOAD_0(frame *rtda.Frame){
	_fload(frame,0)
}
func FLOAD_1(frame *rtda.Frame){
	_fload(frame,1)
}
func FLOAD_2(frame *rtda.Frame){
	_fload(frame,2)
}
func FLOAD_3(frame *rtda.Frame){
	_fload(frame,3)
}