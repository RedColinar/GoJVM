package classfile

import "fmt"
//反映了java虚拟机规范定义的class文件格式
type ClassFile struct {
	//magic			uint32
	minorVersion	uint32
	majorVersion	uint32
	constantPool	ConstantPool
	accessFlags		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		[]uint16
	fields			[] *MemberInfo
	methods			[] *MemberInfo
	attributes		[] AttributeInfo
}
//把字节解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile,err error){
	defer func() {
		if r := recover(); r !=nil{
			var ok bool
			err, ok = r.(error){
				if !ok{
					err = fmt.Errorf("%v", r)
				}
			}
		}
	}()

	cr :=  &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
//read方法依次调用其他方法解析class文件
func (self *ClassFile) read(reader *ClassReader){

}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader){

}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader){

}
func (self *ClassFile) MinorVersion() uint16{

}
func (self *ClassFile) MajorVersion() uint16{

}