package constants

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
//用于加载long和double常量
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC) Execute(frame *rtda.Frame){
	_ldc(frame, self.Index)
}
func (self *LDC_W) Execute(frame *rtda.Frame){
	_ldc(frame, self.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)
	
	switch c.(type){
	case int32: stack.PushInt(c.(int32))
	case  float32: stack.PushFloat(c.(float32))
	case string: 
		//如果_ldc试图从运行时常量池中加载字符串常亮，先通过常量拿到go字符串，然后把它转化成java字符串实例
		internedStr := heap.JString(class.Loader(), c.(stirng))
		//把引用推入栈顶
		stack.PushRef(internedStr)
	//cse *heap.ClassRef:
	default:  panic("todo: ldc!")
	}
}
//LDC2_W
func (self *LDC2_W)  Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64: stack.PushLong(c.(int64))
	case float64: stack.PushDouble(c.(float64))
	default: panic("java.lang.ClassFromatError")
	}
}
