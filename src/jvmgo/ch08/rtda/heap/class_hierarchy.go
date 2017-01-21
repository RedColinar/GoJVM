package heap
//判断是否为实例的逻辑
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	//如果t不是个接口，t可能是 s的子类，
	if !s.IsArray(){
		//s不是数组
		if !s.IsInterface(){
			//s不是接口，是个类
			if !t.IsInterface(){
				return s.IsSubClassOf(t)
			}else {
				return s.IsImplements(t)
			}
		}else{
			if !t.IsInterface(){
				return t.isJlObject()
			}else{
				return t.isSuperInterfaceOf(s)
			}
		}
	}else{
		if !t.IsArray(){
			if !t.IsInterface(){
				return t.isJlObject()
			}else{
				return t.isJlCloneable() || t.isJioSerializable()
			}
		}else{
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}			
	return false
}

func (self *Class) IsSubClassOf(other *Class) bool {
	//一直往上找c的父类
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		//递归
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface){
			return true
		}
	}
	return false
}
// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}
// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}
