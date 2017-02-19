package math

import "math"
import "jvmgo/ch11/instructions/base"
import "jvmgo/ch11/rtda"

type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0{
		panic("java.lang.ArithmeticExeception: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}
func (self *LREM) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0{
		panic("java.lang.ArithmeticExeception: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
//go语言没有浮点型求余操作符，所以用math包的Mod()函数，
//浮点型有Infinity()无穷大，所以不用判断除数非0
func (self *DREM) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v2 == 0{
		panic("java.lang.ArithmeticExeception: / by zero")
	}
	result := math.Mod(v1,v2)
	stack.PushDouble(result)
}
func (self *FREM) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v2 == 0{
		panic("java.lang.ArithmeticExeception: / by zero")
	}
	result := float32(math.Mod(float64(v1),float64(v2)))
	stack.PushFloat(result)
}