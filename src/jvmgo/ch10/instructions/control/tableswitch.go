package control

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"

type TABLE_SWITCH struct{
	//默认情况下执行跳转所需的字节码偏移量
	defaultOffset	int32
	//low和high记录case的取值范围
	low				int32
	high			int32
	//jumpOffsets是一个索引表，里面存放对应case情况个int值，
	//对应着执行跳转所需的字节码偏移量
	jumpOffsets		[]int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame){
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high{
		offset = int(self.jumpOffsets[index-self.low])
	}else{
		offset = int(self.defaultOffset)
	}
	base.Branch(frame,offset)
}