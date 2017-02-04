package heap
//数组对象
//给Object结构体添加几个数组特有的方法,
//没有Boolean，用[]int8来表示布尔数组，用Bytes就行
func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}
func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}
func (self *Object) Chars() []uint16 {
	return self.data.([]uint16)
}
func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}
func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}
func (self *Object) Floats() []float32 {
	return self.data.([]float32)
}
func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}
func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}
//同样没有boolean,返回数组长度
func (self *Object) ArrayLength() int32{
	//进行类型断言的必须是interface{}
	switch self.data.(type){
		case []int8: 
			return int32(len(self.data.([]int8)))
		case []int16: 
			return int32(len(self.data.([]int16)))
		case []uint16: 
			return int32(len(self.data.([]uint16)))
		case []int32: 
			return int32(len(self.data.([]int32)))
		case []int64: 
			return int32(len(self.data.([]int64)))
		case []float32: 
			return int32(len(self.data.([]float32)))
		case []float64: 
			return int32(len(self.data.([]float64)))
		case []*Object: 
			return int32(len(self.data.([]*Object)))
		default: panic("Not array!")
	}
}