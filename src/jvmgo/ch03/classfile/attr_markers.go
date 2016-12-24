package classfile
//定义注解@Deprecated
type DeprecatedAttribute struct {
	MarkerAttribute
}
//定义关键字synthetic
//有synthetic标记的field和method是class内部使用的，正常的源代码里不会出现synthetic field
type SyntheticAttribute struct {
	MarkerAttribute
}
type MarkerAttribute struct {

}

func (self *MarkerAttribute) readInfo(reader *ClassReader){
	//read nothing
}
//由于这两个属性都没有数据，所以readInfo是空的
