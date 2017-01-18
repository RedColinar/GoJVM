package heap

import "fmt"
import "jvmgo/ch08/classfile"

type Constant interface{}
//运行时常量池
type ConstantPool struct{
	//常量池对应类的索引
	class *Class
	consts []Constant
}

//根据索引返回常量
func (self *ConstantPool) GetConstant(index uint) Constant{
	if c := self.consts[index]; c !=nil{
		return c 
	}
	panic(fmt.Sprintf("No constants at index %d",index))
}
//newConstantPool()函数把class文件中的常量池转换成运行时常量池
func newConstantPool(class *Class,cfCp classfile.ConstantPool) *ConstantPool{
	cpCount := len(cfCp)
	consts := make([]Constant,cpCount)
	rtCp := &ConstantPool{class,consts}
	//i是索引
	for i := 1;i < cpCount;i++{
		cpInfo := cfCp[i]
		//类型推断为9种，
		switch cpInfo.(type){
			//前5种为：int，long，float，double，string
			case *classfile.ConstantIntegerInfo:
				intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
				consts[i] = intInfo.Value()
			case *classfile.ConstantFloatInfo:
				floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
				consts[i] = floatInfo.Value()
			//long或double型常量在常量池中的占据两个位置	
			case *classfile.ConstantLongInfo:
				longInfo := cpInfo.(*classfile.ConstantLongInfo)
				consts[i] = longInfo.Value()
				i++ 
			case *classfile.ConstantDoubleInfo:
				doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
				consts[i] = doubleInfo.Value()
				i++
			case *classfile.ConstantStringInfo:
				stringInfo := cpInfo.(*classfile.ConstantStringInfo)
				consts[i] = stringInfo.String()
			//后4种为类，字段，方法，接口方法的符号引用
			case *classfile.ConstantClassInfo:
				classInfo := cpInfo.(*classfile.ConstantClassInfo)
				consts[i] = newClassRef(rtCp, classInfo)
			case *classfile.ConstantFieldrefInfo:
				fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
				consts[i] = newFieldRef(rtCp, fieldrefInfo)
			case *classfile.ConstantMethodrefInfo:
				methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
				consts[i] = newMethodRef(rtCp, methodrefInfo)
			case *classfile.ConstantInterfaceMethodrefInfo:	
				methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
				consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		}
	}
	return rtCp
}