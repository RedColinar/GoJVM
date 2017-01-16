package heap

import "jvmgo/ch07/classfile"

type MethodRef struct{
	MemberRef
	method 		*Method
}
func newMethodRef(cp *ConstantPool,
		refInfo *classfile.ConstantMethodrefInfo) *MethodRef{
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
func (self *MethodRef) ResolvedMethod() *Method{
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}
func (self *MethodRef) resolveMethodRef(){
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface(){
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, self.name,self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAcccessError")
	}
	self.method = method
}
func lookupMethod(class *Class, name, descriptor string) *Method {
	//先从C的继承层次中去找，如果找不到就去C的接口中去找
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
