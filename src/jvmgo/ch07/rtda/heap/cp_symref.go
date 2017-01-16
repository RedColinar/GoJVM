package heap

//符号引用有共性，使用继承减少重复代码
type SymRef struct{
	//运行时常量指针
	cp			*ConstantPool
	//类的完全限定名
	className	string
	//缓存解析后的类结构指针
	class 		*Class 
}
//类符号引用解析
func (self *SymRef) ResolvedClass() *Class{
	if self.class == nil {
		self.resolvedClassRef()
	}
	return self.class
}

func (self *SymRef) resolvedClassRef(){
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}