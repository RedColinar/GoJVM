package rtda

import . "jvmgo/ch06/rtda/heap"
//实现局部变量表的存储，和操作数栈的存储
type Slot struct {
	num int32
	//引用
	ref *Object
}
type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount >0{
		return make([]Slot,maxLocals)
	}
	return nil
}
//存取Slot中的整数
func (self Slots) SetInt(index uint,val int32){
	self[index].num = val
}
func (self Slots) GetInt(index uint) int32{
	return self[index].num
}
//float变量先转成int类型，然后按int类型处理
func (self Slots) SetFloat(index uint,val float32){
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self Slots) GetFloat(index uint) float32{
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}
//long变量拆成两个int变量
func (self Slots) SetLong(index uint,val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
//int64对应java的long
func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high :=  uint32(self[index+1].num)
	return int64(high) << 32 | int64(low)
}
//double变量可以先转成long型，在按照long类型处理
func (self Slots) SetDouble(index uint,val  float64){
	bits := math.Float64bits(val)
	self.SetLong(index ,int64(bits))
}
func (self Slots) GetDouble(index uint) float64{
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}
//引用值
func (self Slots) SetRef(index  uint,ref *Object){
	self[index].ref = ref
}
func (self Slots) GetRef(index uint) *Object{
	return self[index].ref
}