package rtda

type Stack struct {
	maxSize uint
	size 	uint
	//保存栈顶指针
	_top	*Frame
}

func newStack(maxSize uint) *Stack{
	return &Stack{
		maxSize:maxSize,
	}
}
//清空虚拟栈
func (self *Stack) clear(){
	for !self.isEmpty(){
		self.pop()
	}
}
//拿到完整的java虚拟机栈
func (self *Stack) getFrames() []*Frame{
	frames := make([]*Frame, 0, self.size)
	for frame := self._top; frame != nil; frame = frame.lower{
		frames = append(frames,frame)
	}
	return frames
}
//push方法把帧推入栈顶
func (self *Stack) push(frame *Frame){
	if self.size >= self.maxSize{
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil{
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}
//弹出栈顶
func (self *Stack) pop() *Frame{
	if self._top == nil{
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}
//返回栈顶
func (self *Stack) top() *Frame{
	if self._top == nil{
		panic("jvm stack is empty!")
	}
	return self._top 
}
func (self *Stack) isEmpty() bool{
	return self._top == nil
}