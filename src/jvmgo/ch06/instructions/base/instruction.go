package base

import "jvmgo/ch06/rtda"
//指令接口，所有指令都实现这个接口
type Instruction interface {
	//FetchOperands()提取操作数
	FetchOperands(reader *BytecodeReader)
	//Excute执行指令逻辑
	Execute(frame *rtda.Frame)
}
//隐含操作数的命令，嵌套NoOperandsInstruction,
//因为不需要取操作数
type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader){}
//跳转指令嵌套BranchInstruction，
//定义跳转的偏移量，偏移量都是取两字节的操作数，
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader){
	self.Offset = int(reader.ReadInt16())
}
//需要一个局部变量表索引的嵌套Index8Instruction
//这个局部变量表索引读取8位数据获取
type Index8Instruction struct {
	//局部变量表索引
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader){
	self.Index = uint(reader.ReadUint8())
}
//访问运行时常量池，常量池索引由两字节操作数给出
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader){
	self.Index = uint(reader.ReadUint16())
}
