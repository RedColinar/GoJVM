package  lang

import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

const jlClass = "java/lang/Class"

func init(){
	native.Register(jlClass, "getPrimitiveClass","(Ljava/lang/String;)Ljava/lang/Class;",getPrimitiveClass)
	native.Register(jlClass, "getName0","()Ljava/lang/String")
	native.Register(jlClass, "desiredAssertionStatus0","(Ljava/lang/Class;)Z",desiredAssertionStatus0)
}
//static native Class<?> getPrimitiveClass(string name);
func getPrimitiveClass(frame *rtda.Frame){
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}
//private native String getName0
func getName0(frame *rtda.Frame){
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}
//private static native boolean desiredAssertionStatus0(Class<?>, clazz);
func desiredAssertionStatus0(frame *rtda.Frame){
	frame.OperandStack().PushBoolean(false)
}