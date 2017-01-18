package heap

import "jvmgo/ch08/classfile"

type FieldRef struct{
	MemberRef
	//缓存解析后的字段指针
	field 	*Field 
}

func newFieldRef(cp *ConstantPool,
		refInfo *classfile.ConstantFieldrefInfo) *FieldRef{
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class 
	c := self.ResolvedClass()
	field := lookupField(c,self.name,self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}
//字段查找步骤
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor);field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name,descriptor)
	}
	return nil
}