package classfile

import "fmt"
//反映了java虚拟机规范定义的class文件格式
type ClassFile struct {
	//magic			uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlags		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		[]uint16
	fields			[]*MemberInfo
	methods			[]*MemberInfo
	attributes		[]AttributeInfo
	/*
	read(reader *ClassReader)
	readAndCheckMagic(reader *ClassReader)
	readAndCheckVersion(reader *ClassReader)
	*/
}
//把字节解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile,err error){
	defer func() {
		if r := recover(); r !=nil{
			var ok bool
			err, ok = r.(error)
			if !ok{
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	//读取字节码，并解析
	cr :=  &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
//read方法依次调用其他方法解析class文件，注意是依次！  
func (self *ClassFile) read(reader *ClassReader){
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader,self.constantPool)
	self.methods = readMembers(reader,self.constantPool)
	self.attributes = readAttributes(reader,self.constantPool)

}
//魔数，u43很多文件格式都会规定满足该格式的文件必须以固定字节开头
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic  != 0xCAFEBABE {
		//调用panic()方法终止程序执行
		panic("java.lang.ClassFormatError: magic!")
	}
}
//魔数之后是class文件的次版本号和主版本号，都是u2类型的
func (self *ClassFile) readAndCheckVersion(reader *ClassReader){
	//次版本号
	self.minorVersion = reader.readUint16()
	//主版本号
	self.majorVersion = reader.readUint16()
	switch self.majorVersion{
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
//从常量池查找类名
func (self *ClassFile) ClassName() string{
	return self.constantPool.getClassName(self.thisClass)
}
//从常量池查找超类类名
func (self *ClassFile) SuperClassName() string{
	if self.superClass > 0{
		return self.constantPool.getClassName(self.superClass)
	}
	return ""//只有java.lang.Object 没有超类
}
//查找接口名
func (self *ClassFile) InterfaceNames() []string{
	interfaceNames := make([]string,len(self.interfaces))
	for i, cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
//下面的方法是Getter方法，把结构体的字段暴露给其他包使用
func (self *ClassFile) MinorVersion() uint16{
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16{
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool{
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16{
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo{
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo{
	return self.methods
}
//
func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}