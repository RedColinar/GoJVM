package heap

import "jvmgo/ch06/classfile"
//字段不包含字节码
type Field struct{
	//ClassMember包含copyMemberInfo()
	ClassMember
}
//根据class文件的字段信息创建字段表
func newFields(class *Class,cfFields []*classfile.MemberInfo) []*Field{
	fields := make([]*Field,len(cfFields))
	for i, cfField := range cfFields{
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}