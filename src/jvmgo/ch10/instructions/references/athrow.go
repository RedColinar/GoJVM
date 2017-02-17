package  references

import "reflect"
import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

type ATHROW struct{ base.NoOperandsInstruction }
func (self *ATHROW) Execute(frame *rtda.Frame){
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}


	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex){
		handleUncaughtException(thread, ex)
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool{
	for{
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPC > 0{
			stack := frame.OperandStack()
			//找到异常处理，在跳转到异常处理代码之前 ，先把操作数栈清空 ，
			//然后把异常对象引用推入栈顶
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty(){
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtda.Thread, ex *heap.Object){
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage","Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": "+ goMsg)
	//反射
	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len();i++{
		ste := stes.Index(i).Interface().(interface{
			String() string
			})
		println("\tat " + ste.String())
	}
}