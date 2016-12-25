package classfile

type AttributeInfo interface{
	readInfo(reader *ClassReader)
}
//读取属性表
func readAttributes(reader *ClassReader,cp ConstantPool) []AttributeInfo{
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo,attributesCount)
	for i := range attributes{
		attributes[i] = readAttribute(reader,cp)
	}
	return attributes
}

//读取单个属性。先读取属性名索引，根据他从常量池中找到属性名，然后读取属性长度，
func readAttribute(reader *ClassReader,cp ConstantPool) AttributeInfo{
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo:=newAttributeInfo(attrName,attrLen,cp)
	attrInfo.readInfo(reader)
	return attrInfo
}
//接着调用newAttributeInfo()创建具体的属性实例。java虚拟机预定义了23种属性，先读8种
func newAttributeInfo(attrName string,attrLen uint32,cp ConstantPool) AttributeInfo{
	switch attrName{
		case "Code":return &CodeAttribute{cp:cp}//&构造指向CodeAttribute的指针，构造方式为{key：value}
		case "ConstantValue":return &ConstantValueAttribute{}
		case "Deprecated":return &DeprecatedAttribute{}
		case "Exceptions":return &ExceptionsAttribute{}
		case "LineNumberTable":return &LineNumberTableAttribute{}
		case "LocalVariableTable":return &LocalVariableTableAttribute{}
		case "SourceFile":return &SourceFileAttribute{cp:cp}
		case "Synthetic":return &SyntheticAttribute{}
		default:return &UnparsedAttribute{attrName,attrLen,nil}
	} 
}