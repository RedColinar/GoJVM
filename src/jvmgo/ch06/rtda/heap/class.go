package heap

import "jvmgo/ch06/classfile"
type Class struct{
	//访问标志
	accessFlags			uint16
	name				string
	superClassName		string
	interfaceNames		[]string
	constantPool		*ConstantPool
	fields				[]*Field
	methods				[]*Method
	loader				*ClassLoader
	superClass 			*Class
	interfaces   		[]*Class 
	instanceSlotCount	uint 
	staticSlotCount		uint
	staticVars			*Slots
}