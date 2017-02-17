package heap

import "jvmgo/ch10/classfile"
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
	//异常表
	exceptionTable	ExceptionTable
	//行号表
	lineNumberTable *classfile.LineNumberTableAttribute
}

func newMethods(class *Class,cfMethods []*classfile.MemberInfo) []*Method{
	methods := make([]*Method,len(cfMethods))
	for i, cfMethod := range cfMethods{
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method{
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	//如果是个本地方法，则注入字节码和其他信息
	if method.IsNative(){
		method.injectCodeAttribute(md.returnType)
	}
	return method
}
//注入字节码方法
func (self *Method) injectCodeAttribute(returnType string){
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0]{
		case 'V': self.code = []byte{0xfe, 0xb1}//return
		case 'D': self.code = []byte{0xfe, 0xaf}//dreturn
		case 'F': self.code = []byte{0xfe, 0xae}//freturn
		case 'J': self.code = []byte{0xfe, 0xad}//lreturn
		case 'L', '[': self.code = []byte{0xfe, 0xb0}//areturn
		default: self.code = []byte{0xfe, 0xac}//ireturn
	}
}
//把class文件中的memberInfo的maxStack，maxLocals，和字节码复制到方法结构体中
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo){
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil{
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		//从code属性中复制异常处理表
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			self.class.constantPool)
	}
}
//获取行号方法,不是所有的方法都有行号表
func (self *Method) GetLineNumber(pc int)   int{
	if self.IsNative(){
		return -2
	}
	if self.lineNumberTable == nil{
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}
//搜索异常处理表
func (self *Method) FindExceptionHandler(exClass *Class, pc int) int{
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil{
		return handler.handlerPc
	}
	return -1
}
//
func (self *Method) calcArgSlotCount(paramTypes []string) {	
	for _, paramType := range paramTypes {
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