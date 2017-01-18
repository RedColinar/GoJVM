package heap

type Object struct {
	class 	*Class
	data	interface{}
} 
func newObject(class *Class) *Object {
	return &Object{
		class:		class,
		data:		newSlots(class.instanceSlotCount),
	}
}
//实际逻辑在Class结构体中的isAssignableFrom()方法
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}