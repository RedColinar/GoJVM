package comparisons

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

type FCMPG struct { base.NoOperandsInstruction }
type FCMPL struct { base.NoOperandsInstruction }
//gFlag用来讨论两个float变量中至少有一个是NaN时，
//fcmpg返回1，fcmpl返回-1
func _fcmp(frame *rtda.Frame, gFlag bool){
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	}else if v1 == v2{
		stack.PushInt(0)
	}else if v1 < v2{
		stack.PushInt(-1)
	}else if gFlag{
		stack.PushInt(1)
	}else{
		stack.PushInt(-1)
	}
}

func (self *FCMPG) Execute(frame *rtda.Frame){
	_fcmp(frame,true)
}
func (self *FCMPL) Execute(frame *rtda.Frame){
	_fcmp(frame,false)
}
