package classfile

import  "fmt"
import  "unicode/utf16"

type ConstantUtf8Info struct{
	str string
}
func (self *ConstantUtf8Info) readInfo(reader *ClassReader){
	//读两个字节，再转成4个字节
	length := uint32(reader.readUint16())
	bytes :=  reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}
func decodeMUTF8(bytes  []bytes) string{
	return string(bytes)
}