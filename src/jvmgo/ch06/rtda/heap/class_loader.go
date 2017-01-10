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
	return *ClassLoader{
		cp:			cp,
		classMap:	make(map[string]*Class),
	}
}
//把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class{
	if class, ok := self.ClassMap[name]; ok  {
		//类已加载 
		return class
	}
	//数组类和普通类有很大的不同，它的数据并不是来自class文件，
	//而是由java虚拟机在运行期间生成
	return self.loadNonArrayClass(name)
}
func (self *ClassLoader) loadNonArrayClass(name string) *Class{
	data, entry := self.readClass(data)
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
