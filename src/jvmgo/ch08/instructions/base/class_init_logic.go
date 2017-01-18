package base

import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"
func InitClass(thread *rtda.Thread, class *heap.Class){
	//先把类的initStarted状态设置成true，以免进入死循环
	class.StartInit()

	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}
func scheduleClinit(thread *rtda.Thread, class *heap.Class){
	clinit := class.GetClinitMethod()
	if clint != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
func initSuperClass(thread *rtda.Thread, class *heap.Class){
	if !class.IsInterface(){
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted(){
			InitClass(thread, superClass)
		}
	}
}