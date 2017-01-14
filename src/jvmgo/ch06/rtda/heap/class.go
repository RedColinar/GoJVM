package heap

import "jvmgo/ch06/classfile"
type Class struct{
	//访问标志
	accessFlags			uint16
	name				string
	superClassName		string
	interfaceNames		[]string
	constantPool		*ConstantPool
	fields				[]*Field
	methods				[]*Method
	loader				*ClassLoader
	superClass 			*Class
	interfaces   		[]*Class 
	instanceSlotCount	uint 
	staticSlotCount		uint
	staticVars			*Slots
}

func newClass(cf *classfile.ClassFile) *Class{
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames =  cf.InterfaceNames()
	class.constantPool = newConstantPool(class,cf.ConstantPool())
	class.fields = newFields(class,cf.Fields())
	class.methods = newMethods(class,cf.Methods())
	return class
}

func (self *Class) newObject() *Object{
	return newObject(self)
}
//用来判断某个访问标志是否被设置
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
//判断是否有访问权限
func (self *Class) isAccessibleTo(other *Class) bool{
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}
//java/lang/Object返回包名java/lang
func (self *Class) getPackageName() string{
	if i := strings.LastIndex(self.name,"/"); i >= 0{
		return self.name[:i]
	}
	return ""
}
//加载主类
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) getStaticMethod(name,descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name && method.descriptor == descriptor{
				return method
			}
	}
	return nil
}