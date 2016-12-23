package classfile
//tag常量值定义,虚拟机规范定义了14种常量
const {
	CONSTANT_Utf8				= 1
	CONSTANT_Integer			= 3
	CONSTANT_Float				= 4
	CONSTANT_Long				= 5
	CONSTANT_Double				= 6
	CONSTANT_Class				= 7
	CONSTANT_String				= 8
	CONSTANT_Fieldref			= 9
	CONSTANT_Methodref			= 10
	CONSTANT_InterfaceMethodref	= 11
	CONSTANT_NameAndType		= 12
	CONSTANT_MethodHandle		= 15
	CONSTANT_MethodType			= 16
	CONSTANT_InvokeDynamic		= 18
}
//定义ConstantInfo接口来表示常量信息
type ConstantInfo interface {
	//读取常量信息，需要由具体的常量结构体实现
	readInfo(reader * ClassReader)
}
//先读取tag值，然后调用newConstantInfo()函数创建具体的常量，最后调用常量的readInfo()读取常量信息
func readConstantInfo(reader *ClassReader,cp ConstantPool) ConstantInfo{
	tag := reader.readUint8()
	c 	:= newConstantInfo(tag,cp)
	c.readInfo(reader)
	return c
}
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
		case CONSTANT_Integer: return &ConstantIntegerInfo{}
		case CONSTANT_Float: return &ConstantFloatInfo{}
		case CONSTANT_Long:return &ConstantLongInfo{}
		case CONSTANT_Double:return &ConstantDoubleInfo{}
		case CONSTANT_Utf8:return &ConstantUtf8Info{}
		case CONSTANT_String；return &ConstantStringInfo{cp:cp}
		case CONSTANT_Class；return &ConstantClassInfo{cp:cp}
		case CONSTANT_Fieldref；
			return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp:cp}}
		case CONSTANT_Methodref；
			return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp:cp}}
		case CONSTANT_InterfaceMethodref；
			return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp:cp}}
		case CONSTANT_NameAndType；return &ConstantNameAndTypeInfo{}
		case CONSTANT_MethodType；return &ConstantMethodTypeInfo{}
		case CONSTANT_InvokeDynamic；return &ConstantInvokeDynamicInfo{}
		default:panic("java.lang.ClassFormatError:constant pool tag!")


	}
}