package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame){
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	if class.IsInterface() || class.IsAbstract(){
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}