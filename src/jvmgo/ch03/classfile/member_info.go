package classfile
//用一个结构体统一表示字段和方法
type MemberInfo struct {
	//cp ConstantPool
	accessFlags 	uint16
	nameIndex		uint16
	descriptorIndex	uint16
	//attributes		[]AttributeInfo
}