package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	//FetchOperands()提取操作数
	FetchOperands(reader *BytecodeReader)
	//Excute执行指令逻辑
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader){}
//跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader){
	self.Offset = int(reader *ReadInt16())
}

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
