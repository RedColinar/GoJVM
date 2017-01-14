package heap


import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"
//类加载器
type ClassLoader struct{
	cp 			*classpath.Classpath
	//记录已经加载的类数据 
	classMap	map[string]*Class 
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader{
	return &ClassLoader{
		cp:			cp,
		classMap:	make(map[string]*Class),
	}
}
//把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class{
	if class, ok := self.classMap[name]; ok  {
		//类已加载 
		return class
	}
	//数组类和普通类有很大的不同，它的数据并不是来自class文件，
	//而是由java虚拟机在运行期间生成
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(name string) *Class{
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("Loaded %s from %s",name,entry)
	return class
}
//根据文件名找到class文件，，返回文件字节码，和类路径接口
func (self *ClassLoader) readClass(name string) ([]byte,classpath.Entry){
	data, entry, err := self.cp.ReadClass(name)
	if err != nil{
		panic("java.lang.ClassNotFoundException: "+ name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class{
	//把class文件转换成Class结构体
	class := parseClass(data)
	//
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}
func parseClass(data []byte) *Class{
	cf, err  := classfile.Parse(data)
	if err != nil{
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
	//除了java.lang.Object外，所有类都有且仅有一个超类
	//因此，除非是Object类，否则需要递归调用LoadClass()加载超类
func resolveSuperClass (class *Class){
	if class.name != "java/lang/Object"{
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
	//递归调用LoadClass()加载类的每一个直接接口
func resolveInterfaces(class *Class){
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames{
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}	
	//类的链接分为验证和准备两个必要阶段
func link(class *Class){
	verify(class)
	prepare(class)
}
func verify(class *Class){
	//在java虚拟机规范中4.10介绍了类的验证算法，太长不想写…
} 
func prepare(class *Class){
	calcInstantceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}
//计算实例字段(非静态的)的个数，同时编号
func calcInstantceFieldSlotIds(class *Class){
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic(){
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
//计算静态字段的个数
func calcStaticFieldSlotIds(class *Class){
	slotId := uint(0)
	for _, field := range class.fields{
		if field.IsStatic(){
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}
}
//给类变量分配空间，然后给他们赋予初始值
func allocAndInitStaticVars(class *Class){
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields{
		if field.IsStatic() && field.IsFinal(){
			initStaticFinalVar(class, field)
		}
	}
}
//
func initStaticFinalVar(class *Class, field *Field){
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId :=  field.SlotId()

	if cpIndex > 0{
		switch field.Descriptor(){
		case "Z","B","C","S","I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId,val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId,val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId,val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}