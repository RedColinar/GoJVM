//判断是否为实例的逻辑
func (self *Class) isAssignable(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	//如果t不是个接口，t可能是 s的子类，
	if !t.IsInterface(){
		return s.isSubClassOf(t)
	}else {
		return s.isImplements(t)
	}
}

func (self *Class) isSubClassOf(other *Class) bool {
	//一直往上找c的父类
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		//递归
		if superInterface == iface || superInterface.isSubInterfaceOf(iface){
			return true
		}
	}
	return false
}