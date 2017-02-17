package lang

import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

func inti(){
	native.Register("java/lang/Throwable","fillInStackTrace",
		"(I)Ljava/lang/Throwable;",fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame){
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := creatStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}
//记录java虚拟机栈帧信息
type StackTraceElement struct{
	fileName		string
	//声明方法的类名
	className		string
	//方法名
	methodName		string
	//帧正在执行的第几行代码
	lineNumber		int
}
//好像返回值 不能换行
func creatStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement{
	//由于栈顶两帧正在执行fillStackTrace(int)和fillIntStackTrace()方法，所以跳过这两针
	skip := distanceToObject(tObj.Class()) + 2
	//reslice
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames{
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}
//计算好需要跳过的帧数
func distanceToObject(class *heap.Class) int{
	distance := 0
	for c := class.SuperClass();c != nil; c = c.SuperClass(){
		distance ++
	}
	return distance
}
func createStackTraceElement(frame *rtda.Frame) *StackTraceElement{
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:	class.SourceFile(),
		className: 	class.JavaName(),
		methodName:	method.Name(),
		lineNumber:	method.GetLineNumber(frame.NextPC() - 1),	
	}
}
