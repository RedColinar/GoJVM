package classfile

type ConstantMemberrefInfo struct{
	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex 	uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader){
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex =  reader.readUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string{
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string,string){
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
//定义三个结构体“继承”ConstantMemberrefInfo，结构体的嵌套
//字段符号引用
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo}
//普通方法符号引用
type ConstantMethodrefInfo struct { ConstantMemberrefInfo}
//接口方法符号引用
type ConstantInterfaceMethodrefInfo struct {ConstantMemberrefInfo} 
