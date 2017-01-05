package rtda
//实现局部变量表的存储，和操作数栈的存储
type Slot struct {
	num int32
	//引用
	ref *Object
}