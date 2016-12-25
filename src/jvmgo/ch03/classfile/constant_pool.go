package classfile
//ConstantInfo是个接口
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool{
	//常量池的容量
	cpCount := int(reader.readUint16())
	//用make函数来创建ConstantInfo类型切片
	cp := make([]ConstantInfo,cpCount)
	//常量池的0索引是无效索引，所以从1开始
	for i:=1; i < cpCount;i++ {
		//readConstantInfo()先读取tag值，然后调用newConstantInfo()函数创建具体的常量，
		//最后调用常量的readInfo()读取常量信息
		cp[i] = readConstantInfo(reader,cp)
		switch cp[i].(type){
		case *ConstantLongInfo,*ConstantDoubleInfo:
			i++//占两个位置
		}
	}
	return cp
}
//按照索引寻找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	if cpInfo := self[index];cpInfo != nil{
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
//从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string){
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type:= self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
} 
//从常量池查找字段或方法名和描述符
func (self ConstantPool) getClassName(index uint16) string{
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)	
}
//从常量池查找utf8字符串
func (self ConstantPool) getUtf8(index uint16) string{
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

