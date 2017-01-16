package stores

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

//存储指令把变量从操作数栈顶弹出，然后存入局部变量表
type DSTORE struct { base.Index8Instruction}
type DSTORE_0 struct { base.NoOperandsInstruction }
type DSTORE_1 struct { base.NoOperandsInstruction }
type DSTORE_2 struct { base.NoOperandsInstruction }
type DSTORE_3 struct { base.NoOperandsInstruction }

func _dstore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index,val)
}

func (self *DSTORE) Execute(frame *rtda.Frame){
	_dstore(frame,uint(self.Index))
} 
func (self *DSTORE_0) Execute(frame *rtda.Frame){
	_dstore(frame, 0)
}   
func (self *DSTORE_1) Execute(frame *rtda.Frame){
	_dstore(frame, 1)
}   
func (self *DSTORE_2) Execute(frame *rtda.Frame){
	_dstore(frame, 2)
}   
func (self *DSTORE_3) Execute(frame *rtda.Frame){
	_dstore(frame, 3)
}   