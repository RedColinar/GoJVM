package heap

import "strings"
import "jvmgo/ch11/classfile"
type Class struct{
	//访问标志
	accessFlags		uint16
	name			string
	superClassName		string
	interfaceNames		[]string
	constantPool		*ConstantPool
	fields			[]*Field
	methods			[]*Method
	loader			*ClassLoader
	superClass 		*Class
	interfaces   		[]*Class 
	instanceSlotCount	uint 
	staticSlotCount		uint
	//
	staticVars		Slots
	initStarted		bool
	//java.lang.Class实例
	jClass			*Object
	//源文件名
	sourceFile		string
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
	class.sourceFile = getSourceFile(cf)
	return class
}
//源文件名在ClassFile结构的属性表中，getSourceFile()函数提取信息
func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil{
		return sfAttr.FileName()
	}
	return "Unknown"
}
func (self *Class) NewObject() *Object{
	return newObject(self)
}
func (self *Class) ArrayClass() *Class{
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
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
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

//判断是否有访问权限
func (self *Class) isAccessibleTo(other *Class) bool{
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}
//java/lang/Object返回包名java/lang
func (self *Class) GetPackageName() string{
	if i := strings.LastIndex(self.name,"/"); i >= 0{
		return self.name[:i]
	}
	return ""
}
//加载主类
func (self *Class) GetMainMethod() *Method {
	return self.GetStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) GetStaticMethod(name,descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name && method.descriptor == descriptor{
				return method
			}
	}
	return nil
}
func (self *Class) GetClinitMethod() *Method {
	return self.GetStaticMethod("<clinit>", "()V")
}
//改bug时增加的三个方法
func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}
func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field{
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields{
			if field.IsStatic() == isStatic &&
				field.name == name && field.descriptor == descriptor{
					return field
			}
		}
	}
	return nil
}
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
} 
// getters
func (self *Class) Name() string {
	return self.name
}
func (self *Class) JavaName() string{
	return strings.Replace(self.name,"/",".", -1 )
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) Loader() *ClassLoader {
	return self.loader
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
func (self *Class) JClass()*Object{
	return self.jClass
}
func (self *Class) SourceFile() string{
	return self.sourceFile
}
//模拟初始化
func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) StartInit() {
	self.initStarted = true
}
//
func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}
