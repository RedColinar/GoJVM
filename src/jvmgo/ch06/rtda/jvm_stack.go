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
