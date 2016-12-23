package classfile
//表示类或接口的符号引用
type ConstantClassInfo struct{
	cp 			ConstantPool
	nameIndex 	uint16
}

func (self *ConstantClassInfo) readerInfo(reader *ClassReader){
	self.nameIndex = reader.readerUint16()
}
       
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}