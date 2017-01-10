package heap

import "jvmgo/ch06/classfile"

type MemberRef struct{
	//SymRef不是属性，而是嵌套
	SymRef
	name 		string
	descriptor  string
}
//从class文件中存储的字段和方法常量中提取数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfiel.ConstantMemberrefInfo){
	self.className = refInfo.ClassName()
	self.name,self.descriptor = refInfo.NameAndDescriptor()	
}