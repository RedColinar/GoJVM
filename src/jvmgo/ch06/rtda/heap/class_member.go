package heap

import "jvmgo/ch06/classfile"


type ClassMember struct{
	//访问标志
	accessFlags 	uint16
	//名字
	name			string
	//描述符
	descriptor		string
	class 			*Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo){
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

