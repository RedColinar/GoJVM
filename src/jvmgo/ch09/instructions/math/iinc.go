package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//给局部变量表中的int变量增加常量值
type IINC struct{
	Index uint
	Const int32
}
//FetchOperands()从字节码中读取操作数
func (self *IINC) FetchOperands(reader *base.BytecodeReader){
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}
func (self *IINC) Execute(frame *rtda.Frame){
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
