package heap

import "jvmgo/ch06/classfile"
//方法中包含字节码
type Method struct{
	//ClassMember包含accessFlags,name,descriptor
	ClassMember
	//操作数栈
	maxStack		uint
	//局部变量表
	maxLocals		uint
	//字节码
	code			[]byte
}

func newMethods(class *Class,cfMethods []*classfile.MemberInfo) []*Method{
	methods := make([]*Method,len(cfMethods))
	for i, cfMethod := range cfMethods{
		method[i] = &Method{}
		method[i].class = class
		method[i].copyMemberInfo(cfMethod)
		method[i].copyAttributes(cfMethod)
	}
	return methods
}
//把class文件中的memberInfo的maxStack，maxLocals，和字节码复制到方法结构体中
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo){
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil{
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}