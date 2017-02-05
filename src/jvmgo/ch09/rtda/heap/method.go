package heap

import "jvmgo/ch08/classfile"
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
	argSlotCount	uint
}

func newMethods(class *Class,cfMethods []*classfile.MemberInfo) []*Method{
	methods := make([]*Method,len(cfMethods))
	for i, cfMethod := range cfMethods{
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
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
//
func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
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
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

//方法的关键字
func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}