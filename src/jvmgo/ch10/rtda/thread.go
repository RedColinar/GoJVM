package rtda

import "jvmgo/ch10/rtda/heap"

type Thread struct {
	pc 		int
	stack	*Stack
}

func NewThread() *Thread{
	return &Thread{
		stack: newStack(1024),
	}
}
//清空虚拟机栈
func (self *Thread) ClearStack(){
	self.stack.clear()
}
func (self *Thread) GetFrames() []*Frame{
	return self.stack.getFrames()
}
func (self *Thread) PC() int {return self.pc}
func (self *Thread) SetPC(pc int){self.pc = pc}
func (self *Thread) PushFrame(frame *Frame){
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame{
	return self.stack.pop()
}
//两个返回当前帧的方法
func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}
func (self *Thread) TopFrame() *Frame{
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame{
	return newFrame(self,method)
}
//判断虚拟机栈中是否还有帧
func (self *Thread) IsStackEmpty() bool{
	return self.stack.isEmpty()
}