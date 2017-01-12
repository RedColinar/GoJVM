package rtda

type Object struct {
	class 	*Class
	fields	Slots
} 
func newObject(class *Class) *Object {
	return &Object{
		class:		class,
		fields:		newSlots(class.instanceSlotCount),
	}
}
//实际逻辑在Class结构体中的isAssignableFrom()方法
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}