package classfile

import "encoding/binary"
//定义结构体用来读取数据
type ClassReader struct{
	data []byte
}

func (self *ClassReader) readUint8() uint8{
	val := self.data[0]
	self.data = self.data[1:]
	return val
}
/*	
    51	func (littleEndian) Uint16(b []byte) uint16 {
    52		_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
    53		return uint16(b[0]) | uint16(b[1])<<8
    54	}

*/
func (self *ClassReader) readUint16() uint16{
	//标准库encoding/binary包定义了一个变量BigEndian，可以从[]byte中解码多字节数据。
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}
func (self *ClassReader) readUint32() uint32{
	val := binary.BigEndian.Uint32(self.data)
	self.data =  self.data[4:]
	return val
}
//java虚拟机规范没有定义u8
func (self *ClassReader) readUint64() uint64{
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}
func (self *ClassReader) readUint16s() []uint16{
	n := self.readUint16()
	s := make([]uint16,n)
	for i:=range s {
		s[i] =  self.readUint16()
	}
	return s
}
//用于读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte{
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}