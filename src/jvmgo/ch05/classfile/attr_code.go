package classfile

type CodeAttribute struct {
	cp				ConstantPool
	//操作数栈的最大深度
	maxStack		uint16
	//局部变量表大小
	maxLocals		uint16
	//字节码
	code 			[]byte
	//异常处理表
	exceptionTable  []*ExceptionTableEntry
	//属性表
	attributes		[]AttributeInfo
}
//定义异常处理表的结构体
type ExceptionTableEntry struct {
	startPc		uint16
	endPc		uint16
	handlerPc	uint16
	catchType	uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader){
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	//新的方法
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader,self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry{
	//生成ExceptionTableEntry类型的切片，先获取长度
	exceptionTableLength := reader.readUint16()
	exceptionTable :=  make([]*ExceptionTableEntry,exceptionTableLength)
	for i := range exceptionTable{
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:	reader.readUint16(),
			endPc:		reader.readUint16(),
			handlerPc:	reader.readUint16(),
			catchType:	reader.readUint16(),
		}
	}
	return exceptionTable
}
//getter
func (self *CodeAttribute) MaxStack() uint{
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint{
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte{
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry{
	return self.exceptionTable
}