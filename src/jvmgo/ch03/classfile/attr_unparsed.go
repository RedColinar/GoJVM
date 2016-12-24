package classfile
//未解析属性
type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}

func (self *UnparsedAttribute) readInfo(reader *CLassReader){
	self.info = reader.readBytes(self.length)
}