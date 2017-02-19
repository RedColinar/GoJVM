package heap

import "jvmgo/ch11/classfile"
//字段不包含字节码
type Field struct{
	//ClassMember包含copyMemberInfo()
	ClassMember
	constValueIndex	uint
	slotId			uint
}
//根据class文件的字段信息创建字段表
func newFields(class *Class,cfFields []*classfile.MemberInfo) []*Field{
	fields := make([]*Field,len(cfFields))
	for i, cfField := range cfFields{
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}
func (self *Field) isLongOrDouble() bool{
	return self.descriptor == "J" || self.descriptor == "D"
}
func (self *Field) copyAttributes(cfField *classfile.MemberInfo){
	if valAttr := cfField.ConstantValueAttribute();valAttr !=nil{
		self.constValueIndex =  uint(valAttr.ConstantValueIndex())
	}
}
//用来判断某个访问标志是否被设置
func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
//getter
func (self *Field) SlotId() uint {
	return self.slotId
}
func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}