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
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
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
//getter
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}


//用来判断某个访问标志是否被设置
func (self *Method) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Method) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Method) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Method) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Method) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Method) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}