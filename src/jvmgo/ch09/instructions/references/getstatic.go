package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type GET_STATIC struct { base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame){
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	
	if !field.IsStatic(){
		panic("java.lang.IncompatibleClassError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	//根据字段类型，从静态变量中取出相应的值，然后推入操作数栈
	switch descriptor[0] {
	case 'Z','B','C','S','I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}