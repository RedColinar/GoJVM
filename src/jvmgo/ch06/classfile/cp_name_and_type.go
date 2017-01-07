package classfile

type ConstantNameAndTypeInfo struct{
	//字段或方法名		
	nameIndex		uint16
	//字段或方法描述符
	descriptorIndex	uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}