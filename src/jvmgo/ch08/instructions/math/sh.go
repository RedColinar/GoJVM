package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type ISHL struct { base.NoOperandsInstruction }
type ISHR struct { base.NoOperandsInstruction }
type IUSHR struct { base.NoOperandsInstruction }
type LSHL struct { base.NoOperandsInstruction }
type LSHR struct { base.NoOperandsInstruction }
type LUSHR struct { base.NoOperandsInstruction }
//int左移指令
func (self *ISHL) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}
func (self *ISHR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}
//long右移指令，long有64位，
func (self *LSHR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	//long类型有64位
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)   
}
func (self *LSHL) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}
//只有无符号右移，go中没有java的>>>运算符，
//需要先把v1转成无符号整数，位移操作之后，再转回有符号整数
func (self *IUSHR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
func (self *LUSHR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//0x3f是六个1
	s := uint32(v2) & 0x3f
	result := int64(uint32(v1) >> s)
	stack.PushLong(result)
}