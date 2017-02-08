package heap

type Object struct {
	class 	*Class
	data	interface{}
	//记录Object结构体实例额外的信息
	extra	interface{}
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
func (self *Object) GetRefVar(name, descriptor string) *Object{
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *Object) Extra() interface{}{
	return self.extra
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object){
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
func (self *Object) SetExtra(extra interface{}){
	self.extra = extra
}