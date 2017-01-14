package rtda

import "jvmgo/ch06/rtda/heap"

type Thread struct {
	pc 		int
	stack	*Stack
}

func NewThread() *Thread{
	return &Thread{
		stack: newStack(1024),
	}
}
func (self *Thread) PC() int {return self.pc}
func (self *Thread) SetPC(pc int){self.pc = pc}
func (self *Thread) PushFrame(frame *Frame){
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame{
	return self.stack.pop()
}
//返回当前帧
func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}
func (self *Thread) NewFrame(method *heap.Method) *Frame{
	return newFrame(self,method)
}